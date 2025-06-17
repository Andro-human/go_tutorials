package main

import "fmt"

type gasEngine struct{
	mpg int
	gallons int
	owner
}

func (e gasEngine) milesLeft() int {
	return e.gallons*e.mpg
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() int
}

func canMakeIt (e engine, miles int)bool {			// here defining an interface makes it more general so that can be used to call
	if (e.milesLeft() > miles) {					// other structs with the function milesLeft
		return true;
	}
	return false;
}

func myStructs() {
	var myEngine gasEngine = gasEngine{25,15, owner{"Alex"}}
	fmt.Println(myEngine.gallons, myEngine.mpg, myEngine.name, myEngine.milesLeft())			// by default is 0

	var myEngine2 = struct{								// not reusable
		mpg int
		gallons int
	}{25,15}
	fmt.Println(myEngine2.mpg, myEngine2.gallons)

	canMakeIt(myEngine, 25)
}