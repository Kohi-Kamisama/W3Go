package main

import (
	"fmt"
)

func main() { // DataTypes
	dash := "--------"
	fmt.Println(dash)

	var a bool = true     // Boolean
	var b int = 5         // Integer
	var c float32 = 3.14  // Floating point number
	var d string = "HELP" // String

	fmt.Println("Bolean:  ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("float:   ", c)
	fmt.Println("String:  ", d)

	fmt.Println(dash)

	var b1 bool = true // typed declartion with initial value
	var b2 = true      // untyped declartion with initial value
	var b3 bool        // typed declartion without initial value
	b4 := true         // untyped declartion with initial value

	fmt.Println("b1: ", b1, " : typed declartion with initial value")
	fmt.Println("b2: ", b2, " : untyped declartion with initial value")
	fmt.Println("b3: ", b3, ": typed declartion without initial value")
	fmt.Println("b4: ", b4, " : untyped declartion with initial value")

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	// fmt.Println()

}
