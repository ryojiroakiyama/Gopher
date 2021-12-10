package main

import "fmt"

func sliceC(slice []byte) {
	slice[5] = 'C'
}

func arrayC(array [6]byte) {
	array[5] = 'C' // copy
}

func main() {
	sliceA := []byte{'s', 'l', 'i', 'c', 'e', 'A'} //slice
	fmt.Printf("sliceA:%s\n", sliceA)
	sliceB := sliceA
	sliceB[5] = 'B'
	fmt.Printf("sliceB:%s\n", sliceB)
	sliceC(sliceA)
	fmt.Printf("changed:%s\n", sliceA)

	arrayA := [...]byte{'s', 'l', 'i', 'c', 'e', 'A'} //array
	fmt.Printf("arrayA:%s\n", arrayA)
	arrayB := arrayA
	arrayB[5] = 'B'
	fmt.Printf("arrayB:%s\n", arrayB)
	arrayC(arrayA)
	fmt.Printf("changed:%s\n", arrayA)
}
