package main

import "fmt"

func main() {
	s := "cbd"
	fmt.Println(longestPalindrome(s))

}

func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	result := string(s[0])
	for i := 0; i < len(s)-1; i++ {
		current := string(s[i])
		for j := i + 1; j < len(s); j++ {
			if checkedPalindrom(string(current) + string(s[j])) {
				if len(result) < len(string(current)+string(s[j])) {
					result = string(current) + string(s[j])
				}
			}
			current += string(s[j])
			// s
		}
	}

	return result
}

func checkedPalindrom(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}

	return true
}
