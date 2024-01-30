package happynumber

func getNextNumber(digits []int) int {
	var sum int
	for _, digit := range digits {
		sum += digit * digit
	}
	return sum
}

func IsHappy(number int) bool {

	next := number
	//fmt.Println(next)

	for next != 1 {
		var digits []int
		var digit int

		for next > 0 {
			digit = next % 10
			digits = append(digits, digit)
			next = next / 10
		}
		next = getNextNumber(digits)
		//fmt.Println(next)

		if next == number {
			// periodic
			return false
		}

	}

	return true
}
