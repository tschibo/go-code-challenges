package fibonacci

//var fMap map[int]int

func Fibonacci(n int) int {
	//_, ok := fMap[n]
	// if ok {
	// 	return fMap[n]
	// }
	if n <= 1 {
		//fMap[n] = 0
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
