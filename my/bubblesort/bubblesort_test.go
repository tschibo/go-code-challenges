package bubblesort

import (
	"fmt"
	"testing"
)

func Test_Sort1(t *testing.T) {
	input := []int{1, 2, 3}
	expect := []int{1, 2, 3}

	got := Sort(input)

	if fmt.Sprint(got) != fmt.Sprint(expect) {
		t.Errorf("Expected %s, got %s", fmt.Sprint(expect), fmt.Sprint(got))
	}
}

func Test_Sort2(t *testing.T) {
	input := []int{3, 2, 1}
	expect := []int{1, 2, 3}

	got := Sort(input)

	if fmt.Sprint(got) != fmt.Sprint(expect) {
		t.Errorf("Expected %s, got %s", fmt.Sprint(expect), fmt.Sprint(got))
	}
}

func Test_Sort3(t *testing.T) {
	input := []int{6, 20, 8, 19, 56, 23, 87, 41, 49, 53}
	expect := []int{6, 8, 19, 20, 23, 41, 49, 53, 56, 87}

	got := Sort(input)

	if fmt.Sprint(got) != fmt.Sprint(expect) {
		t.Errorf("Expected %s, got %s", fmt.Sprint(expect), fmt.Sprint(got))
	}
}
