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

const (
	BLACK          = "\033[30m"
	RED            = "\033[31m"
	GREEN          = "\033[32m"
	YELLOW         = "\033[33m"
	BLUE           = "\033[34m"
	MAGENTA        = "\033[35m"
	CYAN           = "\033[36m"
	WHITE          = "\033[37m"
	BOLD           = "\033[1m"
	UNDERLINE      = "\033[4m"
	BOLD_UNDERLINE = "\033[1;4m"
	RESET          = "\033[0m"
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

func printHelp() {
	fmt.Println(CYAN, "usage:", RESET)
	fmt.Println("  ./convert -i=png -o=jpg images_directory")
	fmt.Println(CYAN, "available formats:", RESET)
	fmt.Println("  jpg, png, gif")
	fmt.Println(CYAN, "options:", RESET)
	fmt.Println("  i: format convert from, default jpg")
	fmt.Println("  o: format convert to, default png")
}

func main() {
	if len(os.Args) == 1 {
		printError(errors.New("invalid argument"))
		return
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printHelp()
		return
	}
	var informat = flag.String("i", "jpg", "file format source to convert")
	var outformat = flag.String("o", "png", "file format destination")
	flag.Parse()
	var dir = flag.Arg(0)
	if err := converter.Do(dir, *informat, *outformat); err != nil {
		printError(err)
	}
}
