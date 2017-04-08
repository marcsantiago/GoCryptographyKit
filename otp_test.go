package main

import (
	"fmt"
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
	m, _ := otp.Decrypt(msg, key)
	fmt.Println(m)
}
