package main

import (
	"fmt"
	"strings"
)

func main() {

	string := "Let's take LeetCode contest"
	fmt.Println(reverseWords(string))
}

func reverseWords(s string) string {
	listWords := strings.Split(s, " ")

	result := ""

	for i := 0; i < len(listWords); i++ {
		newString := ""
		for j := len(listWords[i]) - 1; j >= 0; j-- {
			newString += string(listWords[i][j])
		}
		result += newString
		result += " "
	}

	return result[:len(result)-1]

}
