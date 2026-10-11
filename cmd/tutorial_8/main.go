package main

import (
	"fmt"
)

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)

	// Go copies the five float64 elements into the function's parameter thing2
	// Resulting in one array for result (or thing2 parameter) and another for thing1
	var result [5]float64 = square(thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1 is: %v\n", thing1)

	// Here we can see the adresses stay the same, and the result is reflected in place
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result2 [5]float64 = square2(&thing1)
	fmt.Printf("\nThe result is: %v", result2)
	fmt.Printf("\nThe value of thing1 is: %v\n", thing1)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}

// This function modifies an array in-place
func square2(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
