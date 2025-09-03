package main

import (
	"fmt"
)

var d int
var e int = 2
var f = 3

/*
g := 4
	":=" can only be used in a function the code above will produce an error
*/

func main() { // Variables
	dash := "--------"
	fmt.Println(dash)

	var student1 string = "John" //type is string
	var student2 = "Jane"        // type is inferred
	x := 2                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	fmt.Println(dash)

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(dash)

	var student3 string
	student3 = "Johnathon"
	fmt.Println(student3)

	fmt.Println(dash)

	d = 1
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	fmt.Println(dash)
}
