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

type Convert struct {
	Dir string
}

func (c Convert) JpgToPng() {
	err := filepath.WalkDir(c.Dir, func(path string, d fs.DirEntry, err error) error {
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
	//fmt.Println("srcFileName: ", srcFileName)
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println(err)
	}
	defer srcFile.Close()
	srcBytes, err := ioutil.ReadAll(srcFile)
	if err != nil {
		fmt.Println(err)
	}
	dstBytes, err := toPng(srcBytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	dstFile, err := os.Create(strings.TrimSuffix(srcFileName, ".jpg") + ".png")
	if err != nil {
		fmt.Println(err)
	}
	defer dstFile.Close()
	_, er := dstFile.Write(dstBytes)
	if er != nil {
		fmt.Println(err)
	}
}

func toPng(srcBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(srcBytes)

	switch contentType {
	case "image/png":
	case "image/jpeg":
		img, err := jpeg.Decode(bytes.NewReader(srcBytes))
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
