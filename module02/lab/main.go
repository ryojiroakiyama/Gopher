package main

import (
	"fmt"
	//"time"
)

var c chan int

func handle(int) {}

func main() {
	//select {
	//case m := <-c:
	//	handle(m)
	//case <-time.After(10 * time.Second):
	//	fmt.Println("timed out")
	//}

	a := make(map[string][]string)
	a["A"] = []string{"a", "b", "c"}
	fmt.Println(a["A"][0])
	fmt.Println(a["B"])
}
