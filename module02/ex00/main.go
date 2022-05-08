package main

import (
	"bufio"
	"context"
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

const ShortDuration = 10 * time.Second

func nextLine(ctx context.Context, sc *bufio.Scanner, ch chan<- string) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			switch {
			case sc.Scan():
				ch <- sc.Text()
			case sc.Err() == nil:
				os.Exit(0)
			default:
				panic(sc.Err())
			}
		}
	}
}

func runGame(ctx context.Context, ch <-chan string) {
	for {
		word := randomwords.Out()
		fmt.Printf("  %v\n", word)
		fmt.Printf("> ")
		select {
		case get := <-ch:
			if word == get {
				fmt.Printf("%sgood!%s\n", GREEN, RESET)
			} else {
				fmt.Printf("%sno..%s\n", RED, RESET)
			}
		case <-time.After(5 * time.Second):
			fmt.Printf("%stime-out%s\n", RED, RESET)
		case <-ctx.Done():
			return
		}
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
	ctx, cancel := context.WithTimeout(context.Background(), ShortDuration)
	defer cancel()
	go nextLine(ctx, sc, ch)
	runGame(ctx, ch)
}
