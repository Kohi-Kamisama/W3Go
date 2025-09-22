package main

import (
	"fmt"
)

func Loops() { // Loops
	dash := "--------"
	fmt.Println(dash)

	for a := 0; a < 5; a++ {
		fmt.Println(a)
	}

	fmt.Println(dash)

	for b := 0; b < 5; b++ {
		if b == 3 {
			fmt.Println("continue")
			continue
		}
		fmt.Println(b)
	}

	fmt.Println(dash)

	for c := 0; c < 5; c++ {
		if c == 3 {
			fmt.Println("break")
			break
		}
		fmt.Println(c)
	}

	fmt.Println(dash)

	adj := [2]string{"Big", "Tasty"}
	fru := [3]string{"apple", "orange", "peach"}

	for d := 0; d < len(adj); d++ {
		for e := 0; e < len(fru); e++ {
			fmt.Println(adj[d], fru[e])
		}
	}

	fmt.Println(dash)

	for idx, val := range fru {
		fmt.Printf("%v: %v\n", idx, val)
	}

	fmt.Println("")

	for _, val := range fru {
		fmt.Printf("%v\n", val)
	}

	fmt.Println("")

	for idx := range fru {
		fmt.Printf("%v\n", idx)
	}

	fmt.Println(dash)

	// Use to copy and past, from f to the next line. Comment out when following the examples.
	//fmt.Println()

}
