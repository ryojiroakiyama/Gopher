package mikujitou

import (
	"encoding/json"
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
	rand.Seed(time.Now().Unix())
	o := Omikuji{Fortune: fortune_message[rand.Intn(len(fortune_message))]}
	j, _ := json.Marshal(o)
	w.Write(j)
}
