package main

import (
	"os"
	"testing"

	"./src/caesar_cipher"
)

const message = `On offering to help the blind man, the man who then stole his car, had not, at that precise moment,
  had any evil intention, quite the contrary, what he did was nothing more than obey those feelings of generosity and altruism which,
  as everyone knows, are the two best traits of human nature and to be found in much more hardened criminals than this one,
  a simple car-thief without any hope of advancing in his profession, exploited by the real owners of this enterprise,
  for it is they who take advantage of the needs of the poor.`

func TestCaserCipher(t *testing.T) {
	// Test basic key cipher
	for i := 1; i < 27; i++ {
		encodedMsg, err := caeser.Encode(message, int16(i))
		if err != nil {
			panic(err)
		}
		decodedMsg, err := caeser.Decode(encodedMsg, int16(i))
		if err != nil {
			panic(err)
		}
		if decodedMsg != message {
			t.Log("Message not the same, key:", i)
			t.Log(decodedMsg)
			t.FailNow()
		}
	}
	// test brute force
	encodedMsg, err := caeser.Encode(message, 5)
	if err != nil {
		panic(err)
	}
	_, err = caeser.BruteForceDecrypt(encodedMsg, 20)
	if err != nil {
		panic(err)
	}

	// test file encrytion/decryption
	f, err := os.Open("the_republic.txt")
	if err != nil {
		panic(err)
	}
	encodedMsg, err = caeser.Encode(f, 5)
	if err != nil {
		panic(err)
	}
	_, err = caeser.Decode(encodedMsg, 5)
	if err != nil {
		panic(err)
	}
}
