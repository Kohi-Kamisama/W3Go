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

	fmt.Print("\n\n")

	fmt.Print(a, "\n")
	fmt.Print(b, "\n")

	fmt.Print("\n\n")

	fmt.Print(a, "\n", b)

	fmt.Print("\n\n")

	fmt.Print(a, " ", b)

	fmt.Print("\n\n")

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

	var flt = 3.14
	var txt = "Pi"

	fmt.Printf("%v \n", flt)   // Prints in default formate
	fmt.Printf("%#v \n", flt)  // Prints in Go-syntax format
	fmt.Printf("%v%% \n", flt) // "%%" Prints the % sign
	fmt.Printf("%T", flt)      // Prints the type

	fmt.Print("\n\n")

	fmt.Printf("%v \n", txt)
	fmt.Printf("%#v \n", txt)
	fmt.Printf("%T \n", txt)

	fmt.Println(dash)

	fmt.Printf("%b \n", f)   // Base 2
	fmt.Printf("%d \n", f)   // Base 10
	fmt.Printf("%+d \n", f)  // Base 10 and always shows sign
	fmt.Printf("%o \n", f)   // Base 8
	fmt.Printf("%O \n", f)   // Base 8, with leading 0o. Wont highlit %O simuler to the other verbs
	fmt.Printf("%x \n", f)   // Base 16, lowercase
	fmt.Printf("%X \n", f)   // Base 16, uppercase
	fmt.Printf("%#X \n", f)  // Base 16, with leading 0x
	fmt.Printf("%4d \n", f)  // Pad with spaces, width 4 right justified
	fmt.Printf("%-4d \n", f) // Pad with spaces, width 4 left justified
	fmt.Printf("%04d \n", f) // Pad with zeroes, width 4

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
