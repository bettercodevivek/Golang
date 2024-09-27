package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	// Let us try to understand recursion using the factorial example
	fmt.Println(factorial(5))
}
