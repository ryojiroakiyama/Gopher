package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func ToTmpFile(src io.Reader) (fileName string, err error) {
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		return "", fmt.Errorf("failt to create tmpfile: %v", err)
	}
	fileName = tmpfile.Name()
	fmt.Println("create:", fileName)
	defer func() {
		if cerr := tmpfile.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()
	_, err = io.Copy(tmpfile, src)
	if err != nil {
		return "", fmt.Errorf("fail to copy: %v", err)
	}
	return
}

func BindFiles(srcNames []string, dstName string) (err error) {
	dstfile, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("failt to create: %v", err)
	}
	defer func() {
		if cerr := dstfile.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
		if err != nil {
			os.Remove(dstName)
		}
	}()
	for _, srcName := range srcNames {
		srcfile, err := os.Open(srcName)
		if err != nil {
			return fmt.Errorf("failt to open: %v", err)
		}
		defer srcfile.Close()
		_, err = io.Copy(dstfile, srcfile)
		if err != nil {
			return fmt.Errorf("fail to copy: %v", err)
		}
	}
	return err
}

func FileBytes(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("fail to open: %v", err)
	}
	defer file.Close()
	srcBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("fail to read: %v", err)
	}
	return srcBytes, nil
}

func ToFile(fileName string, contents []byte) (err error) {
	dstFile, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("fail to create: %v", err)
	}
	defer func() {
		if cerr := dstFile.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()
	_, werr := dstFile.Write(contents)
	if werr != nil {
		return fmt.Errorf("fail to write: %v", werr)
	}
	return
}
