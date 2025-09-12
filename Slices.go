package main

import (
	"fmt"
)

func main() { // Slices
	dash := "--------"
	fmt.Println(dash)

	s1 := []int{}
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	fmt.Println(s1)

	fmt.Println("")

	s2 := []string{"Evrything", "eventualy", "dies", "including", "GODS!"}
	fmt.Println(len(s2))
	fmt.Println(cap(s2))
	fmt.Println(s2)

	fmt.Println(dash)

	arr1 := [8]int{2, 4, 6, 8, 10, 12, 14, 16}
	s3 := arr1[3:6]

	fmt.Printf("Slice 3: %v\n", s3)
	fmt.Printf("Length: %v\n", len(s3))
	fmt.Printf("Capacity: %v\n", cap(s3))

	// The capacity starts from arr1[3] and ends at the end of the array not at arr1[6]

	fmt.Println(dash)

	// Creating slices with make()

	s4 := make([]int, 5, 10)
	fmt.Printf("Slice 4: %v\n", s4)
	fmt.Printf("Length: %v\n", len(s4))
	fmt.Printf("Capacity: %v\n", cap(s4))

	fmt.Println("")

	// Capacity omitted
	s5 := make([]int, 5)
	fmt.Printf("Slice 5: %v\n", s5)
	fmt.Printf("Length: %v\n", len(s5))
	fmt.Printf("Capacity: %v\n", cap(s5))

	fmt.Println(dash)
	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
