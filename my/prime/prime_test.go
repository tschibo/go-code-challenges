package prime

import "fmt"

func ExampleFindPrimeNumbersInInterval1() {
	fmt.Println(FindPrimeNumbersInInterval(1, 1))
	// Output: []
}
func ExampleFindPrimeNumbersInInterval2() {
	fmt.Println(FindPrimeNumbersInInterval(0, 100))
	// Output: [2 3 5 7 11 13 17 19 23 29 31 37 41 43 47 53 59 61 67 71 73 79 83 89 97]
}
func ExampleFindPrimeNumbersInInterval3() {
	fmt.Println(FindPrimeNumbersInInterval(10, 20))
	// Output: [11 13 17 19]
}
