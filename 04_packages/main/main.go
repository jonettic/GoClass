package main

import (
	"fmt"

	"github.com/GoesToEleven/GolangTraining/02_package/stringutil"
	//someAlias "github.com/GoesToEleven/GolangTraining/02_package/icomefromalaska"
)

func main() {
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	fmt.Println(stringutil.MyName)
	//Shows you can call an exported func package which contains an unxported func package
	//Capital letter allows the function to be exported into other packages, such as the main func package
}
