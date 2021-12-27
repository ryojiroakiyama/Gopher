//Package converter provides ability to convert file format
//according to the specified format.
//This process is applied to all files
//in the directory passed as argument.
package converter

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func printError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
}

//Do read all the files in the dir,
//and convert the file format from srcExt to dstExt
//if the file is srcExt format.
//If fail to read dir, regard it as an error and do nothing.
//Else	if something happen, output a message about what happened 
//and go to read the next file.
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
	case JPG:
		return imageJpg{}, nil
	case PNG:
		return imagePng{}, nil
	case GIF:
		return imageGif{}, nil
	default:
		return nil, fmt.Errorf("invalid argment: %v", Extension)
	}
}

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
