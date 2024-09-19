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

	// In go, you can enclose string variables in double quotes only, single quotes will give error

}
