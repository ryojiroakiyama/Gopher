package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	BLACK     = "\033[30m"
	RED       = "\033[31m"
	GREEN     = "\033[32m"
	YELLOW    = "\033[33m"
	BLUE      = "\033[34m"
	MAGENTA   = "\033[35m"
	CYAN      = "\033[36m"
	WHITE     = "\033[37m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	RESET     = "\033[0m"
	ONEDLMAX  = 1000
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

func printHelp() {
	fmt.Println(CYAN, "usage:", RESET)
	fmt.Println("  ./download <URL>")
}

func main() {
	if len(os.Args) == 1 {
		printError(errors.New("invalid argument"))
		return
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printHelp()
		return
	}
	url := os.Args[1]
	if err := DownloadFile("norm.pdf", url); err != nil {
		panic(err)
	}
}

func getRangeValue(start int64, end int64) string {
	return "bytes=" + strconv.FormatInt(start, 10) + "-" + strconv.FormatInt(end, 10)
}

func DownloadFile(filepath string, url string) (err error) {

	// get all length by reponse header
	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	length := resp.ContentLength
	if length <= 0 {
		return fmt.Errorf("unknown content length")
	}

	numDivide := getNumDivide(length)
	sizeDivide := length / int64(numDivide)
	allBody := &bytes.Buffer{}
	for i := 0; i < numDivide; i++ {
		minRange := sizeDivide * int64(i)
		maxRange := sizeDivide * int64(i+1)
		if i == numDivide-1 {
			maxRange += length - maxRange
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
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		allBody.Write(bodyBytes)
	}

	// write into outfile
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()
	_, err = io.Copy(out, allBody)
	return err
}

func getNumDivide(datasize int64) int {
	if datasize < ONEDLMAX {
		return 1
	}
	return 1 + getNumDivide(datasize/ONEDLMAX)
}
