package converter

import (
	"reflect"
	"testing"
)

func Test_getSrcBytes(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := getSrcBytes(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSrcBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSrcBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeDstFile(t *testing.T) {
	type args struct {
		fileName string
		contents []byte
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
			if err := makeDstFile(tt.args.fileName, tt.args.contents); (err != nil) != tt.wantErr {
				t.Errorf("makeDstFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
