package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

func main() {

	// set start
	t1 := time.Now()

	// create session log
	sessionLog := make([]byte, 0)
	// append time stamp
	sessionLog = append(sessionLog, "Date: "+t1.Format("02.01.2006 15:04:05")...)
	// append operating system
	sessionLog = append(sessionLog, "\n"+"OS: "+runtime.GOOS...)

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

	// set end, calculate the elapsed time and append it to sessionLog
	t2 := time.Now()
	elapsed := t2.Sub(t1)
	fmt.Println(elapsed)
	sessionLog = append(sessionLog, "\n"+"Elapsed Time: "+elapsed.String()...)

	// write session log to file
	// '0644' is a Unix code for file permissions
	err := ioutil.WriteFile("./sessionlog.txt", sessionLog, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
