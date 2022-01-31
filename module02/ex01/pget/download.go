package pget

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	ONEDLMAX = 1000
)

func Do(filepath string, url string) (err error) {

	dataLen, err := DataLength(url)
	if err != nil {
		return err
	}
	numDivide := NumDivideRange(dataLen)
	sizeDivide := dataLen / int64(numDivide)
	var files []string
	for i := 0; i < numDivide; i++ {
		minRange := sizeDivide * int64(i)
		maxRange := sizeDivide * int64(i+1)
		if i == numDivide-1 {
			maxRange += dataLen - maxRange
		}
		fmt.Printf("i=%v, min=%v, max=%v\n", i, minRange, maxRange-1)
		rangeValue := getRangeValue(minRange, maxRange-1)
		client := &http.Client{}
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Range", rangeValue)
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		files = append(files, strconv.Itoa(i))
		out, err := os.Create(files[i])
		if err != nil {
			return err
		}
		defer func() {
			if cerr := out.Close(); cerr != nil {
				err = fmt.Errorf("fail to close: %v", cerr)
			}
		}()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
	}
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()
	for _, fileName := range files {
		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		defer os.Remove(fileName)
		_, err = io.Copy(out, file)
		if err != nil {
			return err
		}
	}
	return err
}
