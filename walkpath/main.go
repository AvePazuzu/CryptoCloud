package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	files, err0 := ioutil.ReadDir(".")
	if err0 != nil {
		fmt.Println("Something went wrong! - I will return from the function!")
		return
	}

	// Declare a fileNames var of type slice by using the make func of len(files-1)
	fileNames := make([]string, len(files)-1)
	fmt.Printf("%#v \n", fileNames)

	for i := 0; i < len(files); i++ {
		if files[i].Name() != "main.go" {
			fileNames[i] = files[i].Name()
		}

	}
	fmt.Printf("%#v \n", fileNames)

	// Create .txt file to write the fileNames to:
	var txt = "File_Names.txt"

	// check if file exists
	var _, err = os.Stat(txt)

	// create file if not exists
	if os.IsNotExist(err) {
		var newFile, err = os.Create(txt)
		if isError(err) {
			return
		}
		defer newFile.Close()
	}

	fmt.Println("File Created Successfully", txt)

}

// txt is only declare for the func main() scope--> cannot be used for other functions:

// func deleteFile() {
// 	// delete file
// 	var err = os.Remove(txt)
// 	if isError(err) {
// 		return
// 	}
// }

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}
