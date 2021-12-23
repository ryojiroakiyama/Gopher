//Package converter provides ability to convert
//from one file format to another format.
//This process is applied to all files
//in the directory passed as argument.
package converter

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

func Do(dir string, srcExt string, dstExt string) error {
	c := conversion{
		srcExtension: srcExt,
		dstExtension: dstExt,
	}
	var err error
	c.decoder, err = getImage(srcExt)
	if err != nil {
		return err
	}
	c.encoder, err = getImage(dstExt)
	if err != nil {
		return err
	}
	fmt.Println("---->", c.srcExtension, c.dstExtension)
	//os.Exit(0)
	applyEachFile(dir, c)
	return nil
}

func getImage(Extension string) (images, error) {
	switch Extension {
	case "jpg":
		return imageJpg{}, nil
	case "png":
		return imagePng{}, nil
	case "gif":
		return imageGif{}, nil
	default:
		return nil, fmt.Errorf("invalid argment: %v", Extension)
	}
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

/*
** If fail to read root dir:
**   -> Error. Do Nothing.
**
** Else	if something happen:
**   -> Output what happens, and go to read the next file.
 */
func applyEachFile(rootdir string, c conversion) error {
	err := filepath.WalkDir(rootdir, func(path string, d fs.DirEntry, werr error) error {
		if werr != nil {
			if path == rootdir {
				return werr
			}
			printError(werr)
		} else if d.Type().IsRegular() {
			if cerr := convert(path, c); cerr != nil {
				printError(cerr)
			}
		}
		return nil
	})
	return err
}
