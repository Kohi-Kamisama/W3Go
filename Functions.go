package main

import (
	"fmt"
)

var x int = 8
var y int = 4

func Functions() { // Functions
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

	fmt.Println(Add(x, y))
	fmt.Println(Sub(x, y))
	fmt.Println(Div(x, y))

	multiply := Mul(x, y)
	fmt.Println(multiply)

	fmt.Println("")
	fmt.Println(MRV(y, "Gods"))
	n1, t1 := MRV(x, "Devil")
	fmt.Println(n1, t1)

	fmt.Println("")
	_, t2 := MRV(x, "Gods")
	fmt.Println(t2)
	n3, _ := MRV(x, "Devil")
	fmt.Println(n3)

	fmt.Println(dash)

	TestCount(1)

	fmt.Println(dash)

	FactorialRecursion(4)

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

func Add(x int, y int) int {
	return x + y
}

func Sub(x int, y int) (result int) {
	result = x - y
	return
}

func Div(x int, y int) (result int) {
	result = x / y
	return result
}

func Mul(x int, y int) (result int) {
	result = x * y
	return result
}

func MRV(i int, s string) (dub int, txt string) {
	dub = i + i
	txt = s + " Realm"
	return
}

func TestCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return TestCount(x + 1)
}

func FactorialRecursion(x float64) (y float64) {
	if x > 0 {
		y = x * FactorialRecursion(x-1)
	} else {
		fmt.Println("else:", x, y)
		y = 1
	}
	return
}
