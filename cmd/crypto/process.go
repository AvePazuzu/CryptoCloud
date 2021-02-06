package main

import (
	"fmt"
	"io/ioutil"
	"os"
	// "time"
)

// slice of required dirs
var dirs = []string{"data_to_encrypt", "encrypted_data", "decrypted_data"}

// setupDir scans current dir for the necessary folders and creates them if they are missing.
func setupDir() {

	path := "./"
	fmt.Println("Cheking directory ... ")

	for _, dir := range dirs {
		if _, err := os.Stat(path + dir); os.IsNotExist(err) {
			os.Mkdir(path+dir, 0744)
			fmt.Printf("\nFolder \"%s\" created.\n", dir)
		}
	}
	fmt.Println("\nDirectory is set up... You are ready to lock...\n")

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
//func wrErr(err error) {
//	t := time.Now()
//	fmt.Println("Writing error message to file.")
// mkDir("errors")
// err = ioutil.WriteFile("./errors/"+t.Format("02.01.2006_15:04:05")+".txt", err, 0644)
// }
