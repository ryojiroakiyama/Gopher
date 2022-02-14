package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ryojiroakiyama/pget/pget"
)

const (
	black     = "\033[30m"
	red       = "\033[31m"
	green     = "\033[32m"
	yellow    = "\033[33m"
	blue      = "\033[34m"
	magenta   = "\033[35m"
	cyan      = "\033[36m"
	white     = "\033[37m"
	bold      = "\033[1m"
	underline = "\033[4m"
	reset     = "\033[0m"
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

func printHelp() {
	fmt.Println(cyan, "usage:", reset)
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
	if err := pget.Do(url); err != nil {
		panic(err)
	}
}
