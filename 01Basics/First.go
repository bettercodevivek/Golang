package main

import (
	"fmt"
)

var myname2 string = "jane doe"

// myname3 := "joe doe"

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

	//It is not possible to declare a variable using :=, without assigning a value to it.

	var c int
	c = 60

	d := 80

	fmt.Println(c, d)

	var var1 string
	var var2 int
	var var3 bool
	var var4 int16
	var var5 int32
	var var6 int64
	var var7 int8

	fmt.Println(var1, var2, var3, var4, var5, var6, var7)

	fmt.Println(myname2)

	// another difference between var and := is that var declared  variables can be declared outside and inside the function.

	// whereas := declared variables can only be declared within a function
}
