package main

import (
	"errors"
	"fmt"
)

func main() {
	printValue := "Hello World!"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 0
	var result, remainder, err = intDivision(numerator, denominator)
	/* if err != nil {
		fmt.Printf("%s\n", err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division is %v\n", result)
	} else {
		fmt.Printf("The result of the integer division is %v with remainder %v\n", result, remainder)
	}
	*/

	switch {
	case err != nil:
		fmt.Println(err)
	case remainder == 0:
		fmt.Printf("The result of the integer division is %d\n", result)
	default:
		fmt.Printf("The result of the integer division is %d with remainder %d\n", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("Remainder is 0")
	case 1, 2:
		fmt.Printf("The division was close")
	default:
		fmt.Printf("The division was not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
	// Zero division returns panic: runtime error: integer divide by zero
}
