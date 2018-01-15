package otp

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/marcsantiago/GoCryptographyKit/internal/convert"
	"github.com/marcsantiago/GoCryptographyKit/internal/english"
)

func binToString(s []byte) string {
	output := make([]byte, len(s)/8)
	for i := 0; i < len(output); i++ {
		val, err := strconv.ParseInt(string(s[i*8:(i+1)*8]), 2, 64)
		if err == nil {
			output[i] = byte(val)
		}
	}
	return string(output)
}

func stringToBin(s string) string {
	var b string
	for _, c := range s {
		b = fmt.Sprintf("%s%.8b", b, c)
	}
	return b
}

func generateKey(length int, save bool) string {
	b := make([]byte, length)
	alphaL := len(english.Alpha)
	for i := range b {
		b[i] = english.Alpha[rand.Intn(alphaL)]
	}

	if save {
		t := time.Now()
		source := rand.NewSource(t.Unix())
		r := rand.New(source)
		f, err := os.Create(fmt.Sprintf("key_%d.txt", r.Int63()))
		if err != nil {
			return ""
		}
		defer f.Close()
		f.Write(b)
	}
	return string(b)
}

// Encrypt ...
func Encrypt(msg interface{}, save bool) (string, string, error) {
	message, err := convert.RetrieveDataFromStringOrFile(msg)
	if err != nil {
		return "", "", err
	}

	var buf bytes.Buffer
	binMessage := stringToBin(message)
	keyString := generateKey(len(message), save)
	key := stringToBin(keyString)
	for i := 0; i < len(key); i++ {
		bin := int(binMessage[i]) ^ int(key[i])
		_, err := buf.WriteString(strconv.Itoa(bin))
		if err != nil {
			return "", "", err
		}
	}

	if save {
		f, err := os.Create("encryted_message.txt")
		if err != nil {
			return "", "", err
		}
		defer f.Close()
		f.WriteString(buf.String())
	}
	return buf.String(), keyString, nil
}

// Decrypt ...
func Decrypt(msg interface{}, k interface{}) (string, error) {
	binMessage, err := convert.RetrieveDataFromStringOrFile(msg)
	if err != nil {
		return "", err
	}

	keyS, err := convert.RetrieveDataFromStringOrFile(k)
	if err != nil {
		return "", err
	}

	key := stringToBin(keyS)
	var buf bytes.Buffer
	for i := 0; i < len(key); i++ {
		bin := int(binMessage[i]) ^ int(key[i])
		_, err := buf.WriteString(strconv.Itoa(bin))
		if err != nil {
			return "", err
		}
	}

	str := binToString(buf.Bytes())
	return str, nil
}
