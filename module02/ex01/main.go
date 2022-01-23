package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, akiyama in the house")
//}

//func main() {
//	http.HandleFunc("/", handler)
//	http.ListenAndServe(":8080", nil)
//}

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
	if err := DownloadFile("netpractice.pdf", url); err != nil {
		panic(err)
	}
}

func DownloadFile(filepath string, url string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := out.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()

	_, err = io.Copy(out, resp.Body)
	return err
}
