package main

import (
	"fmt"
	"testing"

	"./src/vigenere_cipher"
)

const plainText = `"Be normal, and the crowd will accept you. Be deranged, and they will make you their leader" -Wheel Of Time`

const key = `Hey, diddle, diddle, The cat and the fiddle, The cow jumped over the moon; The little dog laughed To see such sport, And the dish ran away with the spoon.`

func TestVigenereCipher(t *testing.T) {
	msg, err := vigenere.Encode(plainText, key)
	if err != nil {
		panic(err)
	}
	fmt.Println(msg)
	msg, err = vigenere.Decode(msg, key)
	if err != nil {
		panic(err)
	}
	if plainText != msg {
		t.Log("Plaintext Message not decoded correctly")
		t.FailNow()
	}
}
