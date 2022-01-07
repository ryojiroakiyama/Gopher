package main

import (
	"fmt"
	"io"
	"os"
)

func cat(in io.Reader, out io.Writer) error {
	buf := make([]byte, 101)
	for {
		n, err := in.Read(buf)
		if n > 0 {
			if _, err := out.Write(buf[:n]); err != nil {
				return err
			}
		}
		if n == 0 || err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var in io.Reader
	var err error
	args := len(os.Args)
	for idx, file := range os.Args {
		switch {
		case args == 1:
			in = os.Stdin
		case idx == 0:
			continue
		default:
			if fi, err := os.Stat(file); err != nil {
				fmt.Println(err)
				os.Exit(1)
			} else if fi.IsDir() {
				fmt.Println("Is a directory")
				os.Exit(1) //exit status -> 0 and continue
			}
			in, err = os.Open(file)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if err = cat(in, os.Stdout); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
