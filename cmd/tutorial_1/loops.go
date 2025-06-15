package main

import "fmt"

func myLoops () {
	//whlie loop
	var i int = 0
	for i<10{
		fmt.Println(i)
		i++
	}

	for {
		if (i >= 10) {
			break
		}
		fmt.Println(i)
		i++
	}

	//for
	for i:=0; i<10; i++ {
		fmt.Println(i)
	}
	
}