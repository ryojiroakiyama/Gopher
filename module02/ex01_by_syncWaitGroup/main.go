package main

import (
	"download/pget"
	"errors"
	"fmt"
	"os"
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
	if err := pget.Do("norm.pdf", url); err != nil {
		panic(err)
	}
}
