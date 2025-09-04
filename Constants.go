package main

import (
	"fmt"
)

func main() { // Constants
	dash := "--------"
	fmt.Println(dash)

	const PI = 3.14159265
	fmt.Println(PI)
	// Constants usually typed in uppercase for easy identification, also differntation from variables. Varibales use Pascal case

	fmt.Println(dash)

	const A int = 1
	fmt.Println(A)

	fmt.Println(dash)

	/*
		const B = 3
		B = 4
		fmt.Println(B)

			once you declare a constant you cannot chang it.
			The reult would be a message: cannot assighn to B
	*/

	const (
		C int = 9
		D     = 3.1415
		E     = "Hello World!"
	)

	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)

	fmt.Println(dash)
	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
