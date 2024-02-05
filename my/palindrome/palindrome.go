package palindrome

func IsPalindrome(word string) bool {
	match := true
	letters := len(word)
	for i, j := 0, letters-1; i < letters/2; i, j = i+1, j-1 {
		if word[i] != word[j] {
			match = false
		}
	}
	return match
}

func IsPalindromeEvo(word string) bool {
	return word == reverse(word)
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
