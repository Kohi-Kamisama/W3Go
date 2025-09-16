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
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	temp := 14

	if temp > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
