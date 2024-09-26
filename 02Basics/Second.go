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

	myslice1 := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(myslice1)

	/* In Go, there are two functions that can be used to return the length and capacity of a slice:

	   len() function - returns the length of the slice (the number of elements in the slice)
	   cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

	*/

	// creating a slice from an array

	myslice2 := arr2[0:2]

	fmt.Println(myslice2)

	// can also create slice using the make() function

	// syntax =>  slice_name := make([]type, length, capacity)

	myslice3 := make([]string, 4, 6)
	myslice3[0] = "vivelk"
	myslice3[1] = "vivelk"
	myslice3[2] = "vivelk"
	myslice3[3] = "vivelk"
	fmt.Println(myslice3)

	// So in the above case, capacity jo hai hamara slice ki woh 6 hai but length is 4, lekin 4 se zyda elements daalne mein index out of range
	// error show ho raha hai, why is that???

	// BECAUSE When you try to access myslice3[4] and myslice3[5] directly, it raises an "index out of range" error because those indices are beyond the current length of the slice.
	//  Even though the slice has a capacity of 6, Go slices only allow you to access elements up to the current length. So in this case, your slice has a length of 4, which means valid indices are 0 to 3.

	// TOH ISKA SOLUTION ??
	// If you want to add elements beyond the initial length (4), you should use the append() function to grow the slice.

	myslice3 = append(myslice3, "vivekf", "vivekn")
	fmt.Println(myslice3)
}
