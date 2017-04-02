package vigenere

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"../detect_english"
)

const (
	// LowerCase ...
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	// UpperCase ...
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Alpha ...
	Alpha = LowerCase + UpperCase
)

func buildKey(message, key string) bytes.Buffer {
	var buf bytes.Buffer
	var index int
	for i := range message {
		index = i % len(key)
		if unicode.IsLetter(rune(key[index])) {
			buf.WriteByte(key[index] % 27)
		}
	}
	// add the rest of the buffer
	for i := buf.Len(); i < len(message); i++ {
		index = i % len(key)
		buf.WriteByte(key[index] % 27)
	}
	return buf
}

// Encode ...
func Encode(message, key string) (string, error) {
	var buf bytes.Buffer
	k := buildKey(message, key)
	generatedKey := k.String()
	for i, r := range message {
		if unicode.IsLetter(r) {
			pos := r
			pos += rune(generatedKey[i])

			if strings.ContainsAny(string(r), UpperCase) {
				if pos > 'Z' {
					pos -= 26
				} else if pos < 'A' {
					pos += 26
				}
			} else if strings.ContainsAny(string(r), LowerCase) {
				if pos > 'z' {
					pos -= 26
				} else if pos < 'a' {
					pos += 26
				}
			}
			buf.WriteRune(pos)
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String(), nil
}

// Decode ...
func Decode(message, key string) (string, error) {
	var buf bytes.Buffer
	k := buildKey(message, key)
	generatedKey := k.String()

	for i, r := range message {
		if unicode.IsLetter(r) {
			pos := r
			pos += -rune(generatedKey[i])

			if strings.ContainsAny(string(r), UpperCase) {
				if pos > 'Z' {
					pos -= 26
				} else if pos < 'A' {
					pos += 26
				}
			} else if strings.ContainsAny(string(r), LowerCase) {
				if pos > 'z' {
					pos -= 26
				} else if pos < 'a' {
					pos += 26
				}
			}
			buf.WriteRune(pos)
		} else {
			buf.WriteRune(r)
		}
	}
	return buf.String(), nil
}

// BruteForceDecrypt may work if the key is a single word.  Keys longer then no word have very little change of working with this methods
func BruteForceDecrypt(message string, accuracy int16) (string, error) {
	path, _ := os.Getwd()
	words := filepath.Join(path, "/src/detect_english/dictionary.txt")
	file, err := os.Open(words)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	englishWords := []string{}
	for scanner.Scan() {
		englishWords = append(englishWords, scanner.Text())
		englishWords = append(englishWords, strings.ToLower(scanner.Text()))
	}

	var buf bytes.Buffer
	for _, possibleKey := range englishWords {
		k := buildKey(message, possibleKey)
		generatedKey := k.String()
		for i, r := range message {
			if unicode.IsLetter(r) {
				pos := r
				pos += -rune(generatedKey[i])

				if strings.ContainsAny(string(r), UpperCase) {
					if pos > 'Z' {
						pos -= 26
					} else if pos < 'A' {
						pos += 26
					}
				} else if strings.ContainsAny(string(r), LowerCase) {
					if pos > 'z' {
						pos -= 26
					} else if pos < 'a' {
						pos += 26
					}
				}
				buf.WriteRune(pos)
			} else {
				buf.WriteRune(r)
			}
		}
		msg := buf.String()
		buf.Reset()
		if detect.English(msg, accuracy) {
			return fmt.Sprintf("Key: %s Message: %s\n", possibleKey, msg), nil
		}
	}
	return "", fmt.Errorf("Message could be decoded, try lowering the accuracy level")
}
