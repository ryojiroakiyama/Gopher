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

func from_arg(arg string) error {
	var in io.Reader
	if arg == "-" {
		in = os.Stdin
	} else {
		if fi, err := os.Stat(arg); err != nil {
			return err
		} else if fi.IsDir() {
			return fmt.Errorf("Is directory")
		}
		file, err := os.Open(arg)
		defer file.Close()
		if err != nil {
			return err
		}
		in = file
	}
	if err := cat(in, os.Stdout); err != nil {
		return err
	}
	return nil
}

func main() {
	var status int
	if len(os.Args) == 1 {
		if err := cat(os.Stdin, os.Stdout); err != nil {
			fmt.Fprintln(os.Stderr, "cat:", err)
			status = 1
		}
	}
	for _, arg := range os.Args[1:] {
		if err := from_arg(arg); err != nil {
			fmt.Fprintln(os.Stderr, "cat:", err)
			status = 1
		}
	}
	if status != 0 {
		os.Exit(status)
	}
}
