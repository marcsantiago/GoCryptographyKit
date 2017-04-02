package vigenere

import (
	"bytes"
	"strings"
	"unicode"
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
