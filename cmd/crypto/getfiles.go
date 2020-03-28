package main

import (
	"fmt"
	"io/ioutil"
)

// getFiles returns the file information in dir: "files"
func getFiles() (names []string) {

	files, err := ioutil.ReadDir("./files")
	if err != nil {
		fmt.Println("Reading file information failed! Program is terminated! Error: ", err)
		return
	}

	for _, file := range files {
		n := file.Name()
		names = append(names, n)
	}
	return
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
