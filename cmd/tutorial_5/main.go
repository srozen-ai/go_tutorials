package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString = "résumé"
	var indexed = myString[0]
	fmt.Printf("%c, %v, %T\n", indexed, indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// Thinks length is 8 as 'é' occupies 2 bytes
	fmt.Printf("The length of 'myString' is %v\n", len(myString))

	var myString2 = []rune("résumé")
	var indexed2 = myString2[0]
	fmt.Printf("%c, %v, %T\n", indexed2, indexed2, indexed2)
	for i, v := range myString2 {
		fmt.Println(i, v)
	}

	// Understands length is 6, identifies 'é' as one character
	fmt.Printf("The length of 'myString2' is %v\n", len(myString2))

	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "t", "r", "i", "n", "g"}
	var strBuilder strings.Builder
	/*
		var catStr = ""
		for i := range strSlice {
			catStr += strSlice[i] // This would create a new string every time
		}
	*/
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i]) // Uses an array (slice) to allocate values by appending them
	}
	var catStr = strBuilder.String() // Only at the end is the string created with the appended values (much faster)
	fmt.Printf("\n%v", catStr)
	// catStr[0] = 'x' // Doesn't work since you can not modify strings once created
}
