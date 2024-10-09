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
}
