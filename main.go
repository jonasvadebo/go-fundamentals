package main

import "fmt"

func main() {
	/* 	users := [3]string{"Jonas", "Niklas", "Fredrik"}

	   	userSlice := users[:]

	   	users[0] = "Tomas"
	   	userSlice[1] = "Steven"

	   	fmt.Println(users, userSlice) */

	//Slice
	users := []string{"Jonas", "Niklas", "Fredrik"}

	fmt.Println(users)
	fmt.Printf("The length of slice:  %d and the capacity of slice: %d \n", len(users), cap(users))

	users = append(users, "Jonny")

	fmt.Printf("The length of slice:  %d and the capacity of slice: %d \n", len(users), cap(users))

	user2 := users[1:]
	user3 := users[:3]
	user4 := users[1:2]
	fmt.Println(user2, user3, user4)

}
