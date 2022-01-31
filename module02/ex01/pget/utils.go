package pget

import (
	"strconv"
)

func getRangeValue(start int64, end int64) string {
	return "bytes=" + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10)
}

func getNumDivide(datasize int64) int {
	if datasize < ONEDLMAX {
		return 1
	}
	return 1 + getNumDivide(datasize/ONEDLMAX)
}
