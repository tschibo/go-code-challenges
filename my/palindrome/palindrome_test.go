package palindrome

import "fmt"

func Example_Racecar() {
	fmt.Println(IsPalindrome("racecar"))
	fmt.Println(IsPalindromeEvo("racecar"))
	// Output:
	// true	
	// true	
}
func Example_Raceear() {
	fmt.Println(IsPalindrome("raceear"))
	fmt.Println(IsPalindromeEvo("raceear"))
	// Output:
	// false
	// false
}
func Example_Raccar() {
	fmt.Println(IsPalindrome("raccar"))
	fmt.Println(IsPalindromeEvo("raccar"))
	// Output:
	// true
	// true
}
func Example_Level() {
	fmt.Println(IsPalindrome("level"))
	fmt.Println(IsPalindromeEvo("level"))
	// Output:
	// true
	// true
}
func Example_Kayak() {
	fmt.Println(IsPalindrome("kayak"))
	fmt.Println(IsPalindromeEvo("kayak"))
	// Output:
	// true
	// true
}
func Example_Bus() {
	fmt.Println(IsPalindrome("bus"))
	fmt.Println(IsPalindromeEvo("bus"))
	// Output:
	// false
	// false
}
