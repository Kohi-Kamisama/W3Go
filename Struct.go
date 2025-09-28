package main

import (
	"fmt"
)

// Declaring a new type with struct
type Person struct {
	name   string
	age    int
	job    string
	salary int
}

func Struct() { // Struct
	dash := "--------"
	fmt.Println(dash)

	// Creating varibals with the type we created
	var pers1 Person
	var pers2 Person

	// Specifing the varibauls struct members
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Markiting"
	pers2.salary = 4500

	// Accessing and printing the values
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)
	fmt.Println("")
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	fmt.Println(dash)

	// Calling the function to print each varibal of Person
	PrintPerson(pers1)
	fmt.Println("")
	PrintPerson(pers2)

	fmt.Println(dash)
	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}

// Passing stuct as a function agrgument
func PrintPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}
