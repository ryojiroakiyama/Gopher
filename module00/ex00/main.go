/*
** package main
 */
package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ryojiroakiyama/convertimage/converter"
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

func main() {
	if len(os.Args) != 2 {
		printError(errors.New("invalid argument"))
		return
	}
	dir := os.Args[1]
	if err := converter.JpgToPng(dir); err != nil {
		printError(err)
	}
}
