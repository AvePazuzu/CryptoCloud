package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

// compHash takes hash from additional data
// and compares is to the hash of the decrypted data checks if data before crypto process and after is the same
// Here, data is comapared directly and via comparing the hashes
func compHash() {
	fmt.Println("Hello, Gopher!")
}

// compDirect loads original data into memory and directly compares it with the decrypt plaintext.
func compDirect() {
	// get data from "files"
	files, err := ioutil.ReadDir("./files")
	if err != nil {
		fmt.Println("Reading file information failed! Program is terminated! Error: ", err)
		return
	}

	// get data from "decrypted"
	decs, err := ioutil.ReadDir("./decrypted")
	if err != nil {
		fmt.Println("Reading file information failed! Program is terminated! Error: ", err)
		return
	}

	if len(files) != len(decs) {
		fmt.Println("Uneven amount of files to compare.")
		return
	}

	for _, name1 := range files {
		for _, name2 := range decs {
			if name1.Name() == name2.Name() {
				fmt.Printf("Files to compare: %s and %s\n", name1.Name(), name2.Name())
				dat1, err := ioutil.ReadFile("./files/" + name1.Name())
				if err != nil {
					fmt.Println("Something went wrong reading original file.")
					return
				}

				dat2, err := ioutil.ReadFile("./decrypted/" + name2.Name())
				if err != nil {
					fmt.Println("Something went wrong reading original file.")
					return
				}

				fmt.Printf("Equal: %t\n", bytes.Equal(dat1, dat2))

			}
		}
	}

}
