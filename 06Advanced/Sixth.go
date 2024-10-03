/*Pointer: It's a variable that stores the memory address of another variable.
 In simple words, it points to where some data is stored in the computer's memory.
 Pointers are useful because they allow programs to:

    Access and modify data efficiently: Instead of copying large amounts of data, you can just pass a pointer to the data, which saves memory and speeds up things.
    Share data: Multiple parts of a program can work with the same data without creating multiple copies of it.
	p := &x means the pointer p is storing the memory address of x (the location where x is stored).
*p is called dereferencing the pointer, which means getting the value stored at the memory location p points to
*/

package main

import "fmt"

func main() {
	var num1 int = 20

	var arr1 = [...]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr1); i++ {
		adr1 := &arr1[i]
		fmt.Println(*adr1)
	}

	var adr *int = &num1

	fmt.Println(num1)

	*adr = num1 + 55

	fmt.Println(*adr)

	fmt.Println(arr1)

}
