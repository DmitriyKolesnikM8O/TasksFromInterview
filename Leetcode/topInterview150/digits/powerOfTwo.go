package main

import "fmt"

func main() {
	n := 3
	fmt.Println(isPowerOfTwo(n))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	} else if n == 1 {
		return true
	} else if n < 0 {
		return false
	}

	ostatok := 2
	for ostatok < n {
		if n%ostatok != 0 {
			return false
		}
		ostatok *= 2
	}

	return true
}
