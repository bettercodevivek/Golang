package main

import "fmt"

func main() {
	var num1 int = 20

	var adr *int = &num1

	fmt.Println(num1)

	*adr = num1 + 55

	fmt.Println(*adr)

}
