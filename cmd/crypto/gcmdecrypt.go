package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
)

// func mkDirDec() {
// 	dirName := "decrypted"
// 	path := "./" + dirName
// 	if _, err := os.Stat(path); os.IsNotExist(err) {
// 		os.Mkdir(path, 0744)
// 		fmt.Println("decrypted successfully created")
// 	} else {
// 		fmt.Println("decrypted already exsists")
// 	}
// }

func decrypt(fileName string) {

	ciphertext, err := ioutil.ReadFile("./encrypted/" + fileName)
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

	err = ioutil.WriteFile("./decrypted/"+fileName, plaintext, 0777)

	fmt.Printf("add is  : %x\n", add)
	fmt.Printf("plain is: %x\n", sha256.Sum256(plaintext))
	sum := sha256.Sum256(plaintext)
	fmt.Printf("Equal: %t\n", bytes.Equal(sum[:], add))
	dat2, err := ioutil.ReadFile("./IT_Flye.pdf")
	dat1, err := ioutil.ReadFile("/IT_Flye.pdf")
	// sumDat2 := sha256.Sum256(dat2)
	fmt.Printf("Equal: %t\n", bytes.Equal(dat2, dat1))
}
