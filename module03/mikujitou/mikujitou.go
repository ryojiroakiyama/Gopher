package mikujitou

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type DefaultTimeGetter struct{}

func (d *DefaultTimeGetter) Now() time.Time {
	return time.Now()
}

type TimeGetter interface {
	Now() time.Time
}

func isWithinTime(layout string, start string, end string, timeNow time.Time) (bool, error) {
	startTime, err := time.Parse(layout, start)
	if err != nil {
		return false, err
	}
	timeNowAdjusted, err := time.Parse(layout, timeNow.Format(layout))
	if err != nil {
		return false, err
	}
	endTime, err := time.Parse(layout, end)
	if err != nil {
		return false, err
	}
	// !(startTime > timeNowAdjusted) && !(timeNowAdjusted > endTime) -> (startTime <= timeNowAdjusted) && (timeNowAdjusted <= endTime)
	if !startTime.After(timeNowAdjusted) && !timeNowAdjusted.After(endTime) {
		return true, nil
	} else {
		return false, nil
	}
}

func getRandomContent(contents []string) string {
	rand.Seed(time.Now().Unix())
	return contents[rand.Intn(len(contents))]
}

func getFortuneContent(t TimeGetter) (string, error) {
	var fortuneContents = [...]string{
		"Dai-kichi",
		"Kichi",
		"Chuu-kichi",
		"Sho-kichi",
		"Sue-kichi",
		"Kyo",
		"Dai-kyo",
	}
	layout := "01-02"
	startShougatu := "01-01"
	endShougatu := "01-03"
	if isShougatu, err := isWithinTime(layout, startShougatu, endShougatu, t.Now()); err != nil {
		return "", err
	} else if isShougatu {
		return fortuneContents[0], nil
	}
	return getRandomContent(fortuneContents[:]), nil
}

func getStudyContent() string {
	var studyContents = [...]string{
		"Good",
		"Not so good",
	}
	return getRandomContent(studyContents[:])
}

func DrawOmikuji(t TimeGetter) func(w http.ResponseWriter, _ *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		if t == nil {
			t = &DefaultTimeGetter{}
		}
		omikuji, err := getFortuneContent(t)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		j, err := json.Marshal(omikuji)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(j)
	}
}
