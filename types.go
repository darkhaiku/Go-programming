package main

import (
	"fmt"
)

func main() {

	// integer types
	var integer int
	var floatingPoint float64
	var complexNumber complex64
	fmt.Println(integer, floatingPoint, complexNumber)

	// boolean type
	var boolean bool
	fmt.Println(boolean)

	// Alphabetic 
	var (
	character = 'A'
	utf8Character = 'س' 
	sentence = "hello world"
	)
	fmt.Println(string(character), string(utf8Character), sentence)
	fmt.Println(string(sentence[0]), string(sentence[1]))
}