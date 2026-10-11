package main

import "fmt"

func main() {
	var p *int32
	var i int32

	var ptr *int32 = new(int32) // Allocates an int32 initialized to 0 and assigns its address to ptr

	// When printing, ptr is the actual pointer, *ptr is the value it points to
	fmt.Println(p, i, ptr, *ptr)

	// Pointers have a memory address themselves, they store another memory address
	// *: dereferencing the pointer
	*ptr = 10
	fmt.Println("Address ptr points to now takes value of:", *ptr)

	// &: getting a variables address
	var ptrI = &i
	fmt.Println("i is at address:", ptrI)

	*ptrI = 1
	fmt.Println("The value of i (the value to which ptrI points to) is now:", i)

	var k int32 = 2
	i = k
	fmt.Printf("Value of k copied into i. i:%v, k:%v, *ptrI:%v\n", i, k, *ptrI)

	// Since slices point to an underlying array, when copying a slice,
	// changes to one of these apply to both
	var slice = []int32{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Println(slice)
	fmt.Println(sliceCopy)

	// To avoid this, you would:
	var sliceCopy2 []int32 = make([]int32, len(slice))
	copy(sliceCopy2, slice)
	sliceCopy2[0] = 99
	fmt.Println(slice)
	fmt.Println(sliceCopy2)
	// Here, make(..) allocates a slice with its own underlying array
	// copy() copies the elements into that array

	// A shorter alternative:
	sliceCopy3 := append([]int32(nil), sliceCopy...)
	sliceCopy3[0] = 99
	fmt.Println(slice)
	fmt.Println(sliceCopy3)
}
