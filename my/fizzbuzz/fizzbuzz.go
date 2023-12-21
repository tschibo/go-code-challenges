package fizzbuzz

import "fmt"

func FizzBuzz(end int) {
	for i := 1; i <= end; i++ {
		var s string
		if i%3 == 0 {
			s += "Fizz"
		}
		if i%5 == 0 {
			s += "Buzz"
		}
		if len(s) == 0 {
			fmt.Println(i)
		} else {
			fmt.Println(s)
		}
	}
}
