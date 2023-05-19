package main

func main() {
	//Types of loops
	//All loops will be for-loops

	//1) Loop until condition is met
	//2) loop until condition with post clause
	//3) infinite loops
	//4) loop over collections.

	/** Loop until condution **/
	var i int
	for i < 5 {
		println(i)
		i++

		if i == 3 {
			//break
			continue
		}
		println("Continuing....")
	}
}
