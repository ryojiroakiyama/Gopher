package conversion_test

import (
	"reflect"
	"testing"

	"github.com/ryojiroakiyama/convert/imgconv"
)

func TestNewConverter(t *testing.T) {
	type args struct {
		srcFormat string
		dstFormat string
	}
	tests := []struct {
		name    string
		args    args
		want    *conversion.Converter
		wantErr bool
	}{
		{
			name: "jtop",
			args: args{
				srcFormat: "jpg",
				dstFormat: "png",
			},
			want: &conversion.Converter{
				SrcExtension: "jpg",
				DstExtension: "png",
				Decoder:      conversion.ActionJpg{},
				Encoder:      conversion.ActionPng{},
			},
			wantErr: false,
		},
		{
			name: "gtoj",
			args: args{
				srcFormat: "gif",
				dstFormat: "jpg",
			},
			want: &conversion.Converter{
				SrcExtension: "gif",
				DstExtension: "jpg",
				Decoder:      conversion.ActionGif{},
				Encoder:      conversion.ActionJpg{},
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				srcFormat: "jg",
				dstFormat: "png",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := conversion.NewConverter(tt.args.srcFormat, tt.args.dstFormat)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConverter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConverter() = %v, want %v", got, tt.want)
			}
		})
	}
}
