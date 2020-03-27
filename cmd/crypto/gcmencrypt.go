package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func encrypt() {

	dirName := "encrypted"
	path := "./" + dirName
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0744)
		fmt.Println("encryped successfully created")
	} else {
		fmt.Println("encrypted already exsists")
	}

	text, err := ioutil.ReadFile(".files/IT_Flye.pdf")
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

	ciphertext = append(ciphertext, sum[:]...)

	// write encrypted data to file
	err = ioutil.WriteFile(dirName+"/IT_Flye.pdf", ciphertext, 0777)

	fmt.Printf("%d \n", aesgcm.NonceSize())
	fmt.Printf("len text: %d\n", len(text))
	fmt.Printf("%x len sum: %d\n", sum, len(sum))
	fmt.Printf("len cipher: %d\n", len(ciphertext))
}
