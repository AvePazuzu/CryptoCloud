package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func main() {

	text := []byte("My Super Secret Code Stuff")

	key := []byte("passphrasewhichneedstobe32bytes!")

	// generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key)

	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	aesgcm.NonceSize()
	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, text, nil)

	fmt.Printf("%d \n", aesgcm.NonceSize())
	fmt.Printf("%s\n", text)
	fmt.Printf("%s encoded: %x\n", text, ciphertext)

	// fmt.Printf("%x \n", key)
	// fmt.Printf("%x \n", plain)
}
