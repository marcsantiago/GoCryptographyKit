package main

import (
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
	msg, err = vigenere.Decode(msg, key)
	if err != nil {
		panic(err)
	}
	if plainText != msg {
		t.Log("Plaintext Message not decoded correctly")
		t.FailNow()
	}

	msg, err = vigenere.Encode(plainText, "candy")
	if err != nil {
		panic(err)
	}
	_, err = vigenere.BruteForceDecrypt(msg, 30)
	if err != nil {
		t.Log("Try playing with the accuracy, also remember this only works if the encrypt key is a single word")
		t.FailNow()
	}

}
