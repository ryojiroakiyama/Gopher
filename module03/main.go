package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Omikuji struct {
	Fortune string `json:"fortune-omikuji`
}

func main() {
	o := Omikuji{Fortune: "kichi"}
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		j, _ := json.Marshal(o)
		w.Write(j)
		resp := make(map[string]string)
		resp["message"] = "Status Created"
		jsonResp, _ := json.Marshal(resp)
		w.Write(jsonResp)
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
