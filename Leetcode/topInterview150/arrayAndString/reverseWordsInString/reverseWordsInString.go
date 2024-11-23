package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("a good   example"))
}

func reverseWords(s string) string {
	wordsArray := strings.Split(s, " ")
	newString := ""
	for i := len(wordsArray) - 1; i >= 0; i-- {
		if wordsArray[i] != "" {
			newString += string(wordsArray[i])
			newString += " "
		}
	}

	return newString
}
