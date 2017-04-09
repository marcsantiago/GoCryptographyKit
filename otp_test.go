package main

import (
	"os"
	"testing"

	"./src/one_time_pad"
)

func TestOneTimePad(t *testing.T) {
	msg, err := os.Open("message_test.txt")
	if err != nil {
		panic(err)
	}
	defer msg.Close()
	key, err := os.Open("key_test.txt")
	if err != nil {
		panic(err)
	}
	defer key.Close()
	_, err = otp.Decrypt(msg, key)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

}
