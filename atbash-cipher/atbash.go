package atbash

import (
	"bytes"
	"regexp"
	"strings"
	"unicode"

	"github.com/marcsantiago/GoCryptographyKit/internal/convert"
)

var (
	cipher = map[string]string{"a": "z", "c": "x", "b": "y", "e": "v", "d": "w", "g": "t", "f": "u", "i": "r", "h": "s", "k": "p", "j": "q", "m": "n", "l": "o", "o": "l", "n": "m", "q": "j", "p": "k", "s": "h", "r": "i", "u": "f", "t": "g", "w": "d", "v": "e", "y": "b", "x": "c", "z": "a"}
	reg    = regexp.MustCompile(`[.,\/#!$%\^&\*;:{}=\-_~()\s]`)
)

// Encode ...
func Encode(msg interface{}) (string, error) {
	var buf bytes.Buffer
	message, err := convert.RetrieveDataFromStringOrFile(msg)
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
