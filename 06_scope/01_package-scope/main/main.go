package main

import "fmt"

// This var means the scope of x is the WHOLE package as its NOTt inside the {}. The whole package can use this var, including other .go files inside the package
// AKA package level scope
var x int = 1001

// THIS IS PACKAGE SCOPE
func main() {
	// THIS IS BLOCK level scope
	// shows closures and how to enclose code and variables to areas
	fmt.Println(x)
	foo()
	y := 17
	// y WILL be accesible in this block because it is defined with thwe block {}
	fmt.Println(y)
}

func foo() {
	fmt.Println(x)
	//'y is not defined because the y variable is only in the 1st func block . This is called 'scope'
	// x is able to compile because its in the universal scope, before any blocks
	// func foo states what to be applied to foo() function shown in the 1st block
}
