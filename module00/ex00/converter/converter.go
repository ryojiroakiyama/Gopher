package converter

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

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

type converter interface {
	convert(string) error
}

func JpgToPng(dir string) error {
	c := converterJpgToPng{}
	return applyEachFile(dir, c)
}

func applyEachFile(rootdir string, c converter) error {
	err := filepath.WalkDir(rootdir, func(path string, d fs.DirEntry, werr error) error {
		if werr != nil {
			if path == rootdir {
				return werr
			}
			printError(werr)
		} else if d.Type().IsRegular() {
			if cerr := c.convert(path); cerr != nil {
				printError(cerr)
			}
		}
		return nil
	})
	return err
}

type converterJpgToPng struct{}

func (j converterJpgToPng) convert(srcFileName string) error {
	srcBytes, err := getSrcBytes(srcFileName)
	if err != nil {
		return err
	}
	dstBytes, err := toPng(srcBytes)
	if err != nil {
		return fmt.Errorf("%v %v", srcFileName, err)
	}
	err = makeDstFile(strings.TrimSuffix(srcFileName, ".jpg")+".png", dstBytes)
	if err != nil {
		return err
	}
	return nil
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

func makeDstFile(fileName string, contents []byte) (err error) {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("fail to create: %v", err)
	}
	defer func() {
		if cerr := dstFile.Close(); cerr != nil {
			err = cerr
		}
	}()
	_, werr := dstFile.Write(contents)
	if werr != nil {
		return fmt.Errorf("fail to write: %v", werr)
	}
	return
}

func toPng(srcBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(srcBytes)

	switch contentType {
	case "image/png":
		return nil, fmt.Errorf("is already a png file")
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
