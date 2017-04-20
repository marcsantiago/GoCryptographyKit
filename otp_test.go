package main

import (
	"os"
	"testing"

	"./src/one_time_pad"
)

func TestOneTimePad(t *testing.T) {
	msg, err := os.Open("message_test.txt")
	if err != nil {
		t.Error(err)
	}
	defer msg.Close()
	key, err := os.Open("key_test.txt")
	if err != nil {
		t.Error(err)
	}
	defer key.Close()

	decoded, err := otp.Decrypt(msg, key)
	if err != nil {
		t.Error(err)
	}

	if message != decoded {
		t.Errorf("Message not the same\noriginalMessage: %s\nDecodedMessage: %s\n", decoded, message)
	}

}
