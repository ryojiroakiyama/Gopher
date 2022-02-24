package mikujitou

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var fortuneContents = []string{
	"Dai-kichi",
	"Kichi",
	"Chuu-kichi",
	"Sho-kichi",
	"Sue-kichi",
	"Kyo",
	"Dai-kyo",
}

type Omikuji struct {
	Fortune string `json:"Fortune: "`
}

//time型はafterなど使って比較できる
//parse()で日付から時間を作って, 比較
//現在日付はformatで一旦取得
func Get(w http.ResponseWriter, _ *http.Request) {
	t := time.Now()
	layout := "01-02"
	start, err := time.Parse(layout, "01-01")
	if err != nil {
		log.Fatal(err)
	}
	now, err := time.Parse(layout, t.Format(layout))
	if err != nil {
		log.Fatal(err)
	}
	end, err := time.Parse(layout, "01-03")
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(fmt.Sprintf("s:%v, n:%v, e:%v\n", start.String(), now.String(), end.String())))
	var fortuneIndex int
	// !(start > now) && !(now > end) -> (start <= now) && (now <= end)
	if !start.After(now) && !now.After(end) {
		fortuneIndex = 0
	} else {
		rand.Seed(t.Unix())
		fortuneIndex = rand.Intn(len(fortuneContents))
	}
	o := Omikuji{Fortune: fortuneContents[fortuneIndex]}
	j, _ := json.Marshal(o)
	w.Write(j)
}
