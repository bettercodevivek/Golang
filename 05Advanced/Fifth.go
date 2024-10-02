// LET US NOW STUDY ABOUT DEFER KEYWORD IN GOLANG

/*The defer keyword is used to postpone the execution of a function until the surrounding function returns. It’s often used for cleanup tasks,
like closing files or releasing resources, to ensure they are handled properly even if an error occurs in the function. */

package main

import "fmt"

func main() {

	defer fmt.Println("and this is the 1st defer statement")

	defer fmt.Println("this is the 2nd defer statement")

	defer fmt.Println("this is the 3nd defer statement")

	// it can be clearly seen that in case of multiple defer statements , the defer statements follow LIFO. therefore last statment will get executed first.

	fmt.Println("my name is vivek and i am learning golang !")

}
