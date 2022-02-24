package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"omikuji/mikujitou"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "invalid arg num")
		return
	}
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := mikujitou.Get
	http.HandleFunc("/", h1)
	http.HandleFunc("/omikuji", h2)

	log.Fatal(http.ListenAndServe(":"+os.Args[1], nil))
}
