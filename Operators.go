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

	fmt.Println("")
	fmt.Println("Multiplication: 8*4 =", x*y)

	fmt.Println("")
	fmt.Println("Division: 8/4 =", x/y)
	fmt.Println("Modulus: 8%4 =", x%y)

	fmt.Println("")

	x++
	fmt.Println("Increment: 8++ =", x)

	y--
	fmt.Println("Decrement: 4-- =", y)

	fmt.Println(dash)

	var a = 8
	fmt.Println(a)

	fmt.Println("")
	var b = 8
	b += 4
	fmt.Println(b)

	fmt.Println("")
	var c = 8
	c -= 4
	fmt.Println(c)

	fmt.Println("")
	var d = 8
	d *= 4
	fmt.Println(d)

	fmt.Println("")
	var e = 8
	e /= 4
	fmt.Println(e)

	fmt.Println("")
	var f = 8
	f %= 4
	fmt.Println(f)

	fmt.Println("")
	var g = 8
	fmt.Printf("g is %b \n", g)
	fmt.Printf("4 is %03b \n", 4)
	g &= 4
	fmt.Printf("g now is %03b \n", g)

	fmt.Println("")
	var h = 8
	fmt.Printf("h is %b \n", h)
	fmt.Printf("4 is %03b \n", 4)
	h |= 4
	fmt.Printf("h now is %03b \n", h)

	fmt.Println("")
	var i = 8
	fmt.Printf("i is %b \n", i)
	fmt.Printf("4 is %03b \n", 4)
	i ^= 4
	fmt.Printf("i now is %03b \n", i)

	fmt.Println("")
	var j = 8
	fmt.Printf("j is %b \n", j)
	j >>= 4
	fmt.Printf("j now is %03b \n", j)

	fmt.Println("")
	var k = 8
	fmt.Printf("k is %b \n", k)
	k <<= 4
	fmt.Printf("k now is %03b \n", k)

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
