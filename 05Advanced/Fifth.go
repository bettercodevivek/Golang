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
	vivbyte, err := file.WriteString("The file has been updated now and it will further get updated too\n")

	/* _ (underscore): In Go, the underscore is known as the blank identifier. It's used to discard a value that you're not interested in. In this case, WriteString returns two values:

	    The number of bytes written to the file.
	    An error value.

	Since you don't need to know how many bytes were written, you use _ to discard that value. */

	if err != nil {
		fmt.Println("error writing to the file:", err)
		return
	}

	fmt.Println("File opened and updated successfully!", vivbyte)
}

/*     0: The leading 0 indicates that the number is written in octal (base 8).
6: The owner (user) of the file has read (4) and write (2) permissions. 4 + 2 = 6.
4: The group has read permission (4), but no write or execute permissions.
4: Others (everyone else) have read permission (4), but no write or execute permissions.

  Meaning of 0644:

The owner of the file can read and write to the file.
The group and others can only read the file, but they cannot write or execute it.

*/
