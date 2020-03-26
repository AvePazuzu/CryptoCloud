package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, Gopher")

	ciphertext, err := ioutil.ReadFile("encrypted.data")
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("len file: %d\n", len(ciphertext))

	passphrase := []byte("passphrasewhichneedstobe32bytes!")
	// to enshure the constant lenght of the key the passphrase is encoded by sha256
	key := sha256.Sum256(passphrase)

	block, err := aes.NewCipher(key[:])
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	// ciphertext[12:len(ciphertext[12:])-32],ciphertext[12:len(ciphertext)-32], ciphertext[len(ciphertext)-32:]
	nonce, ciphertext, add := ciphertext[:12], ciphertext[12:len(ciphertext)-32], ciphertext[len(ciphertext)-32:]

	// add := make([]byte, 32)
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, add)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Printf("len nonce: %d\nlen ciphertext: %d\nlen add: %d\n", len(nonce), len(ciphertext), len(add))
	fmt.Printf("add is  : %x\n", add)
	fmt.Printf("plain is: %x\n", sha256.Sum256(plaintext))
	sum := sha256.Sum256(plaintext)
	fmt.Printf("Equal: %t\n", bytes.Equal(sum[:], add))
	fmt.Printf("%s\n", plaintext)

}
