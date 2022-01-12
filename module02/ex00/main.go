package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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
	for {
		word := "apple"
		fmt.Printf("  %v\n", word)
		fmt.Printf("> ")
		go nextLine(sc, ch)
		select {
		case get := <-ch:
			if word == get {
				return
			}
		case <-time.After(5 * time.Second):
			fmt.Println("timed out")
			return
		}
	}
}
