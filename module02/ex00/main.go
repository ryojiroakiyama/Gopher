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
	green = "\033[32m"
	red   = "\033[31m"
	cyan  = "\033[36m"
	reset = "\033[0m"
)

const (
	ShortDuration = 30 * time.Second
)

// write input from sc to ch
// return channel that announce the end
func Scan(ctx context.Context, sc *bufio.Scanner, ch chan<- string) <-chan struct{} {
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
					fmt.Fprintln(os.Stderr, "nextLine:", sc.Err())
					return
				}
			}
		}
	}()
	return quit
}

func runGame(ctx context.Context) uint16 {
	// set up
	sc := bufio.NewScanner(os.Stdin)
	ch := make(chan string)
	defer close(ch)
	scanQuit := Scan(ctx, sc, ch)

	var score uint16
Loop:
	for {
		word := randomwords.Out()
		fmt.Printf("  %v\n", word)
		fmt.Printf("> ")
		select {
		case get := <-ch:
			if word == get {
				fmt.Printf("%sgood!%s\n", green, reset)
				score++
			} else {
				fmt.Printf("%sno..%s\n", red, reset)
			}
		case <-time.After(5 * time.Second):
			fmt.Printf("%stime-out%s\n", red, reset)
		case <-scanQuit:
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
