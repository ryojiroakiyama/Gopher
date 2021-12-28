package conversion

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

const (
	JPG     = "jpg"
	PNG     = "png"
	GIF     = "gif"
	JPGTYPE = "image/jpeg"
	PNGTYPE = "image/png"
	GIFTYPE = "image/gif"
)

func NewConverter(srcFormat string, dstFormat string) (*Converter, error) {
	decodeAction, err := getAction(srcFormat)
	if err != nil {
		return nil, err
	}
	encodeAction, err := getAction(dstFormat)
	if err != nil {
		return nil, err
	}
	return &Converter{
		SrcExtension: srcFormat,
		DstExtension: dstFormat,
		Encoder:      encodeAction,
		Decoder:      decodeAction,
	}, nil
}

func getAction(extension string) (Action, error) {
	switch extension {
	case JPG:
		return ActionJpg{}, nil
	case PNG:
		return ActionPng{}, nil
	case GIF:
		return ActionGif{}, nil
	default:
		return nil, fmt.Errorf("invalid argment: %v", extension)
	}
}

type Converter struct {
	SrcExtension string
	DstExtension string
	Decoder      Action
	Encoder      Action
}

//func (c Converter) GetSrcExt() string {
//	return c.SrcExtension
//}

//func (c Converter) GetDstExt() string {
//	return c.DstExtension
//}

//func (c Converter) GetEncoder() Action {
//	return c.Encoder
//}

//func (c Converter) GetDecoder() Action {
//	return c.Decoder
//}

type Action interface {
	decode(r io.Reader) (image.Image, error)
	encode(w io.Writer, m image.Image) error
}

type ActionJpg struct{}

func (i ActionJpg) decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func (i ActionJpg) encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

type ActionPng struct{}

func (i ActionPng) decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (i ActionPng) encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

type ActionGif struct{}

func (i ActionGif) decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (i ActionGif) encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
