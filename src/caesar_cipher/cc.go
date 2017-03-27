package caeser

import (
	"bytes"
	"log"
	"strings"
	"unicode"
)

const (
	// LOWERCASE ...
	LOWERCASE = "abcdefghijklmnopqrstuvwxyz"
	// UPPERCASE ...
	UPPERCASE = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

// PlainText ...
type PlainText struct {
	originalMessage string
	key             int16
	encodedMessage  string
	decodedMessage  string
}

// NewMessage ...
func NewMessage(message string) (p PlainText) {
	p.originalMessage = message
	return
}

// EncodeKey ...
func (p *PlainText) EncodeKey(key int16) {
	if key > 26 {
		log.Fatal("Key must be less then 27 and greater than 0")
	}
	if key == 0 {
		return
	}
	p.key = key

	var buf bytes.Buffer
	for _, r := range p.originalMessage {
		if unicode.IsLetter(r) {
			pos := r
			pos += rune(p.key)
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
	p.encodedMessage = buf.String()
}

// DecodeKey ...
func (p *PlainText) DecodeKey(key int16) {
	if key > 26 {
		panic("Key must be less than 27")
	}
	if key == 0 {
		panic("Key must be greater than 0")
	}
	p.key = -key
	var buf bytes.Buffer
	for _, r := range p.encodedMessage {
		if unicode.IsLetter(r) {
			pos := r
			pos += rune(p.key)
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
	p.decodedMessage = buf.String()
}

// GetEncodedMessage ...
func (p *PlainText) GetEncodedMessage() string {
	return p.encodedMessage
}

// GetDecodedMessage ...
func (p *PlainText) GetDecodedMessage() string {
	return p.decodedMessage
}
