package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main()  {
	fmt.Println("Hello, Gopher!")

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Something went wrong! - I will return from the function!")
		return
	}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println("Something went wrong! - I will return from the function!")
		return
	}
	fmt.Printf("%s \n", cwd)
	fmt.Println(files[1].Name())
	for i :=0; i<len(files); i++ {
		if files[i].Name() != "main.go" {
			fmt.Println(files[i].Name())
		}
		
	}

}