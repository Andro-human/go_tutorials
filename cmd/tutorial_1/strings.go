package main

import (
	"fmt"
	"strings"
)

func myString() {
	// go uses utf-8 encoding to store strings
	// string are 
	var myString = "rÉsumÉ"
	fmt.Println(myString, len(myString))	// here the length comes 8 because É(special character) takes 2 space

	for i,v := range myString {
		println(i, v);				// it prints idx and a number (value of the byte array of the character)
	}

	fmt.Println(myString[1])		// we get partial value since special character takes 2 idx (1 and 2)
	 
	// runes solves this problem
	var myString2 = []rune("rÉsumÉ")
	var myRune = 'b'
	fmt.Printf("\nstring = %v, rune = %v ", myString2, myRune)

	var strSlice = []string{"s", "u", "p"};
	var catStr = ""
	for i:= range strSlice{
		catStr += strSlice[i]			// strings are immutable so they can't be modified (each time new string is created)	
	}
	fmt.Printf("\n%v", catStr) 

	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}

	catStr = strBuilder.String()		// much faster
	fmt.Printf("n%v", catStr)
}