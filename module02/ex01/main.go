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
	BLACK          = "\033[30m"
	RED            = "\033[31m"
	GREEN          = "\033[32m"
	YELLOW         = "\033[33m"
	BLUE           = "\033[34m"
	MAGENTA        = "\033[35m"
	CYAN           = "\033[36m"
	WHITE          = "\033[37m"
	BOLD           = "\033[1m"
	UNDERLINE      = "\033[4m"
	BOLD_UNDERLINE = "\033[1;4m"
	RESET          = "\033[0m"
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

func rangeStr(start int, end int) string {
	return "bytes=" + strconv.Itoa(start) + "-" + strconv.Itoa(end)
}

func DownloadFile(filepath string, url string) (err error) {

	// get all length by reponse header
	resp_h, _ := http.Head(url)
	maps := resp_h.Header
	length, _ := strconv.Atoi(maps["Content-Length"][0])
	start1 := 0
	end1 := length / 2
	start2 := end1 + 1
	end2 := length - 1
	fmt.Printf("range1:%v-%v, range2:%v-%v, length:%v\n", start1, end1, start2, end2, length)

	// get first range response
	range1 := rangeStr(start1, end1)
	client1 := &http.Client{}
	req1, _ := http.NewRequest("GET", url, nil)
	req1.Header.Add("Range", range1)
	resp1, err := client1.Do(req1)
	if err != nil {
		return err
	}
	defer resp1.Body.Close()

	// get last range response
	range2 := rangeStr(start2, end2)
	client2 := &http.Client{}
	req2, _ := http.NewRequest("GET", url, nil)
	req2.Header.Add("Range", range2)
	resp2, err := client2.Do(req2)
	if err != nil {
		return err
	}
	defer resp2.Body.Close()

	// bind body
	allbody := &bytes.Buffer{}
	body_bytes, _ := io.ReadAll(resp1.Body)
	allbody.Write(body_bytes)
	body_bytes, _ = io.ReadAll(resp2.Body)
	allbody.Write(body_bytes)

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
	_, err = io.Copy(out, allbody)
	return err
}
