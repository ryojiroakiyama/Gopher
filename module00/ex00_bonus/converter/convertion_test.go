package converter

import (
	"testing"
)

func Test_getDstFileName(t *testing.T) {
	type args struct {
		srcFileName string
		c           conversion
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "dstfilename",
			args: args{
				srcFileName: "src",
				c: conversion{
					srcExtension: "jpg",
					dstExtension: "png",
				},
			},
			want: "src.png",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDstFileName(tt.args.srcFileName, tt.args.c); got != tt.want {
				t.Errorf("getDstFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convert(t *testing.T) {
	type args struct {
		srcFileName string
		c           conversion
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := convert(tt.args.srcFileName, tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("convert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
