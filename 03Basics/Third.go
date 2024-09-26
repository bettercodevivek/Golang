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

// Let us see a function that can return multiple values at once

func Divide(num1 int, num2 int) (int, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return num1 / num2, nil
}

// Let us see one more function with multiple return values

func Random(val1 int, val2 int) (int, string) {
	return val1 + val2, "Hello this is the compiler"
}

func Named(val1 int, val2 int) (sum int, prod int) {
	sum = val1 + val2
	prod = val1 * val2

	return
}

// LET US SEE VARIADIC FUNCTIONS NOW
// Go allows functions to accept variadic parameters, which means the function can take an arbitrary number of arguments of the same type.

func variadicexmp(nums ...int) int {
	var total int = 0
	for _, num := range nums {
		total += num
	}
	return total
}

// The ... syntax means the function can accept any number of integers as arguments.

// Inside the function, nums behaves like a slice ([]int), and you can loop over it to process the values.

func variadicexmp2(words ...string) string {
	sentence := " "
	for _, word := range words {
		sentence += word + " "
	}
	return sentence
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

	fmt.Println(Divide(45, 0))

	fmt.Println(Random(34, 56))

	sum, prod := Named(40, 50)

	fmt.Println("Sum is", sum)
	fmt.Println("Product is", prod)

	fmt.Println(variadicexmp(1, 2, 3, 4, 5))
	fmt.Println(variadicexmp(1, 2, 3, 4, 5, 45, 65, 6, 3, 34, 35, 5464, 5, 54, 645, 4, 3, 46))

	fmt.Println(" this is the second variadic function example", variadicexmp2("Hello", "my", "name", "is", "vivek"))
}
