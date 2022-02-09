package pget

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/sync/errgroup"
)

const (
	ONEDLMAX = 1000
)

func Do(url string) error {
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

	divfiles, err := download(context.Background(), url)
	if err != nil {
		return err
	}
	defer func() {
		for _, d := range divfiles {
			if d != "" {
				fmt.Println("remove:", d)
				os.Remove(d)
			}
		}
	}()
	if err := bindFiles(divfiles, url[strings.LastIndex(url, "/")+1:]); err != nil {
		return err
	}
	return nil
}

func bindFiles(srcNames []string, dstName string) (err error) {
	dstfile, err := os.Create(dstName)
	if err != nil {
		return err
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
			return err
		}
		defer srcfile.Close()
		_, err = io.Copy(dstfile, srcfile)
		if err != nil {
			return err
		}
	}
	return err
}
