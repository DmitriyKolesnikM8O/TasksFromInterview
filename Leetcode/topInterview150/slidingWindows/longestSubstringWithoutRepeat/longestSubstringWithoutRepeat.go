package main

import "fmt"

func main() {
	s := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(s))
}

func lengthOfLongestSubstring(s string) int {
	uniqueString := ""
	maxSize := 0

	for i := 0; i < len(s); i++ {
		letter := string(s[i])

		for j := 0; j < len(uniqueString); j++ {
			if string(uniqueString[j]) == letter {
				if len(uniqueString) > maxSize {
					maxSize = len(uniqueString)
				}
				uniqueString = uniqueString[j+1:]
			}
		}

		uniqueString += letter
	}

	if len(uniqueString) > maxSize {
		maxSize = len(uniqueString)
	}

	return maxSize
}
