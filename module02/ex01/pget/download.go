package pget

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
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

func divDownload(url string, minRange int64, maxRange int64) (string, error) {
	body, err := rangeRequest(url, minRange, maxRange)
	if err != nil {
		return "", err
	}
	defer body.Close()
	return toTmpFile(body)
}

func download(ctx context.Context, url string) ([]string, error) {
	eg, ctx := errgroup.WithContext(ctx)
	sizeSum, err := DataLength(url)
	if err != nil {
		return nil, err
	}
	numDiv := NumDivideRange(sizeSum)
	sizeDiv := sizeSum / int64(numDiv)
	divfiles := make([]string, numDiv)
	for i := 0; i < numDiv; i++ {
		i := i
		err := err
		eg.Go(func() error {
			minRange, maxRange := DownloadRange(i, numDiv, sizeDiv, sizeSum)
			select {
			case <-ctx.Done():
			default:
				divfiles[i], err = divDownload(url, minRange, maxRange)
			}
			return err
		})
	}
	if err := eg.Wait(); err != nil {
		return nil, err
	}
	return divfiles, nil
}
