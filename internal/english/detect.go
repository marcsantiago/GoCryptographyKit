package english

import (
	"bytes"
	"os"
	"strconv"
	"strings"
)

var (
	// LetterPercent represents the threshold by which one considers a word english.
	// lowering the value might lead to more false positives, the default is 85
	// this can be set as a enviromental variable -> PERCENT
	LetterPercent float32
)

func init() {
	LetterPercent = 85.0
	percent := os.Getenv("PERCENT")
	if p, err := strconv.ParseFloat(percent, 32); err == nil && len(percent) > 0 {
		LetterPercent = float32(p)
	}
}

func getEnglishCount(message string) float32 {
	possibleWords := strings.Split(removeNonLetters(strings.ToUpper(message)), " ")
	if len(possibleWords) == 0 {
		return 0.0
	}
	var matches int
	for _, word := range possibleWords {
		if _, ok := englishWords[word]; ok {
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

// IsEnglish tries to detemine if a word of sentence is English
func IsEnglish(message string, accuracy int) bool {
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

// GetWordList retrieves the list of words used to build the English dictonary
func GetWordList() []string {
	l := make([]string, 0, len(englishWords))
	for k := range englishWords {
		l = append(l, k)
	}
	return l
}
