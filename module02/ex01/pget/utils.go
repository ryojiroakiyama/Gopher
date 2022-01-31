package pget

import (
	"fmt"
	"net/http"
	"strconv"
)

func RangeValue(start int64, end int64) string {
	return "bytes=" + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10)
}

func NumDivideRange(datasize int64) int {
	if datasize < ONEDLMAX {
		return 1
	}
	return 1 + NumDivideRange(datasize/ONEDLMAX)
}

func DataLength(url string) (int64, error) {
	resp, err := http.Head(url)
	if err != nil {
		return 0, err
	}
	length := resp.ContentLength
	if length <= 0 {
		return 0, fmt.Errorf("unknown content length")
	}
	return length, nil
}
