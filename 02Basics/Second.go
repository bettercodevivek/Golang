package main

import "fmt"

func main() {

	// There are 3 ways of declaring an array in golang ;

	// First => using var keyword with fixed length

	var arr1 = [10]int{1, 2, 3, 4, 5}
	fmt.Println(arr1)

	//  Second => using [...] instead of fixed length which means the compiler itself will infer the length

	var arr2 = [...]string{"john", "ron", "cohn"}
	fmt.Println(arr2)

	// Third => without the var keyword using only :=

	arr3 := [...]string{"lois", "clark", "arkham"}
	fmt.Printf("value is %v and type is %T\n", arr3, arr3)

	// We can also access array elements using their indices.

	fmt.Println(arr2[0])

	// we can also initialise arrays in golang in the following manner

	var arr4 = [5]string{} // not initialized

	var arr5 = [5]int{1, 2, 3} // Partial initialization

	fmt.Println(arr4, arr5)

	// for all array elements whose values have not been initialised, in go they have default values assigned. for int its 0 , for bool its false , for string its " "

	// in Golang you can initialise specific elements of the  array in the following manner :-

	var arr6 = [5]int{1: 10, 3: 50, 0: 5}

	fmt.Println(arr6)

	// Let us now see slices in golang which are similar to arrays

	/* Slices are similar to arrays, but are more powerful and flexible.

	   Like arrays, slices are also used to store multiple values of the same type in a single variable.

	   However, unlike arrays, the length of a slice can grow and shrink as you see fit. */

}
