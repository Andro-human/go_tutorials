package main

import "fmt"

func myArr() {
	var intArr[3] int32
	intArr[1] = 123
	fmt.Println(intArr[1]);
	fmt.Println(intArr[1:3])		// accessing element at index 1 and 2

	fmt.Println(&intArr[0])			// prints out the memory location of 0th index element

	intArr2 := [...]int32{1,2,3}
	fmt.Println(intArr2)

	// slices are arrays with additional functionalities (similar to vectors in c++)
	var intSlice[] int32 = []int32{4,5,6}		// by omitting the length we have a slice
	
	intSlice = append(intSlice, 7)			// similar to push_back in c++
	
	fmt.Printf("the length is %v with capacity %v", len(intSlice), cap(intSlice))			

	intSlice2 := []int32{8,9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)


	for i,v := range intArr {
		fmt.Printf("indev: %v, Value: %v \n", i, v)
	}

	var testSlice[] int = make([]int, 0, 3)
	fmt.Println(testSlice)
} 