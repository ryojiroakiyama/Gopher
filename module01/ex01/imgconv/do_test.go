//Package conversion provides ability to convert file format
//according to the specified format.
package conversion_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	conversion "github.com/ryojiroakiyama/convert/imgconv"
)

func TestConverter_Do(t *testing.T) {
	testdir := "../testdata/"
	jFile := testdir + "OriginalJpg.jpg"
	pngImage := "image/png"
	type fields struct {
		SrcExtension string
		DstExtension string
		Decoder      conversion.Action
		Encoder      conversion.Action
	}
	type args struct {
		srcFileName string
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantErr   bool
		wantFile  string
		wantImage string
	}{
		{
			name: "jtop",
			fields: fields{
				SrcExtension: "jpg",
				DstExtension: "png",
				Decoder:      conversion.ActionJpg{},
				Encoder:      conversion.ActionPng{},
			},
			args: args{
				srcFileName: jFile,
			},
			wantErr:   false,
			wantFile:  dstFileName(jFile, "jpg", "png"),
			wantImage: pngImage,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := conversion.Converter{
				SrcExtension: tt.fields.SrcExtension,
				DstExtension: tt.fields.DstExtension,
				Decoder:      tt.fields.Decoder,
				Encoder:      tt.fields.Encoder,
			}
			if err := c.Do(tt.args.srcFileName); (err != nil) != tt.wantErr {
				t.Errorf("Converter.Do() error = %v, wantErr %v", err, tt.wantErr)
			} else if _, err := os.Stat(tt.wantFile); err != nil {
				t.Errorf("out file dosen't exist, wantFile = %v", tt.wantFile)
			} else if contentType := http.DetectContentType(dstFileBytes(t, tt.wantFile)); contentType != tt.wantImage {
				t.Errorf("Converter.Do()'s out file image = %v, wantImage %v", contentType, tt.wantImage)
			}
		})
	}
}

func dstFileName(srcFileName string, srcExtension string, dstExtension string) string {
	return strings.TrimSuffix(srcFileName, "."+srcExtension) + "." + dstExtension
}

func dstFileBytes(t *testing.T, fileName string) []byte {
	t.Helper()
	file, err := os.Open(fileName)
	if err != nil {
		t.Fatalf("err : %s", err)
	}
	defer file.Close()
	defer os.Remove(fileName)
	srcBytes, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("err : %s", err)
	}
	return srcBytes
}
