package main

import (
	"fmt"
	"io/fs"
	//"os"
	"path/filepath"
)

func main() {
	//tmpDir, err := prepareTestDirTree("dir/to/walk/skip")
	//if err != nil {
	//	fmt.Printf("unable to create test dir tree: %v\n", err)
	//	return
	//}
	//defer os.RemoveAll(tmpDir)
	//os.Chdir(tmpDir)

	//subDirToSkip := "skip"

	fmt.Println("On Unix:")
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		//if info.IsDir() && info.Name() == subDirToSkip {
		//	fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
		//	return filepath.SkipDir
		//}
		fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		//fmt.Printf("error walking the path %q: %v\n", tmpDir, err)
		fmt.Printf("error walking the path: %v\n", err)
		return
	}
}
