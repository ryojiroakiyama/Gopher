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

// 30秒
const ShortDuration = 5 * time.Second

// write input from sc to ch
// return channel that announce the end
func nextLine(ctx context.Context, sc *bufio.Scanner, ch chan<- string) <-chan struct{} {
	quit := make(chan struct{})
	go func() {
		defer close(quit)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				switch {
				case sc.Scan():
					ch <- sc.Text()
				case sc.Err() == nil:
					return
				default:
					fmt.Fprintln(os.Stderr, "extLine:", sc.Err())
					return
				}
			}
		}
	}()
	return quit
}

func runGame(ctx context.Context) int32 {
	sc := bufio.NewScanner(os.Stdin)
	ch := make(chan string)
	defer close(ch)
	quit := nextLine(ctx, sc, ch)
	var score int32
Loop:
	for {
		word := randomwords.Out()
		fmt.Printf("  %v\n", word)
		fmt.Printf("> ")
		select {
		case get := <-ch:
			if word == get {
				fmt.Printf("%sgood!%s\n", GREEN, RESET)
				score++
			} else {
				fmt.Printf("%sno..%s\n", RED, RESET)
			}
		case <-time.After(5 * time.Second):
			fmt.Printf("%stime-out%s\n", RED, RESET)
		case <-quit:
			os.Exit(0)
		case <-ctx.Done():
			break Loop
		}
	}
	return score
}

func main() {
	if err := randomwords.Init(); err != nil {
		fmt.Fprintln(os.Stderr, "init:", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), ShortDuration)
	defer cancel()
	score := runGame(ctx)
	fmt.Println()
	fmt.Println("Time's up! Score:", score)
}
