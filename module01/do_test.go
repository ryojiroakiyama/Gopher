package conversion_test

import (
	"testing"

	"github.com/ryojiroakiyama/convert/conversion"
)

func TestConverter_Do(t *testing.T) {
	type fields struct {
		srcExtension string
		dstExtension string
		encoder      conversion.Action
		decoder      conversion.Action
	}
	type args struct {
		srcFileName string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantFile string
		wantErr  bool
	}{
		{
			name: "jtop",
			fields: fields{
				srcExtension: "jpg",
				dstExtension: "png",
				encoder: conversion.ActionPng{},
				decoder: conversion.ActionJpg{},
			},
			args: args{ srcFileName: "jpgOrignal.jpg" },
			wantFile: "jpgOrignal.png",
			wantErr: false,
		}
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			c := Converter{
				srcExtension: tt.fields.srcExtension,
				dstExtension: tt.fields.dstExtension,
				encoder:      tt.fields.encoder,
				decoder:      tt.fields.decoder,
			}
			//if err := c.Do(tt.args.srcFileName); (err != nil) != tt.wantErr {
			//	t.Errorf("Converter.Do() error = %v, wantErr %v", err, tt.wantErr)
			//}
		})
	}
}
