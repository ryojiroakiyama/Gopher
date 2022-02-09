package pget

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func divDownload(url string, minRange int64, maxRange int64) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("fail to send request: %v", err)
	}
	req.Header.Add("Range", RangeValue(minRange, maxRange-1))
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("fail to get response: %v", err)
	}
	defer resp.Body.Close()
	tmpfile, err := os.CreateTemp("", "")
	if err != nil {
		return "", fmt.Errorf("failt to create: %v", err)
	}
	defer func() {
		if cerr := tmpfile.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()
	fmt.Println("create:", tmpfile.Name())
	_, err = io.Copy(tmpfile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("fail to copy: %v", err)
	}
	return tmpfile.Name(), nil
}
