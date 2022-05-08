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
			name:    "file not exist",
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
	want := ""
	if got := randomwords.Out(); got != want {
		t.Errorf("Out() = %v, want = %v", got, want)
	}
	if err := randomwords.InitWithFile("words.txt"); err != nil {
		t.Fatal("fail to init", err)
	}
	if got := randomwords.Out(); got == "" {
		t.Errorf("Out() returnes empty string")
	}
}
