package main

import "fmt"

// Using 'var' then declaring as a string then assigning it
var c string = "This is a string variable"

func main() {
	a := "Jon Ettics variables"
	b := 10
	d := true
	// Shorthand var, only inside the main func

	fmt.Printf("%T \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
}

// %T displays the TYPE of variables, eg 'string'
// %v tells it to use the default format eg. print variable as it was written
// always use , to tell the printf what to show in place of %v
