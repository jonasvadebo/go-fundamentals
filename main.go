package main

import "fmt"

func main() {
	// Structs

	type user struct {
		ID        int
		firstName string
		lastName  string
	}

	//Declare a used and initilize with its zero values
	var u user
	fmt.Println(u)

	//ser values

	u.ID = 10
	u.firstName = "Jonas"
	u.lastName = "Vadebo"

	fmt.Println(u)

}
