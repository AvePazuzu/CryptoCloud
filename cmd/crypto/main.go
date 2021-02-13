package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {

	// set start
	// t1 := time.Now()
	fmt.Println("\nWelcome to MyFileLocker!\n")

	// setupDir() checks directory for missing folders and creates them
	setupDir()

	fmt.Println("\nSelect from the following options & confrim with \"Enter\":\n",
		"[1]: Encrypt\n",
		"[2]: Decrypt\n",
		"[3]: Compare\n",
		"[4]: Help\n",
		"[5]: Exit")

	for {
		// Read input from console
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\nEnter the your number: ")

		input, err := reader.ReadString('\n')

		// To make input comparable delimiter needs to be removed
		input = strings.Replace(input, "\n", "", -1)

		if err != nil {
			fmt.Println(err)
			return
		}

		switch input {
		case "1":
			fmt.Println("\nData to encrypt:")
			encrypt()
			continue
		case "2":
			fmt.Println("2 has been entered")
			continue
		case "3":
			fmt.Println("3 has been entered")
			continue
		case "4":
			fmt.Println("4 has been entered")
			continue
		case "5":
			fmt.Println("\nExiting program... See you soon!\n")
			time.Sleep(2 * time.Second)
			return
		default:
			fmt.Println("Enter valid input.")
			continue
		}
		break

	}

	fmt.Println("Returned as needed!")

	return
}

// 	// dir name for data to encrypt
// 	// if setupDir() != true {
// 	// 	return
// 	// }
// 	fdata := "files"

// 	dir := "./" + fdata

// 	// create session log
// 	sessionLog := make([]byte, 0)
// 	// append time stamp
// 	sessionLog = append(sessionLog, "Date: "+t1.Format("02.01.2006 15:04:05")...)
// 	// append operating system
// 	sessionLog = append(sessionLog, "\n"+"OS: "+runtime.GOOS...)

// 	fmt.Printf("\nChecking data in %q ... ", fdata)

// 	// retrieve the files to encrypt
// 	files, err := ioutil.ReadDir(dir)
// 	if err != nil {
// 		fmt.Println("Reading file information failed! Program is terminated! Error: ", err)
// 		return
// 	}
// 	// if exists get & print out number of data objects to encrypt
// 	if len(files) > 0 {
// 		fmt.Printf("%q contains %d data object(s).\n", fdata, len(files))

// 	} else {
// 		fmt.Println("There is no data to encrypt. Program is terminating.")
// 		return
// 	}

// 	mkDir("encrypted")

// 	fmt.Printf("\nEncryption started ... ")

// 	for _, name := range files {
// 		encrypt(name.Name())
// 	}

// 	fmt.Printf("%d data objects(s) successfully encrypted.\n", len(files))

// 	// retrieve the files to decrypt
// 	encs := getEncrypted()

// 	mkDir("decrypted")

// 	for _, name := range encs {
// 		decrypt(name)
// 	}

// 	// compHash()
// 	compDirect()

// set end, calculate the elapsed time and append it to sessionLog
// 	t2 := time.Now()
// 	elapsed := t2.Sub(t1)
// 	fmt.Println(elapsed)
// 	sessionLog = append(sessionLog, "\n"+"Elapsed Time: "+elapsed.String()...)
// }

// 	// create dir for and write session log to file
// 	// '0644' is a Unix code for file permissions
// 	// mkDir("sessions")
// 	// err = ioutil.WriteFile("./sessions/"+t1.Format("02.01.2006_15:04:05")+".txt", sessionLog, 0644)
// 	// if err != nil {
// 	// 	fmt.Println(err)
// 	// 	return
// 	// }
// 	// fmt.Println("Press 'Enter' to Exit.")
// 	// if true {
// 	// 	bufio.NewReader(os.Stdin).ReadBytes('\n')
// 	// 	return
// 	// }
// 	// exit()
// }
