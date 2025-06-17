package main

import (
	"fmt"
)

func square(arr1 *[5]int) {			
	for i:= range arr1{
		arr1[i] = arr1[i]*arr1[i]
	}
}

func myPointer() {
	var p *int = new(int)		// new returns a free memory location
	fmt.Println(p)		// nil by default

	var arr1 = [5]int{0,1,2,3,4}
	square(&arr1);						// instead of making a copy, this passing the current array address
	fmt.Println(arr1)				
}