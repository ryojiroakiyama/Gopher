package mikujitou

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var fortune_message = []string{
	"Dai-kichi",
	"Kichi",
	"Chuu-kichi",
	"Sho-kichi",
	"Sue-kichi",
	"Kyo",
	"Dai-kyo",
}

type Omikuji struct {
	Fortune string `json:"fortune-omikuji"`
}

func Get(w http.ResponseWriter, _ *http.Request) {
	t := time.Now()
	rand.Seed(t.Unix())
	layout := "01-02"
	start, _ := time.Parse(layout, "01-01")
	now, _ := time.Parse(layout, t.Format(layout))
	end, _ := time.Parse(layout, "01-03")
	w.Write([]byte(fmt.Sprintf("s:%v, n:%v, e:%v\n", start.String(), now.String(), end.String())))
	var fortune_index int
	// !(start > now) && !(now > end)
	if !start.After(now) && !now.After(end) {
		fortune_index = 0
	} else {
		fortune_index = rand.Intn(len(fortune_message))
	}
	o := Omikuji{Fortune: fortune_message[fortune_index]}
	j, _ := json.Marshal(o)
	w.Write(j)
}
