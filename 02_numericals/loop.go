package main

import "fmt"

func main() {
	// fmt.Printf("%d - %b | %x \n testing \n \t this is a test", 42, 42, 42)
	// fmt.Printf("%d - %b | %#x \n", 42, 42, 42)
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %#X \n", i, i, i)
	}
}
