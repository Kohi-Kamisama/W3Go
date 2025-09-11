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

	// Sighned Integers - can store both positive and negetive values
	// Unsighned Integers - can only store positive values

	var i1 int = 420
	var i2 int = -420

	fmt.Printf("Type: %T, value: %v\n", i1, i1)
	fmt.Printf("Type: %T, value: %v\n", i2, i2)

	fmt.Println(dash)

	fmt.Println("int8: 8 bits, -128 to 127")
	fmt.Println("int16: 16 bits, -32768 to 32767")
	fmt.Println("int32: 32 bits, -2147483648 to 147483647")
	fmt.Println("int64: 64 bits, -9223372036854775808 to 9223372036854775807")

	fmt.Println("\nint will default to eather to int32 or int64 depending on your platform system")

	fmt.Println(dash)

	var u1 uint = 420

	fmt.Printf("Type: %T, value: %v\n", u1, u1)

	fmt.Println(dash)

	//fmt.Println("int8: 8 bits, -128 to 127")

	fmt.Println("unit8: 8 bits, 0 to 255")
	fmt.Println("unit16: 16 bits , 0 to 65535")
	fmt.Println("unit32: 32 bits , 0 to 4294967295")
	fmt.Println("unit64: 64 bits , 0 to 18446744073709551615")

	fmt.Println("\nunit will default to eather to unit32 or unit64 depending on your platform system")

	fmt.Println(dash)

	fmt.Println("float32: 32 bits, -3.4e+38 to 3.4e+38")
	fmt.Println("float64: 64 bits, -1.7e+308 to 1.7e+308")

	uf := 1.7e308

	fmt.Println("\nDefault for a unspecify float is float64")
	fmt.Printf("Type: %T, value: %v\n", uf, uf)

	fmt.Println(dash)

	var s1 string = "HELP"
	var s2 string
	s3 := "PIE"

	fmt.Printf("Type: %T, value: %v", s1, s1)
	fmt.Printf("\nType: %T, value: %v", s2, s2)
	fmt.Printf("\nType: %T, value: %v\n", s3, s3)

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	// fmt.Println()

}
