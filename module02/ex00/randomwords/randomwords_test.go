package randomwords_test

import (
	"testing"
	"typing/randomwords"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "file exists or not",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := randomwords.Init(); (err != nil) != tt.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOut(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "normal",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomwords.Out(); got != tt.want {
				t.Errorf("Out() = %v, want = %v", got, tt.want)
			}
			//randomwords.Init()
			//if got := randomwords.Out(); got == "" {
			//	t.Errorf("after Init(), Out() = %v", got)
			//}
		})
	}
}
