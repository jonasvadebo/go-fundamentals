package main

import "fmt"

func main() {
	//Declare variables

	var i int
	i = 42
	fmt.Println(i)

	var pi float32 = 3.14
	fmt.Println(pi)

	firstName := "Jonas"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(2, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)

	fmt.Println(r)
	fmt.Println(im)

}
