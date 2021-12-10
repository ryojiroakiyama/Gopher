package convert

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

type converter interface {
	convert(string)
}

func JpgToPng(dir string) {
	c := converterJpgToPng{}
	applyEachFile(dir, c)
}

func applyEachFile(dir string, c converter) {
	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
		}
		if d.Type().IsRegular() {
			c.convert(path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("error walking the path: %v\n", err)
	}
}

type converterJpgToPng struct{}

func (j converterJpgToPng) convert(srcFileName string) {
	srcBytes, err := getSrcBytes(srcFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	dstBytes, err := toPng(srcBytes)
	if err != nil {
		fmt.Println("error:", srcFileName, err)
		return
	}
	err = makeDstFile(strings.TrimSuffix(srcFileName, ".jpg")+".png", dstBytes)
	if err != nil {
		fmt.Println(err)
	}
}

func getSrcBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("fail to open: %v", err)
	}
	defer file.Close()
	srcBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("fail to raadall: %v", err)
	}
	return srcBytes, nil
}

func makeDstFile(fileName string, contents []byte) error {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("fail to create: %v", err)
	}
	defer dstFile.Close()
	_, er := dstFile.Write(contents)
	if er != nil {
		return fmt.Errorf("fail to write: %v", err)
	}
	return nil
}

func toPng(srcBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(srcBytes)

	switch contentType {
	case "image/png":
		return nil, fmt.Errorf("is a png file")
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(srcBytes))
		if err != nil {
			return nil, fmt.Errorf(": unable to decode jpeg: %v", err)
		}
		buf := new(bytes.Buffer)
		if err := png.Encode(buf, img); err != nil {
			return nil, fmt.Errorf(": unable to encode png: %v", err)
		}
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("is not a valid file")
}
