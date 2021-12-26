//Package converter provides ability to convert
//from one file format to another format.
//This process is applied to all files
//in the directory passed as argument.
package converter

import (
	"bytes"
	"image"
	"io"
	"reflect"
	"testing"
)

func Test_imageJpg_decode(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		i       imageJpg
		args    args
		want    image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := imageJpg{}
			got, err := i.decode(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageJpg.decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageJpg.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imageJpg_encode(t *testing.T) {
	type args struct {
		m image.Image
	}
	tests := []struct {
		name    string
		i       imageJpg
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := imageJpg{}
			w := &bytes.Buffer{}
			if err := i.encode(w, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("imageJpg.encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("imageJpg.encode() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_imagePng_decode(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		i       imagePng
		args    args
		want    image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := imagePng{}
			got, err := i.decode(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("imagePng.decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imagePng.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imagePng_encode(t *testing.T) {
	type args struct {
		m image.Image
	}
	tests := []struct {
		name    string
		i       imagePng
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := imagePng{}
			w := &bytes.Buffer{}
			if err := i.encode(w, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("imagePng.encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("imagePng.encode() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}

func Test_imageGif_decode(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		i       imageGif
		args    args
		want    image.Image
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := imageGif{}
			got, err := i.decode(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("imageGif.decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("imageGif.decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_imageGif_encode(t *testing.T) {
	type args struct {
		m image.Image
	}
	tests := []struct {
		name    string
		i       imageGif
		args    args
		wantW   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			i := imageGif{}
			w := &bytes.Buffer{}
			if err := i.encode(w, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("imageGif.encode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("imageGif.encode() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
