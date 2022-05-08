//Package randomewords has a list of English words
//and can move them out.
package randomwords

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"time"
)

const DefaultWords = "randomwords/words.txt"

var words []string

//InitWithFile make a word list.
func InitWithFile(filename string) error {
	file, err := os.Open(filename)
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

//Init make a default word list.
func Init() error {
	return InitWithFile(DefaultWords)
}

//List put all words in the list to io.Writer.
func List(out io.Writer) error {
	for _, word := range words {
		if _, err := io.WriteString(out, word+"\n"); err != nil {
			return err
		}
	}
	return nil
}

//Out return a word in the list randomly.
func Out() (word string) {
	len := len(words)
	if 0 < len {
		rand.Seed(time.Now().Unix())
		word = words[rand.Intn(len)]
	}
	return word
}
