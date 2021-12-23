package converter

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func convert(srcFileName string, c conversion) error {
	srcBytes, err := getSrcBytes(srcFileName)
	if err != nil {
		return err
	}
	dstBytes, err := toPng(srcBytes, c)
	if err != nil {
		return fmt.Errorf("%v %v", srcFileName, err)
	}
	err = makeDstFile(strings.TrimSuffix(srcFileName, "."+c.srcExtension)+"."+c.dstExtension, dstBytes)
	if err != nil {
		return err
	}
	return nil
}

func toPng(srcBytes []byte, c conversion) ([]byte, error) {
	contentType := http.DetectContentType(srcBytes)

	switch contentType {
	case "image/" + c.dstExtension:
		return nil, fmt.Errorf("is already a %s format file", c.dstExtension)
	case "image/" + c.srcExtension:
		//img, err := jpeg.Decode(bytes.NewReader(srcBytes))
		img, err := c.decoder.decode(bytes.NewReader(srcBytes))
		if err != nil {
			return nil, fmt.Errorf("unable to decode from %s: %v", c.srcExtension, err)
		}
		buf := new(bytes.Buffer)
		//if err := png.Encode(buf, img); err != nil {
		if err := c.encoder.encode(buf, img); err != nil {
			return nil, fmt.Errorf("unable to encode to %s: %v", c.dstExtension, err)
		}
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("is not a valid file")
}
