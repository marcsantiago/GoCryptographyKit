package atbash

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var cipher = map[string]string{"a": "z", "c": "x", "b": "y", "e": "v", "d": "w", "g": "t", "f": "u", "i": "r", "h": "s", "k": "p", "j": "q", "m": "n", "l": "o", "o": "l", "n": "m", "q": "j", "p": "k", "s": "h", "r": "i", "u": "f", "t": "g", "w": "d", "v": "e", "y": "b", "x": "c", "z": "a"}

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

// Encode ...
func Encode(msg interface{}) (string, error) {
	var buf bytes.Buffer
	message, err := convert(msg)
	if err != nil {
		return "", err
	}
	reg, err := regexp.Compile(`[.,\/#!$%\^&\*;:{}=\-_~()\s]`)
	if err != nil {
		return "", err
	}
	message = strings.ToLower(reg.ReplaceAllString(message, ""))
	for i, r := range message {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			if v, ok := cipher[string(r)]; ok {
				buf.WriteString(v)
			} else {
				buf.WriteRune(r)
			}
			if i > 0 && (i+1)%5 == 0 {
				buf.WriteRune(' ')
			}

		}
	}
	return strings.TrimSpace(buf.String()), nil
}

// Decode ...
func Decode(msg interface{}) (string, error) {
	return Encode(msg)
}
