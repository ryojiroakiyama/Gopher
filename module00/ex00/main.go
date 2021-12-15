/*
** package main
 */
package main

import (
	"fmt"
	"os"

	"github.com/ryojiroakiyama/convertimage/converter"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "invalid argument")
		return
	}
	dir := os.Args[1]
	if err := converter.JpgToPng(dir); err != nil {
		fmt.Fprintln(os.Stderr, "main:", err)
	}
}
