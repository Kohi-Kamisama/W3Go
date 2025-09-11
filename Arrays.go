package main

import (
	"fmt"
)

func Arrays() { // Arrays
	dash := "--------"
	fmt.Println(dash)

	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(dash)

	var arr3 = [...]int{9, 10, 11}
	arr4 := [...]int{12, 13, 14, 15, 16}

	fmt.Println(arr3)
	fmt.Println(arr4)

	fmt.Println(dash)

	var pies = [4]string{"Apple", "Blueberry", "Strawberry", "Pumkin"}
	fmt.Print(pies, "\n\n")

	// Access elments of an arry
	fmt.Println(pies[1])
	fmt.Println(pies[2])

	fmt.Println("")

	// Change elements of an arry
	pies[3] = "Pumpkin"
	fmt.Println(pies)

	fmt.Println(dash)

	arr5 := [5]int{}              // Not initialized
	arr6 := [5]int{1, 2}          // Partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} // Fully initialized
	arr8 := [5]int{1: 2, 3: 4}    //Specificly initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)
	fmt.Println(arr8)

	fmt.Println(dash)

	// Find the length
	fmt.Println(len(pies))
	fmt.Println(len(arr4))

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
