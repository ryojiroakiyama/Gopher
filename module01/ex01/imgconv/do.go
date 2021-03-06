//Package conversion provides ability to convert file format
//according to the specified format.
package conversion

import (
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

//Do convert srcFileName file format from SrcExtension to DstExtension
//and create the file.
func (c Converter) Do(srcFileName string) error {
	srcBytes, err := getBytes(srcFileName)
	if err != nil {
		return err
	}
	dstBytes, err := c.convertFormat(srcBytes)
	if err != nil {
		return fmt.Errorf("%v %v", srcFileName, err)
	}
	if dstBytes != nil {
		if err = makeFile(c.dstFileName(srcFileName), dstBytes); err != nil {
			return err
		}
	}
	return nil
}

func (c Converter) dstFileName(srcFileName string) string {
	return strings.TrimSuffix(srcFileName, "."+c.SrcExtension) + "." + c.DstExtension
}

func (c Converter) convertFormat(srcBytes []byte) ([]byte, error) {
	contentType := http.DetectContentType(srcBytes)

	switch contentType {
	case getTypeName(c.SrcExtension):
		img, err := c.Decoder.decode(bytes.NewReader(srcBytes))
		if err != nil {
			return nil, fmt.Errorf("fail to decode from %s: %v", c.SrcExtension, err)
		}
		buf := new(bytes.Buffer)
		if err := c.Encoder.encode(buf, img); err != nil {
			return nil, fmt.Errorf("fail to encode to %s: %v", c.DstExtension, err)
		}
		return buf.Bytes(), nil
	case JPGTYPE, PNGTYPE, GIFTYPE:
		return nil, nil
	default:
		return nil, fmt.Errorf("is not a valid file")
	}
}

func getTypeName(extesion string) string {
	switch extesion {
	case JPG:
		return JPGTYPE
	case PNG:
		return PNGTYPE
	case GIF:
		return GIFTYPE
	default:
		return ""
	}
}
