package main

import "fmt"

func main() {

	fmt.Println("Hello, Gopher!")

}
3. Writes file names in a .txt
// 4. Saves .txt in new folder
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	// Loop over the files in path return the len of each file
	// and add it to total; +1 as additional place for line separator '\n'

	// Create new Folder in dir
	dirName := "newdir"
	path := "./" + dirName
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0744)
		fmt.Println("newdir successfully created")
	} else {
		fmt.Println("newdir already exsists")
	}
	// Read Files in dir
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Something went wrong! - I will return from the function!")
		return
	}

	// To optimize the program the backing array is created only once
	// Therefore the size of the backing array needs to be calculated
	// var total is the capacity of the backing array

	var total int
	for _, file := range files {
		if file.Name() != "main.go" || file.Name() != dirName {
			total += len(file.Name()) + 1
		}
	}

	// names var is declared wiht the make func
	// only bytes are written to files
	names := make([]byte, 0, total)
	// Loop over the files, check for size and append the names
	for _, file := range files {
		if file.Name() != "main.go" && file.Name() != dirName {
			n := file.Name()
			names = append(names, n...)
			// '\n' seperator needed to separate the bytes by line
			names = append(names, '\n')
		}
	}
	// '0644' is a Unix code for file permissions
	err = ioutil.WriteFile(dirName+"/out.txt", names, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// for _, v := range names {
	// 	fmt.Println(v.Name())
	// }

	// Declare a fileNames var of type slice by using the make func of len(files-1)
	// fileNames := make([]string, len(files)-1)
	// fmt.Printf("%#v \n", fileNames)

	// for i := 0; i < len(files)-1; i++ {
	// 	if files[i].Name() != "main.go" {
	// 		fileNames[i] = files[i].Name()
	// 	}
	// }
	// fmt.Printf("%#v \n", names)

	// Create .txt file to write the fileNames to:
	var txt = "File_Names.txt"

	// check if file exists
	// var _, err = os.Stat(txt)

	// create file if not exists
	// if os.IsNotExist(err) {
	// 	var newFile, err = os.Create(txt)
	// 	if isError(err) {
	// 		return
	// 	}
	// 	defer newFile.Close()
	// }

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

// func isError(err error) bool {
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	return (err != nil)
// }
