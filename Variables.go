package main

import (
	"fmt"
)

var d int
var f int = 9
var g = 7

/*
h := 1
	":=" can only be used in a function the code above will produce an error
*/

func main() { // Variables
	dash := "--------"
	fmt.Println(dash)

	var student1 string = "John" //type is string
	var student2 = "Jane"        // type is inferred
	a := 1                       //type is inferred

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(a)

	fmt.Println(dash)

	var b string
	var c int
	var d bool

	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println(dash)

	var student3 string
	student3 = "Johnathon"
	fmt.Println(student3)

	fmt.Println(dash)

	e := 3
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	fmt.Println(dash)

	var i, j, k, l int = 1, 3, 9, 7

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)

	fmt.Println(dash)

	var m, n = 1, "Hello "
	o, p := "world!", 3

	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)
	fmt.Println(p)

	fmt.Println(dash)

	var (
		q int
		r int    = 9
		s string = "Janet"
	)

	fmt.Println(q)
	fmt.Println(r)
	fmt.Println(s)

	fmt.Println(dash)
}
