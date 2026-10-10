package main

import "fmt"

func main() {
	printValue := "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder int = intDivision(numerator, denominator)
	fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)

}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
	// Zero division returns panic: runtime error: integer divide by zero
}
