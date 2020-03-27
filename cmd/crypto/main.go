package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Gopher")

	fileNames := getFiles()

	fmt.Println(fileNames)

	mkDirEnc()

	encs := getEncrypted()
	fmt.Println(encs)

	mkDirDec()

	for _, name := range fileNames {
		encrypt(name)
	}
	for _, name := range encs {
		decrypt(name)
	}
}
