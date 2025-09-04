package main

import (
	"fmt"
)

func main() { // Output
	dash := "--------"
	fmt.Println(dash)

	var a, b string = "Hello", "World!"

	fmt.Print(a)
	fmt.Print(b)

	fmt.Println("\n")

	fmt.Print(a, "\n")
	fmt.Print(b, "\n")

	fmt.Println("\n")

	fmt.Print(a, "\n", b)

	fmt.Println("\n")

	fmt.Print(a, " ", b)

	fmt.Println("\n")

	var c, d = 10, 20
	fmt.Print(c, d)

	fmt.Println("")
	fmt.Println(dash)

	fmt.Println(a, b)

	fmt.Println(dash)

	var e string = "John"
	var f int = 15

	fmt.Printf("e has value: %v and type: %T \n", e, e)
	fmt.Printf("f has value: %v and type: %T \n", f, f)

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
