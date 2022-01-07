package main

import (
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
	cat(os.Stdin, os.Stdout)
}
