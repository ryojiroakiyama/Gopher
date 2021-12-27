//Package converter provides ability to convert
//from one file format to another format.
//This process is applied to all files
//in the directory passed as argument.
package converter

import (
	"reflect"
	"testing"
)

func Test_printError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			printError(tt.args.err)
		})
	}
}

func TestDo(t *testing.T) {
	type args struct {
		dir    string
		srcExt string
		dstExt string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := Do(tt.args.dir, tt.args.srcExt, tt.args.dstExt); (err != nil) != tt.wantErr {
				t.Errorf("Do() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getImages(t *testing.T) {
	type args struct {
		Extension string
	}
	tests := []struct {
		name    string
		args    args
		want    images
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := getImages(tt.args.Extension)
			if (err != nil) != tt.wantErr {
				t.Errorf("getImages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getImages() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyEachFile(t *testing.T) {
	type args struct {
		rootdir string
		c       conversion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if err := applyEachFile(tt.args.rootdir, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("applyEachFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
