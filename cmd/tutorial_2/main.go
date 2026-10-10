package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 9223372036854775807
	intNum = intNum + 1
	fmt.Println(intNum)

	var floatNum float32 = 12345678.9
	fmt.Println(floatNum)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2)
	fmt.Println(intNum1 % intNum2)

	fmt.Println("len() count for γ:", len("γ"))
	fmt.Println("RuneCountInString() for γ:", utf8.RuneCountInString("γ"))

	var myBoolean bool = false
	fmt.Println(myBoolean)

	intNum3 := 3
	fmt.Println(intNum3)
}
