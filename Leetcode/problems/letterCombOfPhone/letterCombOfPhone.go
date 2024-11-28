package main

import "fmt"

func main() {
	digits := "2347"
	fmt.Println(letterCombinations(digits))

}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	mapDigits := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	result := []string{}

	var back func(comb string, nextDigit string)
	back = func(comb string, nextDigit string) {
		if len(nextDigit) == 0 {
			result = append(result, comb)
			return
		}

		current := nextDigit[0]
		stringLetter := mapDigits[current]

		for i := 0; i < len(stringLetter); i++ {
			back(comb+string(stringLetter[i]), nextDigit[1:])
		}
	}

	back("", digits)

	return result
}
