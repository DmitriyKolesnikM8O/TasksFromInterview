package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(myAtoi("10000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000522545459"))
}

func myAtoi(s string) int {

	result := 0
	sign := 1
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}

	if s[0] > '9' || s[0] < '0' {
		if s[0] == '-' {
			sign *= -1
		} else if s[0] == '+' {
			sign *= 1
		} else {
			return 0
		}
	} else {
		result = result*10 + int(s[0]-'0')
	}
	for i := 1; i < len(s); i++ {
		if s[i] > '9' || s[i] < '0' {
			break
		} else {
			result = result*10 + int(s[i]-'0')
			if result > 2147483647 && sign == 1 {
				return 2147483647
			}
			if -result < -2147483648 && sign == -1 {
				return -2147483647
			}
		}
	}

	result *= sign

	return result
}
