package stringcalc

import "fmt"

func ExampleAdd() {
	fmt.Println(Add(""))
	fmt.Println(Add("1,2,3,4,5"))
	// Output:
	// 0
	// 15
}
