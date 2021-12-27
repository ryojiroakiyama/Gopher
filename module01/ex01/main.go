/*
** package main
 */
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/ryojiroakiyama/convert/converter"
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

func main() {
	if len(os.Args) == 1 {
		printError(errors.New("invalid argument"))
		return
	}
	var informat = flag.String("i", "jpg", "file format source to convert")
	var outformat = flag.String("o", "png", "file format destination")
	flag.Parse()
	var dir = flag.Arg(0)
	fmt.Println(*informat, *outformat, dir)
	if err := converter.Do(dir, *informat, *outformat); err != nil {
		printError(err)
	}
}
