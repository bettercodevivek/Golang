package main

import "fmt"

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func StringPrinter(str1 string, i int) string {
	if i >= len(str1) {
		return " "
	}
	fmt.Printf("%c", str1[i])

	return StringPrinter(str1, i+1)
}

func main() {
	// Let us try to understand recursion using the factorial example
	fmt.Println(factorial(5))

	var str string = "golang"

	fmt.Println(StringPrinter(str, 0))

}
