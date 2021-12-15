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

type converter interface {
	convert(string) error
}

func applyEachFile(rootdir string, c converter) error {
	err := filepath.WalkDir(rootdir, func(path string, d fs.DirEntry, werr error) error {
		if werr != nil {
			if path == rootdir {
				return werr
			}
			printError(werr)
		} else if d.Type().IsRegular() {
			if cerr := c.convert(path); cerr != nil {
				printError(cerr)
			}
		}
		return nil
	})
	return err
}
