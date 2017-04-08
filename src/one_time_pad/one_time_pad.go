package otp

import (
	"bytes"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	// LowerCase ...
	LowerCase = "abcdefghijklmnopqrstuvwxyz"
	// UpperCase ...
	UpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Alpha ...
	Alpha = LowerCase + UpperCase
)

func binToString(s []byte) string {
	var counter int
	chunk := make([]byte, 7)
	var buf bytes.Buffer
	for i := range s {
		if i%8 == 0 {
			counter = 0
			if i, err := strconv.ParseInt(string(chunk), 2, 64); err == nil {
				buf.WriteString(string(i))
			}
		} else {
			chunk[counter] = s[i]
			counter++
		}
	}
	return buf.String()
}

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

func stringToBin(s string) (b string) {
	for _, c := range s {
		b = fmt.Sprintf("%s%.8b", b, c)
	}
	return
}

func generateKey(length int, save bool) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = Alpha[rand.Intn(len(Alpha))]
	}
	if save {
		t := time.Now()
		source := rand.NewSource(t.Unix())
		r := rand.New(source)
		f, err := os.Create(fmt.Sprintf("key_%d.txt", r.Int63()))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.Write(b)
	}
	return stringToBin(string(b))
}

// Encrypt ...
func Encrypt(msg interface{}, save bool) (string, error) {
	message, err := convert(msg)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	binMessage := stringToBin(message)
	key := generateKey(len(message), save)
	for i := 0; i < len(key); i++ {
		bin := int(binMessage[i]) ^ int(key[i])
		_, err := buf.WriteString(strconv.Itoa(bin))
		if err != nil {
			panic(err)
		}
	}

	if save {
		f, err := os.Create("encryted_message.txt")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString(buf.String())
	}
	return buf.String(), nil
}

// Decrypt ...
func Decrypt(msg interface{}, k interface{}) (string, error) {
	binMessage, err := convert(msg)
	if err != nil {
		return "", err
	}
	keyS, err := convert(k)
	if err != nil {
		return "", err
	}
	key := stringToBin(keyS)
	var buf bytes.Buffer
	for i := 0; i < len(key); i++ {
		bin := int(binMessage[i]) ^ int(key[i])
		_, err := buf.WriteString(strconv.Itoa(bin))
		if err != nil {
			panic(err)
		}
	}
	str := binToString(buf.Bytes())
	return str, nil
}
