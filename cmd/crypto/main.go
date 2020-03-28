package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println(time.Now())

	// set start
	t1 := time.Now()

	// retrieve the files to encrypt
	fileNames := getFiles()

	fmt.Println(fileNames)

	mkDirEnc()

	// retrieve the files to decrypt
	encs := getEncrypted()
	fmt.Println(encs)

	for _, name := range fileNames {
		encrypt(name)
	}

	mkDirDec()

	for _, name := range encs {
		decrypt(name)
	}

	// set end and calculate the elapsed time
	t2 := time.Now()
	elapsed := t2.Sub(t1)
	fmt.Println(elapsed)
	fmt.Println(runtime.GOOS)
}
