package pget

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
)

const (
	ONEDLMAX = 1000
)

func Do(filepath string, url string) (err error) {
	download := func(ctx context.Context, filepath string, url string) ([]string, error) {
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
				client := &http.Client{}
				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					return err
				}
				req.Header.Add("Range", RangeValue(minRange, maxRange-1))
				resp, err := client.Do(req)
				if err != nil {
					return err
				}
				defer resp.Body.Close()
				tmpfile, err := os.CreateTemp("", "")
				if err != nil {
					return err
				}
				defer func() {
					if cerr := tmpfile.Close(); cerr != nil {
						err = fmt.Errorf("fail to close: %v", cerr)
					}
				}()
				divfiles[i] = tmpfile.Name()
				fmt.Println("create:", tmpfile.Name())
				_, err = io.Copy(tmpfile, resp.Body)
				if err != nil {
					return err
				}
				return nil
			})
			if err = eg.Wait(); err != nil {
				return nil, err
			}
		}
		return divfiles, nil
	}

	divfiles, err := download(context.Background(), filepath, url)
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
