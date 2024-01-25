package stringcalc

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	if numbers == "" {
		return 0
	} else {
		split := strings.Split(numbers, ",")
		var sum int
		for _, num := range split {
			i, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			sum += i
		}
		return sum
	}
}
