package main

import "fmt"

func main() {
	// fmt.Printf("%d - %b | %x \n testing \n \t this is a test", 42, 42, 42)
	// fmt.Printf("%d - %b | %#x \n", 42, 42, 42)
	for i := 50; i < 150; i++ {
		fmt.Printf("%d \t %b \t %#X \t %q \n", i, i, i, i)
		//%q = print a letter in UTF8
	}
}
