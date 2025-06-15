package main

import "fmt"

func printMe(printValue string) string {
	fmt.Println("hello world");
	return printValue
}

func mult (num1 int, str1 string) (int, string ) {		// we can specify that this function is going to return multiple values
	return num1, str1
}