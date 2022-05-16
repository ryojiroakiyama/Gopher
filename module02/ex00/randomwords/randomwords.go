//Package randomewords has a list of English words
//and can move them out.
package randomwords

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"time"
)

const DefaultWords = "randomwords/words.txt"

var wordlist []string

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
		wordlist = append(wordlist, line)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	if len(wordlist) <= 0 {
		return fmt.Errorf("words is empty")
	}
	return nil
}

//Init make a default word list.
func Init() error {
	return InitWithFile(DefaultWords)
}

//List put all words in the list to io.Writer.
func List(out io.Writer) error {
	for _, word := range wordlist {
		if _, err := io.WriteString(out, word+"\n"); err != nil {
			return err
		}
	}
	return nil
}

//Out return a word in the list randomly.
func Out() string {
	var outword string
	len := len(wordlist)
	if 0 < len {
		rand.Seed(time.Now().UnixNano())
		outindex := rand.Intn(len)
		outword = wordlist[outindex]
		wordlist = append(wordlist[:outindex], wordlist[outindex+1:]...) // remove outword from wordlist
	}
	return outword
}
