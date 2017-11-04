package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42
	// this wont compile
	// you cannot declare and assign a var AFTER calling the var
	// x = 42 MUST be stated before the Println (or calling the var)
}

var y = 42

// var y is outside of a block and can be called anywhere in the packge
// It is called the OUTER SCOPE
