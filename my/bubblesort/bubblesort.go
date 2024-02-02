package bubblesort

func Sort(array []int) []int {
	sorted := make([]int, len(array))
	copy(sorted, array)

	for i := len(sorted) - 1; i >= 1; i-- {
		for j := 0; j < i; j++ {
			if sorted[j] > sorted[j+1] {
				temp := sorted[j]
				sorted[j] = sorted[j+1]
				sorted[j+1] = temp
			}
		}
	}
	return sorted
}
