package main

import (
	"fmt"
)

func main() { // Functions
	dash := "--------"
	fmt.Println(dash)

	Message()
	fmt.Println("")
	Message()

	fmt.Println(dash)

	// This part I'm modifying the example to use what I learned about arrays and loops

	names := [4]string{"Micheal", "Jack", "Luna", "Ana"}
	for n := range names {
		FamilyName(names[n])
	}

	fmt.Println("")

	age := [4]int{42, 24, 23, 1}
	for a := range age {
		FamilyAge(names[a], age[a])
	}

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}

func Message() {
	fmt.Println("Gods Die")
}

func FamilyName(name string) {
	fmt.Println("Welcome", name, "Hana")
}

func FamilyAge(name string, age int) {
	fmt.Println(name, "Hana is", age, "Years Old")
}
