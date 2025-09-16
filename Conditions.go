package main

import (
	"fmt"
)

func main() { // Conditions
	dash := "--------"
	fmt.Println(dash)

	if 20 > 18 {
		fmt.Println("20 is grater than 18 ")
	}

	a := 20
	b := 18

	if a > b {
		fmt.Println("a is greater than b")
	}

	fmt.Println(dash)

	time := 20

	if time < 18 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}

	temp := 14

	if temp > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}

	fmt.Println(dash)

	time = 22

	if time < 10 {
		fmt.Println("Good morning")
	} else if time < 20 {
		fmt.Println("Good day")
	} else {
		fmt.Println("Good evening")
	}

	c := 14
	d := 14

	if c < d {
		fmt.Println("c is less than d.")
	} else if c > d {
		fmt.Println("c is more than d")
	} else {
		fmt.Println("c and d are equal")
	}

	e := 30

	if e >= 10 {
		fmt.Println("e is larger or equal to 10")
	} else if e > 20 {
		fmt.Println("e is larger than 20")
	} else {
		fmt.Println("e is less than 10")
	}

	fmt.Println(dash)

	f := 20

	if f >= 10 {
		fmt.Println("f is more than 10")
		if f > 15 {
			fmt.Println("f is also more than 15")
		}
	} else {
		fmt.Println("f is less than 10")
	}

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
