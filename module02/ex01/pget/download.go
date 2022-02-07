package pget

import (
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
	chStr := make([]chan string, numDivide)
	tmpFileNames := make([]string, numDivide)
	defer func() {
		for _, t := range tmpFileNames {
			if t != "" {
				fmt.Println("\ntmpfile:", t)
				os.Remove(t)
			}
		}
	}()
	for i := 0; i < numDivide; i++ {
		chStr[i] = make(chan string)
		wg.Add(1)
		go download(i, numDivide, sizeDivide, sizeTotal, url, chStr[i])
	}
	fmt.Println("\n---------------------")
	dstFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	for _, ch := range chStr {
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
	wg.Wait()
	return err
}

func download(index int, numDivide int, sizeDivide int64, sizeTotal int64, url string, chStr chan<- string) {
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
	fmt.Printf("%v ", tmpfile.Name())
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if err = tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
	chStr <- tmpfile.Name()
}
