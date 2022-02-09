package pget

import (
	"context"
	"fmt"
	"io"
	"os"

	"golang.org/x/sync/errgroup"
)

const (
	ONEDLMAX = 1000
)

func Do(filepath string, url string) (err error) {
	download := func(ctx context.Context, url string) ([]string, error) {
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
			eg.Go(func() (err error) {
				minRange, maxRange := DownloadRange(i, numDiv, sizeDiv, sizeSum)
				divfiles[i], err = divDownload(url, minRange, maxRange)
				return err
			})
			if err = eg.Wait(); err != nil {
				return nil, err
			}
		}
		return divfiles, nil
	}

	divfiles, err := download(context.Background(), url)
	defer func() {
		for _, d := range divfiles {
			if d != "" {
				fmt.Println("remove:", d)
				os.Remove(d)
			}
		}
	}()
	dstfile, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := dstfile.Close(); cerr != nil {
			err = fmt.Errorf("fail to close: %v", cerr)
		}
	}()
	for _, srcfileName := range divfiles {
		srcfile, err := os.Open(srcfileName)
		if err != nil {
			os.Remove(dstfile.Name())
			return err
		}
		_, err = io.Copy(dstfile, srcfile)
		if err != nil {
			return err
		}
	}
	return err
}
