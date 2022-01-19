package randomwords

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

const (
	FILENAME = "randomwords/words.txt"
)

var words []string

func Init() error {
	file, err := os.Open(FILENAME)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		words = append(words, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func Out() string {
	rand.Seed(time.Now().Unix())
	word := words[rand.Intn(len(words))]
	return word
}
