package main

import (
	"fmt"
	"log"
	"os"
)

// Go provides a powerful package called os to work with files

func main() {

	// 1. You can create a file using the os.Create() function. This will either create a new file or truncate the file if it already exists.

	file, err := os.Create("demo.txt")

	if err != nil {
		fmt.Println("An Error occured while creating the file => ", err)
	}

	fmt.Println("File Create Successfully Boss with the following name => ", file.Name())

	// 2. If you need to open a file that already exists, use os.Open() or os.OpenFile() for more specific options.

	// The os.OpenFile() function gives more control over the file’s read/write permissions. This is useful when you need to specify flags for file access, like append mode, read-only, etc.

	file1, err1 := os.OpenFile("demo.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)

	if err1 != nil {
		fmt.Println("An error occured while opening the file :- ", err1)
	}

	file1.WriteString("This is me using openfile function to write to this demo file after its been created ")

	fmt.Println("File opened and modified successfully !!")

	// 3. You can read a file in Go using a few different methods:

	// Reading Entire File at Once

	data, err := os.ReadFile("demo.txt")

	if err != nil {
		fmt.Println("Error occurred while reading the file => ", err)
	}

	fmt.Println("File Content:-\n", string(data))
	// conversion of data to string is must to make it in human readable form as the readfile function returns the data of a file in  []byteslice form which basically means ASCII code for each character.

	// Reading with a Buffer

	// To read data in chunks, you can use the Read() method on the file object, which reads the file into a byte slice (buffer).

	buffer := make([]byte, 100) // Buffer to hold chunks of file data
	bytesRead, err := file.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Bytes read: %d\n", bytesRead)
	fmt.Println("Data read from file:\n", string(buffer[:bytesRead]))

	// 4. The Stat() method on a file or os.Stat() on a filename provides details about a file, such as its size, permissions, and modification time.

	filedetails, err := os.Stat("demo.txt")

	if err != nil {
		fmt.Println("An error occurred while fetching file details => ", err)
	}

	fmt.Println("Filename : ", filedetails.Name())
	fmt.Println("FileSize : ", filedetails.Size())
	fmt.Println("File Last Modified : ", filedetails.ModTime())
	fmt.Println("File Permissions : ", filedetails.Mode())

	// 5. You can delete a file using os.Remove().

	// os.Remove("demo.txt")

	// fmt.Println("file deleted successfully")

	defer file.Close()
}
