package main

import (
	"fmt"
)

func Slices() { // Slices
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

	// Create slice from array
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

	// Capacity omitted (excluded)
	s5 := make([]int, 5)
	fmt.Printf("Slice 5: %v\n", s5)
	fmt.Printf("Length: %v\n", len(s5))
	fmt.Printf("Capacity: %v\n", cap(s5))

	fmt.Println(dash)

	// Access elements
	juu := []int{10, 20, 30}
	fmt.Printf("juu 0: %v\n", juu[0])
	fmt.Printf("juu 2: %v\n", juu[2])

	// Change elements
	juu[2] = 50
	fmt.Printf("\nUpdated juu 2: %v\n", juu[2])

	fmt.Println(dash)

	s6 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("Slice 6: %v\n", s6)
	fmt.Printf("Length: %v\n", len(s6))
	fmt.Printf("Capacity: %v\n", cap(s6))

	fmt.Println("")

	// Append elements
	s6 = append(s6, 7, 8)
	fmt.Printf("Slice 6: %v\n", s6)
	fmt.Printf("Length: %v\n", len(s6))
	fmt.Printf("Capacity: %v\n", cap(s6))

	fmt.Println(dash)

	// Append two slices together
	s7 := []int{2, 4, 6, 8}
	s8 := []int{10, 12, 14, 16}
	s9 := append(s7, s8...)

	fmt.Printf("Slice 9: %v\n", s9)
	fmt.Printf("Length: %v\n", len(s9))
	fmt.Printf("Capacity: %v\n", cap(s9))

	fmt.Println(dash)

	// Changing the length
	arr2 := [6]int{0, 1, 2, 3, 4, 5}

	fmt.Println("Slice array")
	s10 := arr2[1:5]

	fmt.Printf("Slice 10: %v\n", s10)
	fmt.Printf("Length: %v\n", len(s10))
	fmt.Printf("Capacity: %v\n", cap(s10))

	fmt.Println("\nChange lenght be re-slicing the array")
	s10 = arr2[1:3]

	fmt.Printf("Slice 10: %v\n", s10)
	fmt.Printf("Length: %v\n", len(s10))
	fmt.Printf("Capacity: %v\n", cap(s10))

	fmt.Println("\nChange lenght by appending items")
	s10 = append(s10, 3, 4, 5, 6)

	fmt.Printf("Slice 10: %v\n", s10)
	fmt.Printf("Length: %v\n", len(s10))
	fmt.Printf("Capacity: %v\n", cap(s10))

	fmt.Println(dash)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println("Original slice")
	fmt.Printf("numbers: %v\n", numbers)
	fmt.Printf("Length: %v\n", len(numbers))
	fmt.Printf("Capacity: %v\n", cap(numbers))

	// Memory efficiency with copy()
	fmt.Print("\nCreate copy with only needed numbers\n\n")

	needednumbers := numbers[:len(numbers)-10] // numbers[:5]

	fmt.Println("Needed numbers:", needednumbers)
	fmt.Println("length:", len(needednumbers))
	fmt.Println("Capacity:", cap(needednumbers))
	fmt.Print("\n")

	numberscopy := make([]int, len(needednumbers))

	fmt.Println("numberscopy before:", numberscopy)
	fmt.Println("length:", len(numberscopy))
	fmt.Println("Capacity:", cap(numberscopy))
	fmt.Print("\n")

	copy(numberscopy, needednumbers)

	fmt.Printf("numberscopy after: %v\n", numberscopy)
	fmt.Printf("Length: %v\n", len(numberscopy))
	fmt.Printf("Capacity: %v\n", cap(numberscopy))

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
