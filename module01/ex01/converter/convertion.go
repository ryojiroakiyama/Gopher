package converter

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

func getDstFileName(srcFileName string, c conversion) string {
	return strings.TrimSuffix(srcFileName, "."+c.srcExtension) + "." + c.dstExtension
}

func convert(srcFileName string, c conversion) error {
	srcBytes, err := getSrcBytes(srcFileName)
	if err != nil {
		return err
	}
	dstBytes, err := convertionImage(srcBytes, c)
	if err != nil {
		return fmt.Errorf("%v %v", srcFileName, err)
	}
	err = makeDstFile(getDstFileName(srcFileName, c), dstBytes)
	if err != nil {
		return err
	}
	return nil
}

func getTypeName(extesion string) string {
	switch extesion {
	case JPG:
		return "image/jpeg"
	default:
		return "image/" + extesion
	}
}

func convertionImage(srcBytes []byte, c conversion) ([]byte, error) {
	contentType := http.DetectContentType(srcBytes)
	srcType := getTypeName(c.srcExtension)
	dstType := getTypeName(c.dstExtension)

	switch contentType {
	case dstType:
		return nil, fmt.Errorf("is already a %s format file", dstType)
	case srcType:
		img, err := c.decoder.decode(bytes.NewReader(srcBytes))
		if err != nil {
			return nil, fmt.Errorf("unable to decode from %s: %v", srcType, err)
		}
		buf := new(bytes.Buffer)
		if err := c.encoder.encode(buf, img); err != nil {
			return nil, fmt.Errorf("unable to encode to %s: %v", dstType, err)
		}
		return buf.Bytes(), nil
	}
	return nil, fmt.Errorf("is not a valid file")
}
