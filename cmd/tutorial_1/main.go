package main 
import "fmt"


func main() {
	var intNum int			// statically typed language
	fmt.Println(intNum)				// we need to use each variable we define

	var floatNum float64 = 12345678.9		// float 64 is more precise then float 32
	fmt.Printf("hello %v", floatNum)		// printf can be used to insert variables

	var result float64 = floatNum + float64(intNum)		// cannot perform operations with mixed types
	fmt.Println(result)

	var myRune rune = 'a'			// rune represents characters
	fmt.Println(myRune)

	myVar := "text"			// the datatype is infered this way
	fmt.Println(myVar)

	var1, var2 := 1,2						
	fmt.Println(var1, var2)

	var myFunc string = printMe("return this")		// defining type when it isn't obvious is a good practice
	fmt.Println(myFunc);

	res1, res2 := mult(9, "hii")
	fmt.Println(res1, res2);			// default value is 0 if not assigned

	myArr()
	myMap()
	myLoops()
	myString()
}
