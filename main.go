package main

import "fmt"

func main() {
	//Array
	/* 	var users [3]string
	   	users[0] = "Jonas"
	   	users[1] = "Niklas"
	   	users[2] = "Fredrik" */

	users := [3]string{"Jonas", "Niklas", "Fredrik"}

	fmt.Println(users)
	fmt.Println(users[1])

}
