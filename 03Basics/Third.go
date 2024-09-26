// BASICALLY, IF YOU WANT TO WRITE FUNCTIONS NORAMLLY IN GO THEN THEY CAN ONLY BE DEFINED OUT OF THE MAIN FUNCTION, IF YOU TRY TO DO THAT WITHIN THE MAIN FUNCTION IT THROWS ERROR.
// SO, WHAT TO DO? WELL YOU CAN DEFINE A ANONYMOUS FUNCTION IN MAIN AND ASSIGN IT TO A VARIABLE
package main

import "fmt"

func Adder(num1 int, num2 int) {
	if num1 > 10 && num2 > 10 {
		fmt.Println(num1 + num2)
	} else {
		fmt.Println("num1 and num2 should be greater than 10")
	}
}

func main() {

	myfunc := func() {
		fmt.Println("this is a function inside the main function")
	}

	// func Myfunc(){
	// 	fmt.Println("Hello i am a new function !!!")
	// }

	// Myfunc

	myfunc()
	Adder(2, 4)
}
