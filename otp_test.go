package main

import (
	"bytes"
	"io"
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

	var buf bytes.Buffer
	f, err := os.Open("message_test.txt")
	if err != nil {
		t.Error(err)
	}
	defer f.Close()
	io.Copy(&buf, f)
	original := buf.String()

	if original != decoded {
		t.Errorf("Message not the same\noriginalMessage: %s\nDecodedMessage: %s\n", decoded, original)
	}

}
