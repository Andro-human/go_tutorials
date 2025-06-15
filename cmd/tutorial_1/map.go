package main

import "fmt"

func myMap() {
	var myMap map[string]int = make(map[string]int)
	fmt.Println(myMap)

	var myMap2 = map[string]int{"Adam": 23, "Sarah":21}
	fmt.Println(myMap2)

	var age, ok = myMap2["Jason"]
	if (ok) {
		fmt.Println(age)
	} else {
		fmt.Println("not found")
	}

	delete(myMap2, "Adam")			// deletes a specific key in map

	for name, age := range myMap2 {
		fmt.Printf("Name: %v, Age:%v \n", name, age);
	}

	
}