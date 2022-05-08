package pget

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	DivDownLoadMax = 1000
)

type ctxKeyUrl struct{}

var wg sync.WaitGroup

// - errorgroupsとstringのスライスを使わず, waitgroupとチャネルのスライスで実装
// - 一時的なファイルを生成せず, goroutineからmainへチャネルで直接レスポンスのボディを伝えていた, mainではチャネルをスライスで管理
// - 全体的な流れ,
// main:         必要な情報用意→download関数: 事前にチャネルをmainに戻してgoroutine
// goroutine: ダウンロードしてresponseをmainに送り続ける
// (キャンセルされるまでresponseをclose()しないための送り続ける処理)
// main:         responseを受け取って処理が終わればcancel()を実行,
// contextなので全てのgoroutineでキャンセルされる
// goroutine: キャンセルを受け取って終了, done()を実行, responseをclose
// main:          waitでgoroutineの終わりを待つ, done()が全てのgoroutineで呼ばれたら終了
// (処理が終われば今度は全てのresponseは閉じてほしいのでgoroutineの終了を待つ)
// - contextのkey
//     - 非公開型で指定することで当パッケージ内でのみ考慮すれば, パッケージ間でのkey被りを防げる, さらにパッケージ内でも型自体を毎回定義することで被らなくなる
//     - struct{}で容量の削減, struct{}で渡してるのでおそらくその型のnilがkeyになってるのかな
func Do(filepath string, url string) (err error) {
	sizeSum, err := DataLength(url)
	if err != nil {
		return err
	}
	numDiv := NumDivideRange(sizeSum)
	sizeDiv := sizeSum / int64(numDiv)
	ctx, cancel := context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, ctxKeyUrl{}, url)
	readers := make([]chan io.Reader, numDiv)
	for i := 0; i < numDiv; i++ {
		wg.Add(1)
		readers[i] = download(ctx, i, numDiv, sizeDiv, sizeSum)
	}
	dstFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := dstFile.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()
	for _, fileName := range readers {
		get := <-fileName
		if get == nil { //gorougine側でerrorが起きればnil
			os.Remove(dstFile.Name())
			return fmt.Errorf("error")
		}
		_, err = io.Copy(dstFile, get)
		if err != nil {
			return err
		}
	}
	cancel()
	wg.Wait()
	return err
}

func download(ctx context.Context, index int, numDiv int, sizeDiv int64, sizeSum int64) chan io.Reader {
	outFileName := make(chan io.Reader)
	go func() {
		defer wg.Done()
		minRange, maxRange := downloadRange(index, numDiv, sizeDiv, sizeSum)
		client := &http.Client{}
		url, ok := ctx.Value(ctxKeyUrl{}).(string)
		if !ok {
			log.Fatal("No value")
		}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("Range", RangeValue(minRange, maxRange-1))
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case outFileName <- resp.Body:
			}
		}
		close(outFileName)
	}()
	return outFileName
}
