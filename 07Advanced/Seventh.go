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

	var person1 person
	//  var person2 person ;

	person1.name = "vivek"
	person1.age = 20
	person1.address = "san francisco"
	person1.orders = 40

	fmt.Println(person1)

}
