package converter

import (
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
)

const (
	JPG = "jpg"
	PNG = "png"
	GIF = "gif"
)

type conversion struct {
	srcExtension string
	dstExtension string
	encoder      images
	decoder      images
}

type images interface {
	decode(r io.Reader) (image.Image, error)
	encode(w io.Writer, m image.Image) error
}

type imageJpg struct{}

func (i imageJpg) decode(r io.Reader) (image.Image, error) {
	return jpeg.Decode(r)
}

func (i imageJpg) encode(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

type imagePng struct{}

func (i imagePng) decode(r io.Reader) (image.Image, error) {
	return png.Decode(r)
}

func (i imagePng) encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}

type imageGif struct{}

func (i imageGif) decode(r io.Reader) (image.Image, error) {
	return gif.Decode(r)
}

func (i imageGif) encode(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
