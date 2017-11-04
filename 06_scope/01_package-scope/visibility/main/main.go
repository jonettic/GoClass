package main

import (
	"fmt"

	"github.com/jonettic/goclass/01_package-scope/visibility/vis"
)

// Importing vis package now enables the vis package to be used
// IMPORTING IS ALWAYS AT THE FILE LEVEL, NOT PACKAGE level

func main() {
	fmt.Println(vis.MyName) // vis.MyName looks at
	vis.PrintVar()

}
