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

var wg sync.WaitGroup

func Do(filepath string, url string) (err error) {
	sizeTotal, err := DataLength(url)
	if err != nil {
		return err
	}
	numDivide := NumDivideRange(sizeTotal)
	sizeDivide := sizeTotal / int64(numDivide)
	ch := make([]chan string, numDivide)
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < numDivide; i++ {
		wg.Add(1)
		ch[i] = download(ctx, i, numDivide, sizeDivide, sizeTotal, url)
	}
	dstFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	for _, ch := range ch {
		name := <-ch
		srcfile, err := os.Open(name)
		if err != nil {
			return err
		}
		_, err = io.Copy(dstFile, srcfile)
		if err != nil {
			return err
		}
	}
	cancel()
	wg.Wait()
	return err
}

func download(ctx context.Context, index int, numDivide int, sizeDivide int64, sizeTotal int64, url string) chan string {
	fileCh := make(chan string)
	go func() {
		defer wg.Done()
		minRange := sizeDivide * int64(index)
		maxRange := sizeDivide * int64(index+1)
		if index == numDivide-1 {
			maxRange += sizeTotal - maxRange
		}
		fmt.Printf("index=%v, min=%v, max=%v\n", index, minRange, maxRange-1)
		client := &http.Client{}
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
		tmpfile, err := os.CreateTemp("", "")
		if err != nil {
			log.Fatal(err)
		}
		defer func() {
			fmt.Println("\nrm %v", tmpfile.Name())
			os.Remove(tmpfile.Name())
		}()
		fmt.Printf("%v ", tmpfile.Name())
		_, err = io.Copy(tmpfile, resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if err = tmpfile.Close(); err != nil {
			log.Fatal(err)
		}
	LOOP:
		for {
			select {
			case <-ctx.Done():
				break LOOP
			case fileCh <- tmpfile.Name():
			}
		}
		close(fileCh)
	}()
	return fileCh
}
