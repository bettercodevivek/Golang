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

func Stringer(message1 string, message2 string) string {
	return message1 + message2
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
	result := Stringer("hello", "i am learning golang")
	fmt.Println(result)

	// In golang, when parameters are defined in a function their datatype needs to specified too else it will throw an error.

	// If using return keyword in golang then you have to specify the return type of the function explicitly just before the {} of function

}
