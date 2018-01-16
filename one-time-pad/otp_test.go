package otp

import (
	"log"
	"testing"

	"github.com/marcsantiago/GoCryptographyKit/internal"
)

func TestOneTimePad(t *testing.T) {
	decoded, err := Decrypt(internal.TestBinMessage, internal.TestKey)
	if err != nil {
		t.Error(err)
	}

	if internal.TestMessage != decoded {
		t.Errorf("Message not the same\noriginalMessage: %s\nDecodedMessage: %s\n", decoded, internal.TestMessage)
	}
}

func BenchmarkOneTimePad(b *testing.B) {
	b.StopTimer()
	for i := 0; i < 100; i++ {
		b.StartTimer()
		encryptedMsg, key, err := Encrypt(internal.TestMessage, false)
		if err != nil {
			log.Fatalln(err)
		}
		_, err = Decrypt(encryptedMsg, key)
		if err != nil {
			log.Fatalln(err)
		}
		b.StopTimer()
	}
	b.StopTimer()

}
