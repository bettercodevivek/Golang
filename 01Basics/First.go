package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world. i am learning golang!")
	fmt.Println("This is 2nd statement")

	// Comments are written in go using double slash and multiline comments can be written in the following manner:-

	/* This is a multi line comment
	   This is the syntax for it
	*/

	// In Golang you can declare variables in the following 2 manners :-

	// 1 . var variablename type = value

	var myname string = "vivek"
	fmt.Println(myname)

	var a int = 40
	var b int = 80

	fmt.Println(a * b)

	// In go, you can enclose string variables in double quotes only, single quotes will give error

	// 2. using := sign ,  variablename := value

	// In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

	myname1 := "john doe"

	fmt.Println(myname1)
}
