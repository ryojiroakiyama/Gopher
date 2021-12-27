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
					srcExtension: JPG,
					dstExtension: PNG,
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

