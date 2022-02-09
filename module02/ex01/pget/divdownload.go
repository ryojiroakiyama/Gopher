package pget

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func rangeRequest(url string, minRange int64, maxRange int64) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("fail to send request: %v", err)
	}
	req.Header.Add("Range", RangeValue(minRange, maxRange-1))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("fail to get response: %v", err)
	}
	return resp.Body, nil
}

func toTmpFile(src io.Reader) (fileName string, err error) {
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		return "", fmt.Errorf("failt to create: %v", err)
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

func divDownload(url string, minRange int64, maxRange int64) (string, error) {
	body, err := rangeRequest(url, minRange, maxRange)
	if err != nil {
		return "", err
	}
	defer body.Close()
	return toTmpFile(body)
}
