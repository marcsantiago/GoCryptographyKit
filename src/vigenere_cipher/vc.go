package vigenere

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
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

func convert(i interface{}) (string, error) {
	switch i.(type) {
	case string:
		return i.(string), nil
	case *os.File:
		var buf bytes.Buffer
		f := i.(*os.File)
		defer f.Close()
		io.Copy(&buf, f)
		return buf.String(), nil
	default:
		return "", fmt.Errorf("Message must be of type string or file")
	}
}

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
func Encode(msg interface{}, key string) (string, error) {
	var buf bytes.Buffer

	message, err := convert(msg)
	if err != nil {
		return "", err
	}

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
func Decode(msg interface{}, key string) (string, error) {
	var buf bytes.Buffer
	message, err := convert(msg)
	if err != nil {
		return "", err
	}

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
func BruteForceDecrypt(message string, accuracy int) (string, error) {
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
