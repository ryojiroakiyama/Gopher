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
	ONEDLMAX = 1000
)

type ctxKeyUrl struct{}

var wg sync.WaitGroup

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
	defer dstFile.Close()
	for _, fileName := range readers {
		get := <-fileName
		if get == nil {
			os.Remove(dstFile.Name())
			return fmt.Errorf("error")
		}
		_, err = io.Copy(dstFile, <-fileName)
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
		minRange, maxRange := DownloadRange(index, numDiv, sizeDiv, sizeSum)
		//fmt.Printf("index=%v, min=%v, max=%v\n", index, minRange, maxRange-1)
		client := &http.Client{}
		url, ok := ctx.Value(ctxKeyUrl{}).(string)
		if !ok {
			log.Fatal("no")
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
