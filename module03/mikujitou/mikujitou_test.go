package mikujitou_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"omikuji/mikujitou"
	"testing"
	"time"
)

type ShogatuTimeGetter struct {
	t time.Time
}

func (s *ShogatuTimeGetter) Now() time.Time {
	return s.t
}
func TestDrawOmikuji(t *testing.T) {
	type args struct {
		t mikujitou.TimeGetter
	}
	tests := []struct {
		name           string
		args           args
		accessUrl      string
		wantStatus     int
		DoShougatuTest bool
		wantOmikuji    string
	}{
		{
			name: "normal",
			args: args{
				t: nil,
			},
			accessUrl:      "/omikuji",
			wantStatus:     http.StatusOK,
			DoShougatuTest: false,
			wantOmikuji:    "",
		},
		{
			name: "shougatu",
			args: args{
				t: &ShogatuTimeGetter{
					time.Date(0000, time.January, 2, 00, 00, 00, 0, time.UTC),
				},
			},
			accessUrl:      "/omikuji",
			wantStatus:     http.StatusOK,
			DoShougatuTest: true,
			wantOmikuji:    "\"Dai-kichi\"",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			handlerGot := mikujitou.DrawOmikuji(tt.args.t)
			w := httptest.NewRecorder()
			r := httptest.NewRequest("GET", tt.accessUrl, nil)
			handlerGot(w, r)
			res := w.Result()
			defer res.Body.Close()
			if res.StatusCode != tt.wantStatus {
				t.Fatal("unexpected status code")
			}
			b, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal("unexpected error")
			}
			if tt.DoShougatuTest {
				if s := string(b); s != tt.wantOmikuji {
					t.Fatalf("unexpected response: %s: want->%s", s, tt.wantOmikuji)
				}
			}
		})
	}
}
