package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("practice.txt")
	if err != nil {
		fmt.Println("An error occurred in creating the file named : ", file.Name(), err)
	}

	file.WriteString("THIS IS MY PRACTICE FILE FOR LEARNING FILE HANDLING IN GOLANG")

	defer file.Close()

	fmt.Println("file Create and updated successfully !!!")

	file1, err := os.OpenFile("practice.txt", os.O_APPEND|os.O_RDWR, 0644)

	if err != nil {
		fmt.Println("Error occurred in opening a file", err)
	}

	file1.WriteString("\nFirst Ammendment here boss!!!")

	data, err := os.ReadFile("practice.txt")

	if err != nil {
		fmt.Println("error occurred in reading file", err)
	}

	fmt.Println("data in the file is :- \n", string(data))

	buffer := make([]byte, 100)

	bytesread, err := file.Read(buffer)

	fmt.Printf("no. of bytes in file are:- %d", bytesread)

	fmt.Println("content of file is:-", string(buffer[:bytesread]))

	details, err := os.Stat("practice.txt")

	if err != nil {
		fmt.Println("Error occured in showing file details", err)
	}

	fmt.Println("file name :-", details.Name())
	fmt.Println("file name :-", details.Name())
	fmt.Println("file name :-", details.Name())
	fmt.Println("file name :-", details.Name())
	fmt.Println("file name :-", details.Name())
}
