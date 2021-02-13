package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func encrypt() {

	// retrieve data to encrypt
	inDir := "./data_to_encrypt/"

	inData, err := ioutil.ReadDir(inDir)
	if err != nil {
		fmt.Println("Reading file information failed! Program is terminated! Error: ", err)
		return
	}

	fmt.Println("\nRetrieving data to encrypt...")
	time.Sleep(1 * time.Second)
	// retrieve names of data to encrypt
	for _, inNames := range inData {
		n := inNames.Name()
		//encs := append(encs, n)
		fmt.Printf("\n %q", n)
	}

	// text, err := ioutil.ReadFile("./data_to_encrypt/" + fileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(text)
	// to comape the decrypted data with the source data the checksum is passed to the Seal func
	// sum := sha256.Sum256(text)

	// Ask for passphrase by reading input from console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println()
	fmt.Print("\nTo continue Enter passphrase: ")

	input, err := reader.ReadString('\n')
	fmt.Println("... and better remember it.\n")
	if err != nil {
		fmt.Println(err)
		return
	}
	// To make input comparable delimiter needs to be removed
	input = strings.Replace(input, "\n", "", -1)
	// fmt.Printf("%q", input)

	// convert user input into byte slice
	passphrase := []byte(input)
	// fmt.Println(passphrase)

	// to enshure the constant lenght of the key the passphrase is encoded by sha256
	key := sha256.Sum256(passphrase)
	fmt.Println("Key: ", key)

	fmt.Println("\nEncryption started...")

	for _, inNames := range inData {

		// read data into text
		text, err := ioutil.ReadFile("./data_to_encrypt/" + inNames.Name())
		if err != nil {
			log.Fatal(err)
		}

		// to comape the decrypted data with the source data the checksum is passed to the Seal func
		sum := sha256.Sum256(text)

		// generate a new aes cipher using the 32 byte long key
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
		// compute encoding
		ciphertext := aesgcm.Seal(nonce, nonce, text, sum[:])

		// append the checksum of the data to the encoded data
		ciphertext = append(ciphertext, sum[:]...)

		// write encrypted data to file
		err = ioutil.WriteFile("./encrypted_data/"+inNames.Name(), ciphertext, 0777)

		fmt.Printf("\n%q successfully encrypted.", inNames.Name())

	}
	fmt.Println()
	fmt.Println("\nEncryption successfull... Your data is ready to be uploaded...")

}
