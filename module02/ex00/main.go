package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"typing/randomwords"
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
				fmt.Println("good!")
			} else {
				fmt.Println("bad")
			}
		case <-time.After(5 * time.Second):
			fmt.Println("\ntimed out (>_<;)")
			return
		}
	}
}
