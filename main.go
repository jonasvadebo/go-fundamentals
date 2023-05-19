package main

func main() {
	//Types of loops
	//All loops will be for-loops

	//1) Loop until condition is met
	//2) loop until condition with post clause
	//3) infinite loops
	//4) loop over collections.

	/** loop over collections **/
	slice := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(slice); i++ {
		println(slice[i])
	}

	for i, v := range slice {
		println(i, v)
	}

	wellKnownPorts := map[string]int{"http": 80, "https": 443}

	for _, v := range wellKnownPorts {
		println(v)
	}

	for k, v := range wellKnownPorts {
		println(k, v)
	}

}
