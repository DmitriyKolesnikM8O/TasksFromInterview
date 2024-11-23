package main

import "fmt"

func main() {
	fmt.Println(mergeAlternately("abcd", "pq"))
}

func mergeAlternately(word1 string, word2 string) string {
	long := 0
	if len(word1) < len(word2) {
		long = -1
	} else if len(word1) > len(word2) {
		long = 1
	}

	resultStrings := ""
	if long == -1 || long == 0 {
		for i := 0; i < len(word1); i++ {
			resultStrings += string(word1[i])
			resultStrings += string(word2[i])
		}
	} else if long == 1 {
		for i := 0; i < len(word2); i++ {
			resultStrings += string(word1[i])
			resultStrings += string(word2[i])
		}
	}

	if long == -1 {
		resultStrings += string(word2[len(word1):])
	} else if long == 1 {
		resultStrings += string(word1[len(word2):])
	}

	return resultStrings
}
