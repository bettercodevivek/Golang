package main

import (
	"fmt"
	"os"
)

// Go provides a powerful package called os to work with files

func main() {

	// 1. You can create a file using the os.Create() function. This will either create a new file or truncate the file if it already exists.

	file, err := os.Create("demo.txt")

	if err != nil {
		fmt.Println("An Error occured while creating the file => ", err)
	}

	defer file.Close()

	fmt.Println("File Create Successfully Boss with the following name => ", file.Name())

	// If you need to open a file that already exists, use os.Open() or os.OpenFile() for more specific options.

	// The os.OpenFile() function gives more control over the file’s read/write permissions. This is useful when you need to specify flags for file access, like append mode, read-only, etc.

	file1, err1 := os.OpenFile("demo.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err1 != nil {
		fmt.Println("An error occured while opening the file :- ", err1)
	}

	file1.WriteString("This is me using openfile function to write to this demo file after its been created ")

	fmt.Println("File opened and modified successfully !!")
}
