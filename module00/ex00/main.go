/*
** package main
 */
package main

import (
	"fmt"
	"github.com/ryojiroakiyama/convert"
	"os"
)

func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("error:", handler)
		}
	}()
	if len(os.Args) != 2 {
		panic("invalid argument")
	}
	dir := os.Args[1]
	convert.JpgToPng(dir)
}
