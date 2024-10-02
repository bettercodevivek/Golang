package main

import "fmt"

func main() {
	// LET US NOW TRY TO UNDERSTAND THE CONCEPT OF POINTERS

	var a int = 20

	var ptr *int = &a

	fmt.Println(*ptr)

	//If you have a pointer, you can modify the original value by dereferencing the pointer.

	*ptr = 50

	fmt.Println("new value is :-", *ptr)
}
