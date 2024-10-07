package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
	orders  int
}

type struct1 struct {
	name    string
	age     int
	male    bool
	address string
}

// You can also pass a structure as a function argument, like this:

func SayHello(pers person) {
	a := "Hello, Mr. " + pers.name + " how are you this evening ?"
	fmt.Println(a)
}

func main() {
	// let us study about structs in go.
	/* A struct (short for structure) is used to create a collection of members of different data types, into a single variable.

	While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.

	A struct can be useful for grouping data together to create records.*/

	//  var person2 person ;

	// person1.name = "vivek"
	// person1.age = 20
	// person1.address = "san francisco"
	// person1.orders = 40

	person1 := person{"vivek", 30, "pasadena", 50}

	//With literal notation, you provide values for the fields in the order they are defined.

	person2 := person{
		name:    "johnny",
		age:     30,
		address: "san francisco",
		orders:  100,
	}

	fmt.Println(person1, " ", person2)

	// With named field initialization, you explicitly specify which value corresponds to which field. This can improve readability, especially for larger structs.

	// Fields are accessed and modified using the dot . operator.

	person2.name = "david fincher"

	fmt.Println(person2)

	defer SayHello(person1)

	SayHello(person2)

	man1 := struct1{"josh", 21, true, "india"}

	man2 := struct1{
		name:    "bale",
		age:     24,
		male:    true,
		address: "delhi",
	}

	fmt.Print(man1, "\n")

	fmt.Println(man2)

}
