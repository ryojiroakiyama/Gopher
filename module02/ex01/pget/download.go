package pget

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	ONEDLMAX = 1000
)

func Do(filepath string, url string) (err error) {
	sizeTotal, err := DataLength(url)
	if err != nil {
		return err
	}
	numDivide := NumDivideRange(sizeTotal)
	sizeDivide := sizeTotal / int64(numDivide)
	tmpFileNames := make([]string, numDivide)
	defer func() {
		for _, t := range tmpFileNames {
			if t != "" {
				fmt.Println("tmpfile:", t)
				os.Remove(t)
			}
		}
	}()
	for i := 0; i < numDivide; i++ {
		minRange := sizeDivide * int64(i)
		maxRange := sizeDivide * int64(i+1)
		if i == numDivide-1 {
			maxRange += sizeTotal - maxRange
		}
		fmt.Printf("i=%v, min=%v, max=%v\n", i, minRange, maxRange-1)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Range", RangeValue(minRange, maxRange-1))
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		tmpfile, err := os.CreateTemp("", "")
		if err != nil {
			return err
		}
		tmpFileNames[i] = tmpfile.Name()
		_, err = io.Copy(tmpfile, resp.Body)
		if err != nil {
			return err
		}
		if err = tmpfile.Close(); err != nil {
			return err
		}
	}
	dstFile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer dstFile.Close()
	for _, srcFileName := range tmpFileNames {
		srcfile, err := os.Open(srcFileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(dstFile, srcfile)
		if err != nil {
			return err
		}
	}
	return err
}
