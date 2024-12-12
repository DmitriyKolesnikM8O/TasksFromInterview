package main

import (
	"fmt"
)

func main() {
	s := "aab"
	fmt.Println(partition(s))
}

func partition(s string) [][]string {
	if s == "" {
		return [][]string{}
	}

	result := [][]string{}
	current := []string{}
	lenS := len(s)

	var explore func(index int)
	explore = func(index int) {
		if index >= lenS {
			result = append(result, append([]string(nil), current...))
			return
		}

		for i := index; i < lenS; i++ {
			subStr := s[index : i+1]
			if isPalindrom(subStr) {
				current = append(current, subStr)
				explore(i + 1)
				current = current[:len(current)-1]
			}
		}
	}

	explore(0)
	return result

}

func isPalindrom(s string) bool {
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
