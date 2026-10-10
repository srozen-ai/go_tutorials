package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
	// owner // would add name: ""
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	ownerInfo := owner{"Carlos"}
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 15, ownerInfo: ownerInfo}

	var myEngine2 = struct {
		mpg     uint8
		gallons uint8
	}{25, 15} // Not reusable

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.ownerInfo.name, myEngine.int)
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	fmt.Printf("Total miles left in tank: %v\n", myEngine.milesLeft())

	var eEngine electricEngine = electricEngine{mpkwh: 30, kwh: 40}

	miles := uint8(255)

	fmt.Println("My engine: ")
	canMakeIt(myEngine, miles)

	fmt.Println("Electric engine: ")
	canMakeIt(eEngine, miles)
}
