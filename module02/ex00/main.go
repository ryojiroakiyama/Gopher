package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func nextLine() (string, error) {
	switch {
	case sc.Scan():
		return sc.Text(), nil
	case sc.Err() == nil:
		return "", fmt.Errorf("eof")
	default:
		fmt.Fprintln(os.Stderr, sc.Err())
		return "", fmt.Errorf("error")
	}
}

func main() {
	for {
		word := "apple"
		fmt.Printf("  %v\n", word)
		fmt.Printf("> ")
		if get, err := nextLine(); err != nil {
			break
		} else if word == get {
			break
		}
	}
}
