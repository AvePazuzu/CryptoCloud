// package main

// import (
// 	"bytes"
// 	"crypto/sha256"
// 	"fmt"
// )

// func createKey() [32]byte {
// 	passPhrase := "Some Pass Phrase"
// 	key := sha256.Sum256([]byte(passPhrase))
// 	return key
// }

// func getFiles() []byte {
// 	plain := []byte("exampleplaintext")
// 	return plain
// }

// func encrypt() {
// 	block := cipher.NewGCM()
// }

// func main() {

// 	key := createKey()

	// plain := getFiles()

	// text := bytes.Repeat([]byte("i"), 96)

	// fmt.Println(len(text))

	// iv := make([]byte, aes.BlockSize)
	// if _, err := rand.Read(iv); err != nil {
	// log.Fatal(err)
	// }

	// fmt.Printf("%x \n", key)
	// fmt.Printf("%x \n", plain)
// }