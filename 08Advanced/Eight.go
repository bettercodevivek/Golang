package main

import "fmt"

func main() {

	/* Maps are used to store data values in key:value pairs.

	Each element in a map is a key:value pair.

	A map is an unordered and changeable collection that does not allow duplicates.

	The length of a map is the number of its elements. You can find it using the len() function.

	The default value of a map is nil.

	Maps hold references to an underlying hash table.

	Go has multiple ways for creating maps. */

	var map1 = map[string]string{"name": "nexon", "brand": "tata"}

	fmt.Println(map1)

	map2 := map[int]string{1: "vivek", 2: "john", 3: "dave", 4: "batista"}

	fmt.Println(map2)

	// let us now see how to create a map using make() function

	var map3 = make(map[string]string)
	map3["name"] = "sam"
	map3["name"] = "montana"
	map3["origin"] = "asia"
	map3["gender"] = "male"

	fmt.Println(map3)

	// The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.

	// Accessing map elements

	fmt.Println(map2[1])

	// map ka name and then in square brackets key name jiski value access karni hai.

	map2[1] = "holden"

	fmt.Println(map2)

	// Removing elements is done using the delete() function.

	delete(map2, 2)

	fmt.Println(map2)
}
