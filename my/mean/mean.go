package mean

func Mean(numbers []int) float64 {
	var sum int
	//var num int = 0
	num := len(numbers)
	for _, element := range numbers {
		sum += element
		//num += 1
	}
	mean :=  float64(sum) / float64(num)
	return mean
}