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

func from_file(fileName string) error {
	if fi, err := os.Stat(fileName); err != nil {
		return err
	} else if fi.IsDir() {
		return fmt.Errorf("Is directory")
	}
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		return err
	}
	if err = cat(file, os.Stdout); err != nil {
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
	for _, file := range os.Args[1:] {
		if err := from_file(file); err != nil {
			fmt.Fprintln(os.Stderr, "cat:", err)
			status = 1
		}
	}
	if status != 0 {
		os.Exit(status)
	}
}
