package pget

import (
	"fmt"
	"net/http"
	"os"
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
	tmpfiles := make([]tmpfile, numDivide)
	defer func() {
		for _, t := range tmpfiles {
			t.remove()
		}
	}()
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
		if err = toFile("", &tmpfiles[i], resp.Body); err != nil {
			return err
		}
	}
	var dfile dstfile
	for _, tmpfile := range tmpfiles {
		file, err := os.Open(tmpfile.name)
		if err != nil {
			return err
		}
		if err = toFile(filepath, &dfile, file); err != nil {
			return err
		}
	}
	return err
}
