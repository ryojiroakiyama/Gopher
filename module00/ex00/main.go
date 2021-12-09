/*
** package main
 */
package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"net/http"
	"os"
	//"path/filepath"
	"strings"
)

func main() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("error:", handler)
		}
	}()
	if len(os.Args) != 2 {
		panic("invalid argument")
	}
	dir := os.Args[1]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		Convert(file.Name())
	}
	//result, err := FilePathWalkDir(dir)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//for _, file := range result {
	//	fmt.Println(file)
	//}
}

//func FilePathWalkDir(root string) ([]string, error) {
//	var files []string
//	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//		if !info.IsDir() {
//			files = append(files, path)
//		}
//		return nil
//	})
//	return files, err
//}

func Convert(srcFileName string) {
	fmt.Println("srcFileName: ", srcFileName)
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	srcByte, err := ioutil.ReadAll(srcFile)
	if err != nil {
		panic(err)
	}
	dstByte, err := ToPng(srcByte)
	if err != nil {
		panic(err)
	}
	dstFile, err := os.Create(strings.TrimSuffix(srcFileName, ".jpg") + ".png")
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	_, er := dstFile.Write(dstByte)
	if er != nil {
		panic(err)
	}
}

func ToPng(imageBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			return nil, fmt.Errorf("%v: unable to decode jpeg", err)
		}

		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return nil, fmt.Errorf("%v: unable to encode png", err)
		}
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("unable to convert %#v to png", contentType)
}
