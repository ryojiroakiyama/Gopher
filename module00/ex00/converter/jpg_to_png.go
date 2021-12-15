package converter

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"
)

func JpgToPng(dir string) error {
	c := converterJpgToPng{}
	return applyEachFile(dir, c)
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
