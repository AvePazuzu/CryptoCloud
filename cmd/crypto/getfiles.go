package main

import (
	"fmt"
	"io/ioutil"
)

func getFiles() (names []string) {

	// Read Files in dir
	files, err := ioutil.ReadDir("./files")
	if err != nil {
		fmt.Println("Reading files failed! Program is terminated! Error: ", err)
		return
	}

	for _, file := range files {
		n := file.Name()
		names = append(names, n)
	}
	return
}

func getEncrypted() (encs []string) {
	// Read Files in in folder "encrypted"
	files, err := ioutil.ReadDir("./encrypted")
	if err != nil {
		fmt.Println("Reading encrypted files failed! Program is terminated! Error: ", err)
		return
	}

	for _, file := range files {
		n := file.Name()
		encs = append(encs, n)
	}
	return
}
