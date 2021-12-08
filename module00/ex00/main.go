package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("invalid number of arguments")
		return
	}
	fmt.Println(os.Args[1])
}
