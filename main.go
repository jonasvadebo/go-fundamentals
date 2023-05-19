package main

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func main() {
	//Create instances of struct
	u1 := User{
		ID:        1,
		FirstName: "Jonas",
		LastName:  "Vadebo",
	}
	u2 := User{
		ID:        2,
		FirstName: "Negin",
		LastName:  "Moradi",
	}
	if u1.ID == u2.ID {
		println("Same user")
	}

	if u1.ID != u2.ID {
		println("Not the same user!")
	}

	if u1.ID == u2.ID {
		println("Same user")
	} else {
		println("Not the same user!")
	}

	if u1.ID == u2.ID {
		println("Same user")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user!")
	} else {
		println("Diferent user!")
	}

	if u1 == u2 {
		println("Same user")
	} else if u1.FirstName == u2.FirstName {
		println("Similar user!")
	}
}
