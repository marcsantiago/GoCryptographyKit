package caeser

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

const (
	// LOWERCASE ...
	LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	// UPPERCASE ...
	UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
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
			if strings.ContainsAny(string(r), UPPERCASE) {
				if pos > 'Z' {
					pos -= 26
				} else if pos < 'A' {
					pos += 26
				}
			} else if strings.ContainsAny(string(r), LOWERCASE) {
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
			if strings.ContainsAny(string(r), UPPERCASE) {
				if pos > 'Z' {
					pos -= 26
				} else if pos < 'A' {
					pos += 26
				}
			} else if strings.ContainsAny(string(r), LOWERCASE) {
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
