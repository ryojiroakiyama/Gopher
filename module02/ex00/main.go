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
	words := [3]string{"apple", "water", "42tokyo"}
	for idx := 0; true; idx++ {
		if idx == 3 {
			idx = 0
		}
		fmt.Printf("  %v\n", words[idx])
		fmt.Printf("> ")
		go nextLine(sc, ch)
		select {
		case get := <-ch:
			if words[idx] == get {
				fmt.Println("Congratulations!")
				return
			}
		case <-time.After(5 * time.Second):
			fmt.Println("\ntimed out (>_<;)")
			return
		}
	}
}
