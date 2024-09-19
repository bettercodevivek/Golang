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

	var e, f, g, h int = 2, 4, 6, 8

	fmt.Println(e, f, g, h)

	// Multiple variable declarations can also be grouped together into a block for greater readability:

	var (
		a1 string = "hello"
		a2 string = "my name is"
		a3 string = "sully"
	)
	fmt.Println(a1 + a2 + a3)

	// var name_1 string = "jimmy"

	// var name_2 string = "stuart"

	// var name_3 string = "mitchell"

	// fmt.Println(name_1, name_2, name_3)

	var (
		name_1 string = "jimmy"

		name_2 string = "mitchell"

		name_3 string = "stuart"
	)

	fmt.Println("hello " + name_1 + " I am " + name_2 + " and this is " + name_3)

	//  If a variable should have a fixed value that cannot be changed, you can use the const keyword.

	//  The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.

	const PI float32 = 3.14

	fmt.Println(PI)

	// There are 2 types of constants :- typed constants and untyped constants

	// typed --- in which you declare a variable with defined type

	// untyped --- in which you don't declare variable with defined type

	var var_1 string = "virat"

	fmt.Printf("%T", var_1)

	/* The Printf() function first formats its argument based on the given formatting verb and then prints them.

	Here we will use two formatting verbs:

	    %v is used to print the value of the arguments
	    %T is used to print the type of the arguments
	*/

}
