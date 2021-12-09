package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	fmt.Println("On Unix:")
	err := filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if d.Type().IsRegular() {
			fmt.Print("file!: ")
		}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}

	//fmt.Println("On Unix:")
	//err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
	//	if err != nil {
	//		fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
	//		return err
	//	}
	//	if info.Mode().IsRegular() {
	//		fmt.Print("file!: ")
	//	}
	//	fmt.Printf("visited file or dir: %q\n", path)
	//	return nil
	//})
	//if err != nil {
	//	fmt.Printf("error walking the path: %v\n", err)
	//	return
	//}
}
