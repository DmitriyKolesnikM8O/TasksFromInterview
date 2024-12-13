package main

import "fmt"

func main() {
	s := "cbd"
	fmt.Println(longestPalindrome(s))

}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	result := ""
	current := ""
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			current = s[i : j+1]
			if checkedPalindrom(string(current)) {
				if len(result) < len(string(current)) {
					result = string(current)
				}
			}
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
