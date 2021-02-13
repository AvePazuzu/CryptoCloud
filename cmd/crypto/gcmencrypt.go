package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func encrypt() {

	// retrieve the files to encrypt
	inDir := "./data_to_encrypt/"

	inData, err := ioutil.ReadDir(inDir)
	if err != nil {
		fmt.Println("Reading file information failed! Program is terminated! Error: ", err)
		return
	}

	// retrieve names of data to encrypt
	for _, inNames := range inData {
		n := inNames.Name()
		//encs := append(encs, n)
		fmt.Println(n)
	}

	// text, err := ioutil.ReadFile("./data_to_encrypt/" + fileName)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(text)
	// to comape the decrypted data with the source data the checksum is passed to the Seal func
	// sum := sha256.Sum256(text)

	// Ask for passphrase
	for {
		// Read input from console
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nEnter passphrase: ")

		input, err := reader.ReadString('\n')

		// To make input comparable delimiter needs to be removed
		input = strings.Replace(input, "\n", "", -1)
		passphrase := []byte(input)

		// to enshure the constant lenght of the key the passphrase is encoded by sha256
		key := sha256.Sum256(passphrase)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(passphrase)
		fmt.Println(input)
		fmt.Println(key)
		return
	}

	// // generate a new aes cipher using our 32 byte long key
	// block, err := aes.NewCipher(key[:])

	// if err != nil {
	// 	panic(err.Error())
	// }

	// aesgcm, err := cipher.NewGCM(block)
	// if err != nil {
	// 	panic(err.Error())
	// }

	// // Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	// nonce := make([]byte, 12)

	// if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
	// 	panic(err.Error())
	// }

	// ciphertext := aesgcm.Seal(nonce, nonce, text, sum[:])

	// ciphertext = append(ciphertext, sum[:]...)

	// // write encrypted data to file
	// err = ioutil.WriteFile("./encrypted/"+fileName, ciphertext, 0777)

}
