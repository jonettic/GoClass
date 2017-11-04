package vis

import "fmt"

func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourname)

	// This is a function and can be called into other packages
	// yourname variable IS NOW ABLE be exported through PrintVar function!! Because PrintVar is capitalised
}
