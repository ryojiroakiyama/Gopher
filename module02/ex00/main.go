package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"typing/randomwords"
)

const (
	GREEN = "\033[32m"
	RED   = "\033[31m"
	CYAN  = "\033[36m"
	RESET = "\033[0m"
)

func nextLine(sc *bufio.Scanner, ch chan<- string) {
	switch {
	case sc.Scan():
		ch <- sc.Text()
	case sc.Err() == nil:
		os.Exit(0)
	default:
		panic(sc.Err())
	}
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	ch := make(chan string)
	defer close(ch)
	if err := randomwords.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "init:", err)
		return
	}
	for {
		word := randomwords.Out()
		fmt.Printf("  %v\n", word)
		fmt.Printf("> ")
		go nextLine(sc, ch)
		select {
		case get := <-ch:
			if word == get {
				fmt.Printf("%sgood!%s\n", GREEN, RESET)
			} else {
				fmt.Printf("%sno..%s\n", RED, RESET)
			}
		case <-time.After(5 * time.Second):
			fmt.Printf("\n%stimed out (>_<;)%s\n", CYAN, RESET)
			return
		}
	}
}
