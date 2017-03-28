package detect

import (
	"bufio"
	"bytes"
	"os"
	"path/filepath"
	"strings"
)

// EnglishWords ...
var EnglishWords map[string]error

const (
	// AlphaSpace ...
	AlphaSpace = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\t\n "
	// LetterPercent lowering the value might lead to more false positives
	LetterPercent = 85
)

func init() {
	path, _ := os.Getwd()
	words := filepath.Join(path, "/src/detect_english/dictionary.txt")
	file, err := os.Open(words)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	EnglishWords = make(map[string]error)
	for scanner.Scan() {
		EnglishWords[strings.TrimSpace(scanner.Text())] = nil
	}
}

func getEnglishCount(message string) float32 {
	possibleWords := strings.Split(removeNonLetters(strings.ToUpper(message)), " ")
	if len(possibleWords) == 0 {
		return 0.0
	}
	var matches int
	for _, word := range possibleWords {
		if _, ok := EnglishWords[word]; ok {
			matches++
		}
	}
	return float32(matches) / float32(len(possibleWords))
}

func removeNonLetters(message string) string {
	var buf bytes.Buffer
	for _, r := range message {
		if strings.ContainsAny(string(r), AlphaSpace) {
			buf.WriteRune(r)
		}
	}
	return buf.String()
}

// English ...
func English(message string, accuracy int16) bool {
	count := getEnglishCount(message) * 100
	wordMatch := false
	if count >= float32(accuracy) {
		wordMatch = true
	}
	numLetters := len(removeNonLetters(message))
	messageLetterPercent := (float32(numLetters) / float32(len(message))) * 100.0
	lettersMatch := false
	if messageLetterPercent >= LetterPercent {
		lettersMatch = true
	}
	return wordMatch && lettersMatch

}
