package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// setupDir scans current dir for "files". It is used as the destination for data to encrypt.
// The program returns if it needs to be created or sets var s=true to continune main.
func setupDir() (s bool) {
	dirName := "files"
	path := "./" + dirName
	fmt.Printf("\nCheking directory ... ")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0744)
		fmt.Printf("Folder %q is missing and needs to be created.\n", dirName)
		fmt.Printf("\nFolder %q successfully created.\n", dirName)
		fmt.Printf("\nPlease copy data into %q and restart the program.\n", dirName)
		exit()
	} else {
		fmt.Printf("Folder %q already exsists.\n", dirName)
		s = true
	}
	return s
}

// mkDir scans dir for desired folder and creates it if necessary
func mkDir(dirName string) {

	path := "./" + dirName
	fmt.Printf("\nCheking directory ... ")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0744)
		fmt.Printf("Folder %q successfully created.\n", dirName)
	} else {
		fmt.Printf("Folder %q already exsists.\n", dirName)
	}
}

// getEcnypted retruns the file information in dir: "encrypted"
func getEncrypted() (encs []string) {

	files, err := ioutil.ReadDir("./encrypted")
	if err != nil {
		fmt.Println("Reading encrypted file information failed! Program is terminated! Error: ", err)
		return
	}

	for _, file := range files {
		n := file.Name()
		encs = append(encs, n)
	}
	return
}

// exit() returns from the function
func exit() {
	fmt.Printf("\nPress %q to Exit.", "Enter")
	var input string
	fmt.Scanln(&input)
	return
}

// wrErr function to write out all errors. It takes the error if
// occures and writes the error message to file.
func wrErr(err error) {
	t := time.Now()
	fmt.Println("Writing error message to file.")
	mkDir("errors")
	err = ioutil.WriteFile("./errors/"+t.Format("02.01.2006_15:04:05")+".txt", err, 0644)
}
