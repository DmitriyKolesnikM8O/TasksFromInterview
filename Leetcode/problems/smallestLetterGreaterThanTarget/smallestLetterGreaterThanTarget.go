package main

import "fmt"

func main() {
	letters := []byte{'c', 'f', 'j'}
	target := 'z'

	fmt.Println(string(nextGreatestLetter(letters, byte(target))))
}

func nextGreatestLetter(letters []byte, target byte) byte {

	start := 0
	end := len(letters) - 1
	for start <= end {
		med := (end + start) / 2

		if letters[med] <= target {
			start = med + 1
		} else {
			if (med > 0) && (letters[med-1] > target) {
				end = med - 1
			} else {
				return letters[med]
			}
		}
	}

	return letters[0]
}
