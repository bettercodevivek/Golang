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
