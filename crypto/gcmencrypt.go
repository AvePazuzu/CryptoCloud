package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
)

func main() {

	text := []byte("My Super Secret Code Stuff")

	// to comape the decrypted data with the source data the checksum is passed to the Seal func
	sum := sha256.Sum256(text)

	passphrase := []byte("passphrasewhichneedstobe32bytes!")
	// to enshure the constant lenght of the key the passphrase is encoded by sha256
	key := sha256.Sum256(passphrase)

	// generate a new aes cipher using our 32 byte long key
	block, err := aes.NewCipher(key[:])

	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nonce, nonce, text, sum[:])

	// write encrypted data to file
	err = ioutil.WriteFile("encrypted.data", nonce, 0777)

	fmt.Printf("%d \n", aesgcm.NonceSize())
	fmt.Printf("%s\n", text)
	fmt.Printf("%x\n", sum)
	fmt.Printf("%s encoded: %x\n", text, ciphertext)

	// fmt.Printf("%x \n", key)
	// fmt.Printf("%x \n", plain)
}
