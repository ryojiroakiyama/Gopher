//Package converter provides ability to convert
//from one file format to another format.
//This process is applied to all files
//in the directory passed as argument.
package converter

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type conversion struct {
	srcExtension string
	dstExtension string
	encoder      images
	decoder      images
}

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

func Do(dir string, srcExt string, dstExt string) error {
	c := conversion{
		srcExtension: srcExt,
		dstExtension: dstExt,
	}
	var err error
	c.decoder, err = getImages(srcExt)
	if err != nil {
		return err
	}
	c.encoder, err = getImages(dstExt)
	if err != nil {
		return err
	}
	return applyEachFile(dir, c)
}

func getImages(Extension string) (images, error) {
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
