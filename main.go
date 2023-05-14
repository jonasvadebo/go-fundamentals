package main

import "fmt"

func main() {
	m := map[string]int{"Jonas": 1980, "Johan": 1981}

	fmt.Println(m)

	fmt.Println(m["Jonas"])

	m["Johan"] = 1982

	fmt.Println(m)

	fmt.Println("Full map!")

	for k, v := range m {
		fmt.Printf("%s -> %d \n", k, v)

	}

	delete(m, "Johan")

	fmt.Println("Reduced map!")

	for k, v := range m {
		fmt.Printf("%s -> %d \n", k, v)

	}

}
