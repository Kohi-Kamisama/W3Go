package main

import (
	"fmt"
)

func main() { // Operators
	dash := "--------"
	fmt.Println(dash)

	x := 8
	y := 4

	fmt.Print("Arithmetic Operators\n\n")

	fmt.Println("Addition: 8+4 =", x+y)
	fmt.Println("Subtraction: 8-4 =", x-y)
	fmt.Println("Multiplication: 8*4 =", x*y)
	fmt.Println("Division: 8/4 =", x/y)
	fmt.Println("Modulus: 8%4 =", x%y)

	x++
	fmt.Println("Increment: 8++ =", x)

	y--
	fmt.Println("Decrement: 4-- =", y)

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
