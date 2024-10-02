// LET US NOW STUDY ABOUT DEFER KEYWORD IN GOLANG

/*The defer keyword is used to postpone the execution of a function until the surrounding function returns. It’s often used for cleanup tasks,
like closing files or releasing resources, to ensure they are handled properly even if an error occurs in the function. */

package main

import "fmt"

func main() {

	defer fmt.Println("and this is the first defer statement")

	fmt.Println("my name is vivek and i am learning golang !")

}
