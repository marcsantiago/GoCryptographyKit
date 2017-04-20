package main

import (
	"log"
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

func BenchmarkOneTimePad(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		b.StartTimer()
		encryptedMsg, key, err := otp.Encrypt(message, false)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = otp.Decrypt(encryptedMsg, key)
		if err != nil {
			log.Fatalln(err)
		}
		b.StopTimer()
	}
	b.StopTimer()

}
