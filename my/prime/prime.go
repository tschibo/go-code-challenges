package prime

import "math"

func FindPrimeNumbersInInterval(start, end int) []int {
	var primes []int

	// No prime numbers until 1
	if end <= 1 {
		return primes
	}

	// Approach: https://de.wikipedia.org/wiki/Sieb_des_Eratosthenes
	//var pMap map[int]bool
	pMap := make(map[int]bool)
	for i := 2; i <= end; i++ {
		pMap[i] = true
	}
	stop := math.Sqrt(float64(end))
	stop = math.Ceil(stop)
	for num := 2; num <= end; num++ {
		if pMap[num] {
			if num >= start {
				primes = append(primes, num)
			}
			if num <= int(stop) {
				for j := num * num; j <= end; j += num {
					pMap[j] = false
				}
			}
		}
	}
	return primes
}
