package main

import "fmt"

func main() {
	//Pointers

	/* var firstNamePtr *string = new(string)

	*firstNamePtr = "Jonas"

	fmt.Println(*firstNamePtr) */

	firstName := "Jonas"
	fmt.Println(firstName)

	firstNamePtr := &firstName
	fmt.Println(firstNamePtr, *firstNamePtr)

	firstName = "Niklas"
	fmt.Println(firstNamePtr, *firstNamePtr)

}
