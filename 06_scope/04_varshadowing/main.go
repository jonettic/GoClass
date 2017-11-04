package main

import "fmt"

func max(x int) int {
	//This func max is actually at package level. It can be called through the package. IT IS NOT EXPORTED as it is not a capital letter
	return 42 + x
}

// You can then call function max from anywhere

func main() {
	max := max(7)
	// This is declaring a variable at BLOCK level but it is also calling the max function form package level
	// this is bad coding and is called variable shadowing.
	fmt.Println(max)
	// you cannot call the variable as max variable is named the same as the func block above
}
