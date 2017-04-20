package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"./src/vigenere_cipher"
)

const plainText = `"Be normal, and the crowd will accept you. Be deranged, and they will make you their leader" -Wheel Of Time`

const key = `Hey, diddle, diddle, The cat and the fiddle, The cow jumped over the moon; The little dog laughed To see such sport, And the dish ran away with the spoon.`

func TestVigenereCipher(t *testing.T) {
	// test basic encryption
	msg, err := vigenere.Encode(plainText, key)
	if err != nil {
		t.Error(err)
	}
	msg, err = vigenere.Decode(msg, key)
	if err != nil {
		t.Error(err)
	}
	if plainText != msg {
		t.Error("Plaintext Message not decoded correctly")
	}

	// test brute force
	msg, err = vigenere.Encode(plainText, "candy")
	if err != nil {
		t.Error(err)
	}
	_, err = vigenere.BruteForceDecrypt(msg, 30)
	if err != nil {
		t.Error("Try playing with the accuracy, also remember this only works if the encrypt key is a single word")
	}

	// test file encrytion/decryption
	f, err := os.Open("the_republic.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	msg, err = vigenere.Encode(f, key)
	if err != nil {
		t.Error(err)
	}

	var buf bytes.Buffer
	f, err = os.Open("the_republic.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	io.Copy(&buf, f)
	original := buf.String()

	decoded, err := vigenere.Decode(msg, key)
	if original != decoded {
		t.Errorf("Message not the same\nkey: %d\noriginalMessage: %s\nDecodedMessage: %s\n", 5, decoded, original)
	}

}
