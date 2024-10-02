package main

import (
	"fmt"
	"os"
)

func main() {

	// defer fmt.Println("and this is the 1st defer statement")
	// defer fmt.Println("this is the 2nd defer statement")
	// defer fmt.Println("this is the 3rd defer statement")
	// fmt.Println("my name is vivek and i am learning golang !")

	// Corrected use case of defer for closing the file
	file, err := os.OpenFile("demo.txt", os.O_RDWR|os.O_APPEND, 0644) // Open file with read and write permissions
	if err != nil {
		fmt.Println("error occurred in opening file:", err)
		return
	}

	// Ensure the file is closed at the end of the function
	defer file.Close()

	// Write to the file before closing it
	_, err = file.WriteString("The file has been updated now\n")
	if err != nil {
		fmt.Println("error writing to the file:", err)
		return
	}

	fmt.Println("File opened and updated successfully!")
}
