package main

import (
	"io"
	"os"
	"strings"
)

func putError(out io.Writer, message string) {
	if _, err := out.Write([]byte("ft_cat: " + message + "\n")); err != nil {
		os.Exit(1)
	}
}

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

func from_arg(arg string) string {
	var in io.Reader
	if arg == "-" {
		in = os.Stdin
	} else {
		file, err := os.Open(arg)
		if err != nil {
			return strings.TrimPrefix(err.Error(), "open ")
		}
		defer file.Close()
		if fileInfo, err := os.Stat(arg); err != nil {
			return strings.TrimPrefix(err.Error(), "stat ")
		} else if fileInfo.IsDir() {
			return "Is directory"
		}
		in = file
	}
	if err := cat(in, os.Stdout); err != nil {
		return err.Error()
	}
	return ""
}

func main() {
	var status int
	if len(os.Args) == 1 {
		if err := cat(os.Stdin, os.Stdout); err != nil {
			putError(os.Stderr, err.Error())
			status = 1
		}
	}
	for _, arg := range os.Args[1:] {
		if err_message := from_arg(arg); err_message != "" {
			putError(os.Stderr, err_message)
			status = 1
		}
	}
	if status != 0 {
		os.Exit(status)
	}
}
