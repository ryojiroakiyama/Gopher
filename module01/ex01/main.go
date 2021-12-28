//Package main apply conversion process to all files
//in the directory passed as argument.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/ryojiroakiyama/convert/imgconv"
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

func argument() (string, string, string, bool) {
	if len(os.Args) == 1 {
		printError(errors.New("invalid argument"))
		return "", "", "", true
	} else if os.Args[1] == "-h" || os.Args[1] == "--help" {
		printHelp()
		return "", "", "", true
	}
	var informat = flag.String("i", "jpg", "file format convert from")
	var outformat = flag.String("o", "png", "file format convert to")
	flag.Parse()
	var dir = flag.Arg(0)
	return dir, *informat, *outformat, false
}

//main read all the files in the dir and call Do method.
//If fail to read dir, output the error and do nothing.
//Else if something happen, output a message about the thing
//and go to read the next file.
func main() {
	rootdir, srcExtension, dstExtension, is_end := argument()
	if is_end {
		return
	}
	c, err := conversion.NewConverter(srcExtension, dstExtension)
	if err != nil {
		printError(err)
		return
	}
	err = filepath.WalkDir(rootdir, func(path string, d fs.DirEntry, werr error) error {
		if werr != nil {
			if path == rootdir {
				return werr
			}
			printError(werr)
		} else if d.Type().IsRegular() {
			if cerr := c.Do(path); cerr != nil {
				printError(cerr)
			}
		}
		return nil
	})
	if err != nil {
		printError(err)
	}
}
