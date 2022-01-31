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

	datasize, err := DataLength(url)
	if err != nil {
		return err
	}
	numDivide := NumDivideRange(datasize)
	sizeDivide := datasize / int64(numDivide)
	for i := 0; i < numDivide; i++ {
		minRange := sizeDivide * int64(i)
		maxRange := sizeDivide * int64(i+1)
		if i == numDivide-1 {
			maxRange += datasize - maxRange
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
		if err = toFile(strconv.Itoa(i), resp.Body); err != nil {
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
	for i := 0; i < numDivide; i++ {
		srcFileName := strconv.Itoa(i)
		file, err := os.Open(srcFileName)
		if err != nil {
			return err
		}
		defer file.Close()
		defer os.Remove(srcFileName)
		_, err = io.Copy(out, file)
		if err != nil {
			return err
		}
	}
	return err
}
