package main

import "fmt"

func main() {
	// Fixed length
	var intArr [3]int32 = [3]int32{1, 2, 3}

	// Indexable
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])
	fmt.Println(intArr)

	// Contiguous in memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Printf("The length is %d with capacity %d", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("\nThe length is %d with capacity %d\n", len(intSlice), cap(intSlice))

	// Slices can share data, can be created with specific len, cap
	// To append more than one value use:
	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	var intSlice3 []int32 = make([]int32, 3, 8)
	// intSlice3 := make([]int32, 3, 8)
	fmt.Println(intSlice3)

	// {"":#, ...}
	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap)

	myMap2 := map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(myMap2["Adam"])
	var age, ok = myMap2["Jason"]
	fmt.Println(age, ok)
	delete(myMap2, "Adam")
	fmt.Println(myMap2["Adam"])

	myMap2["Christian"] = 37
	myMap2["Samir"] = 12
	for name, age := range myMap2 {
		fmt.Printf("Name: %v\n, Age: %v \n ", name, age)
	}

	for i, v := range intArr {
		fmt.Printf("Index: %v, Value: %v \n", i, v)
	}

	// While keyword doesn't exist but a for loop with a condition can be used
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
