package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	rand.Seed(time.Now().Unix())
	words := [3]string{"apple", "water", "42tokyo"}
	for idx := 0; true; idx++ {
		if idx == 3 {
			idx = 0
		}
		word := words[rand.Intn(len(words))]
		fmt.Printf("  %v\n", word)
		fmt.Printf("> ")
		go nextLine(sc, ch)
		select {
		case get := <-ch:
			if word == get {
				fmt.Println("Congratulations!")
				return
			}
		case <-time.After(5 * time.Second):
			fmt.Println("\ntimed out (>_<;)")
			return
		}
	}
}
