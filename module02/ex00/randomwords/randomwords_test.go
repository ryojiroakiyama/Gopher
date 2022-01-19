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

//func nextLine(t *testing.T, sc *bufio.Scanner) {
//	switch {
//	case sc.Scan():
//		ch <- sc.Text()
//	case sc.Err() == nil:
//		os.Exit(0)
//	default:
//		panic(sc.Err())
//	}
//}

//func TestOut(t *testing.T) {
//	tests := []struct {
//		name string
//		want string
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Out(); got != tt.want {
//				t.Errorf("Out() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
