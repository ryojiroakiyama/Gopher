package pget

import (
	"context"
	"fmt"
	"os"
	"strings"
)

const (
	DivDownLoadMax = 1000
)

func Do(url string) error {
	divfiles, err := download(context.Background(), url)
	if err != nil {
		return fmt.Errorf("Do: %v", err)
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
		return fmt.Errorf("Do: %v", err)
	}
	return nil
}
