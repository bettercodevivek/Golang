package main

import "fmt"

type person struct {
	name    string
	age     int
	address string
	orders  int
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

}
