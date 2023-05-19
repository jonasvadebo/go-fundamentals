package main

func main() {
	//Types of loops
	//All loops will be for-loops

	//1) Loop until condition is met
	//2) loop until condition with post clause
	//3) infinite loops
	//4) loop over collections.

	/** infinite loops, ugly way **/
	var i int
	//for ; ;{ ugly way
	for {
		if i == 5 {
			break
		}
		println(i)
		i++
	}
}
