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
	fmt.Println("\nProgram starts ...")

	// dir name for data to encrypt
	fdata := "files"
	mkDir(fdata)
	dir := "./" + fdata

	// create session log
	sessionLog := make([]byte, 0)
	// append time stamp
	sessionLog = append(sessionLog, "Date: "+t1.Format("02.01.2006 15:04:05")...)
	// append operating system
	sessionLog = append(sessionLog, "\n"+"OS: "+runtime.GOOS...)

	fmt.Printf("\nChecking data in %q ... ", fdata)

	// retrieve the files to encrypt
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println("Reading file information failed! Program is terminated! Error: ", err)
		return
	}
	// get & print out number of data objects to encrypt
	fmt.Printf("%q contains %d data object(s)\n", fdata, len(files))

	mkDir("encrypted")

	fmt.Printf("\nEncryption started ... ")

	for _, name := range files {
		encrypt(name.Name())
	}

	// retrieve the files to decrypt
	encs := getEncrypted()

	mkDir("decrypted")

	for _, name := range encs {
		decrypt(name)
	}

	// set end, calculate the elapsed time and append it to sessionLog
	t2 := time.Now()
	elapsed := t2.Sub(t1)
	fmt.Println(elapsed)
	sessionLog = append(sessionLog, "\n"+"Elapsed Time: "+elapsed.String()...)

	// create dir for and write session log to file
	// '0644' is a Unix code for file permissions
	mkDir("sessions")
	// err = ioutil.WriteFile("./sessions/"+t1.Format("02.01.2006_15:04:05")+".txt", sessionLog, 0644)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

}
