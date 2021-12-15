package converter

import (
	"fmt"
	"io/ioutil"
	"os"
)

func getSrcBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("fail to open: %v", err)
	}
	defer file.Close()
	srcBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("fail to raadall: %v", err)
	}
	return srcBytes, nil
}

func makeDstFile(fileName string, contents []byte) (err error) {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("fail to create: %v", err)
	}
	defer func() {
		if cerr := dstFile.Close(); cerr != nil {
			err = cerr
		}
	}()
	_, werr := dstFile.Write(contents)
	if werr != nil {
		return fmt.Errorf("fail to write: %v", werr)
	}
	return
}
