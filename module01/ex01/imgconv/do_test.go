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
	gFile := testdir + "OriginalGif.gif"
	tFile := testdir + "test.txt"
	pngImage := "image/png"
	jpgImage := "image/jpeg"
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
		{
			name: "gtoj",
			fields: fields{
				SrcExtension: "gif",
				DstExtension: "jpg",
				Decoder:      conversion.ActionGif{},
				Encoder:      conversion.ActionJpg{},
			},
			args: args{
				srcFileName: gFile,
			},
			wantErr:   false,
			wantFile:  dstFileName(gFile, "gif", "jpg"),
			wantImage: jpgImage,
		},
		{
			name: "form txt",
			fields: fields{
				SrcExtension: "png",
				DstExtension: "gif",
				Decoder:      conversion.ActionPng{},
				Encoder:      conversion.ActionGif{},
			},
			args: args{
				srcFileName: tFile,
			},
			wantErr:   true,
			wantFile:  "",
			wantImage: "",
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
			var err error
			if err = c.Do(tt.args.srcFileName); (err != nil) != tt.wantErr {
				t.Errorf("Converter.Do() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil {
				if _, err := os.Stat(tt.wantFile); err != nil {
					t.Errorf("out file dosen't exist, wantFile = %v", tt.wantFile)
				} else if contentType := http.DetectContentType(dstFileBytes(t, tt.wantFile)); contentType != tt.wantImage {
					t.Errorf("Converter.Do()'s out file image = %v, wantImage %v", contentType, tt.wantImage)
				}
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
