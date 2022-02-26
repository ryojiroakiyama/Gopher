package mikujitou

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var defaultFortuneContents = [...]string{
	"Dai-kichi",
	"Kichi",
	"Chuu-kichi",
	"Sho-kichi",
	"Sue-kichi",
	"Kyo",
	"Dai-kyo",
}

var defaultStudyFortuneContents = [...]string{
	"Good",
	"Not so good",
}

type DefaultTimeGetter struct{}

func (d *DefaultTimeGetter) Now() time.Time {
	return time.Now()
}

type TimeGetter interface {
	Now() time.Time
}

func DrawOmikuji(t TimeGetter) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		if t == nil {
			t = &DefaultTimeGetter{}
		}
		timeOfDay := t.Now()
		layout := "01-02"
		start, err := time.Parse(layout, "01-01")
		if err != nil {
			log.Fatal(err)
		}
		now, err := time.Parse(layout, timeOfDay.Format(layout))
		if err != nil {
			log.Fatal(err)
		}
		end, err := time.Parse(layout, "01-03")
		if err != nil {
			log.Fatal(err)
		}
		w.Write([]byte(fmt.Sprintf("s:%v, n:%v, e:%v\n", start.String(), now.String(), end.String())))
		var omikuji string
		// !(start > now) && !(now > end) -> (start <= now) && (now <= end)
		if !start.After(now) && !now.After(end) {
			omikuji = defaultFortuneContents[0]
		} else {
			rand.Seed(timeOfDay.Unix())
			omikuji = defaultFortuneContents[rand.Intn(len(defaultFortuneContents))]
		}
		j, _ := json.Marshal(omikuji)
		w.Write(j)
	}
}
