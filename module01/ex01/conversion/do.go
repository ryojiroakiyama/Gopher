//Package conversion provides ability to convert file format
//according to the specified format.
//This process is applied to all files
//in the directory passed as argument.
package conversion

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

//Do read all the files in the dir,
//and convert the file format from srcExt to dstExt
//if the file is srcExt format.
//If fail to read dir, regard it as an error and do nothing.
//Else	if something happen, output a message about what happened
//and go to read the next file.
func (c Converter) Do(srcFileName string) error {
	srcBytes, err := getSrcBytes(srcFileName)
	if err != nil {
		return err
	}
	dstBytes, err := c.convertionImage(srcBytes)
	if err != nil {
		return fmt.Errorf("%v %v", srcFileName, err)
	}
	err = makeDstFile(c.getDstFileName(srcFileName), dstBytes)
	if err != nil {
		return err
	}
	return nil
}

func (c Converter) getDstFileName(srcFileName string) string {
	return strings.TrimSuffix(srcFileName, "."+c.srcExtension) + "." + c.dstExtension
}

func (c Converter) convertionImage(srcBytes []byte) ([]byte, error) {
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

func getTypeName(extesion string) string {
	switch extesion {
	case JPG:
		return "image/jpeg"
	default:
		return "image/" + extesion
	}
}
