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
}
