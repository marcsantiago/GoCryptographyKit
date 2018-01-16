package caeser

import (
	"testing"

	"github.com/marcsantiago/GoCryptographyKit/internal"
)

func TestCaserCipher(t *testing.T) {
	// Test basic key cipher
	for i := 1; i < 27; i++ {
		encodedMsg, err := Encode(internal.TestMessage, i)
		if err != nil {
			t.Error(err)
		}
		decodedMsg, err := Decode(encodedMsg, i)
		if err != nil {
			t.Error(err)
		}
		if decodedMsg != internal.TestMessage {
			t.Errorf("Message not the same\nkey: %d\noriginalMessage: %s\nDecodedMessage: %s\n", i, internal.TestMessage, decodedMsg)
		}
	}
	// test brute force
	encodedMsg, err := Encode(internal.TestMessage, 5)
	if err != nil {
		t.Error(err)
	}
	_, err = BruteForceDecrypt(encodedMsg, 20)
	if err != nil {
		t.Error(err)
	}

	encodedMsg, err = Encode(internal.TheRepublic, 5)
	if err != nil {
		t.Error(err)
	}

	decoded, err := Decode(encodedMsg, 5)
	if err != nil {
		t.Error(err)
	}
	original := internal.TheRepublic

	if original != decoded {
		t.Errorf("Message not the same\nkey: %d\noriginalMessage: %s\nDecodedMessage: %s\n", 5, decoded, original)
	}
}
