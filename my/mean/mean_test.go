package mean

import "fmt"

func ExampleMean1() {
	numbers := []int{1, 2, 3}
	mean := Mean(numbers)
	fmt.Println(mean)
	// Output: 2
}

func ExampleMean2() {
	numbers := []int{10, -1, 50, 23, 77}
	mean := Mean(numbers)
	fmt.Println(mean)
	// Output: 31.8
}
