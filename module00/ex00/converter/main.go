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

/*
** If fail to read root dir:
**   -> Error. Do Nothing.
**
** Else	if something happen:
**   -> Output what happens, and go to read the next file.
 */
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
