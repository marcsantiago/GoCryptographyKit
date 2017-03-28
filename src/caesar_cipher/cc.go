package caeser

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"

	"../detect_english"
)

const (
	// LowerCase ...
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	// UpperCase ...
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// Encode ...
func Encode(message string, key int16) (string, error) {
	if key > 26 {
		return "", fmt.Errorf("Key must be less than 27")
	}
	if key == 0 {
		return "", fmt.Errorf("Key must be greater than 0")
	}

	var buf bytes.Buffer
	for _, r := range message {
		if unicode.IsLetter(r) {
			pos := r
			pos += rune(key)
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
func Decode(message string, key int16) (string, error) {
	if key > 26 {
		return "", fmt.Errorf("Key must be less than 27")
	}
	if key == 0 {
		return "", fmt.Errorf("Key must be greater than 0")
	}
	key = -key
	var buf bytes.Buffer
	for _, r := range message {
		if unicode.IsLetter(r) {
			pos := r
			pos += rune(key)
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

// BruteForceDecrypt ...
func BruteForceDecrypt(encodedMessage string, accuracy int16) (string, error) {
	var buf bytes.Buffer
	for i := 1; i < 27; i++ {
		key := -i
		for _, r := range encodedMessage {
			if unicode.IsLetter(r) {
				pos := r
				pos += rune(key)
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
		message := buf.String()
		buf.Reset()
		if detect.English(message, accuracy) {
			return fmt.Sprintf("Key: %d Message: %s\n", key*-1, message), nil
		}
	}
	return "", fmt.Errorf("Message could be decoded, try lowering the accuracy level")
}
