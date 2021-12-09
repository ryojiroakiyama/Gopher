/*
** package main
 */
package main

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
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
	ConvertAll(dir)
}

func ConvertAll(dir string) {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}
		if d.Type().IsRegular() {
			//fmt.Print("file!: ")
			jpg_to_png(path)
		}
		//fmt.Printf("visited file or dir: %q\n", path)
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
		return
	}
}

func jpg_to_png(srcFileName string) {
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
	dstFile, err := os.Create(strings.TrimSuffix(srcFileName, ".jpg") + ".png")
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	err = ToPng(srcByte, dstFile)
	if err != nil {
		panic(err)
	}
	//dstByte, err := ToPng(srcByte)
	//if err != nil {
	//	panic(err)
	//}
	//dstFile, err := os.Create(strings.TrimSuffix(srcFileName, ".jpg") + ".png")
	//if err != nil {
	//	panic(err)
	//}
	//defer dstFile.Close()
	//_, er := dstFile.Write(dstByte)
	//if er != nil {
	//	panic(err)
	//}
}

func ToPng(imageBytes []byte, dstfile *os.File) error {
	contentType := http.DetectContentType(imageBytes)

	switch contentType {
	case "image/png":
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			return fmt.Errorf("%v: unable to decode jpeg", err)
		}
		if err := png.Encode(dstfile, img); err != nil {
			return fmt.Errorf("%v: unable to encode png", err)
		}
		return nil
	}
	return fmt.Errorf("unable to convert %#v to png", contentType)
}

//func ToPng(imageBytes []byte) ([]byte, error) {
//	contentType := http.DetectContentType(imageBytes)

//	switch contentType {
//	case "image/png":
//	case "image/jpeg":
//		img, err := jpeg.Decode(bytes.NewReader(imageBytes))
//		if err != nil {
//			return nil, fmt.Errorf("%v: unable to decode jpeg", err)
//		}

//		buf := new(bytes.Buffer)
//		if err := png.Encode(buf, img); err != nil {
//			return nil, fmt.Errorf("%v: unable to encode png", err)
//		}
//		return buf.Bytes(), nil
//	}
//	return nil, fmt.Errorf("unable to convert %#v to png", contentType)
//}
