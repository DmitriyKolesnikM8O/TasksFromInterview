package main

import "fmt"

func main() {
	fmt.Println(reverse(-2147483412))
}

func reverse(x int) int {
	sign := 1
	if x <= -2147483648 || x > 2147483647 {
		return 0
	}

	if x < 0 { //works only with positive numbers
		sign *= -1
		x *= -1
	}

	result := 0
	for x > 0 {
		result = result*10 + (x % 10)
		x /= 10
	}

	if result <= -2147483412 || result >= 2147483647 {
		return 0
	}

	return sign * result
}
