package main

import (
	"fmt"
)

func main() {
	var originalText, key, action string
	fmt.Scanln(&originalText)
	fmt.Scanln(&key)
	fmt.Scanln(&action)

	keyLongText := key
	for i := 0; i < len(originalText)-len(key); i++ {
		keyLongText += string(key[i%len(key)])
	}

	if action == "encode" {
		shoferText := ""
		for i := 0; i < len(keyLongText); i++ {
			newDigit := int(originalText[i]) + int(keyLongText[i])%int('A')
			if newDigit > 90 {
				newDigit = newDigit%90 + 64
			}
			shoferText += string(newDigit)
		}

		fmt.Println(shoferText)

	} else if action == "decode" {
		deshoferText := ""
		for i := 0; i < len(keyLongText); i++ {
			newDigit := int(originalText[i]) - int(keyLongText[i])%int('A')
			if newDigit < 65 {
				newDigit += 26
			}
			deshoferText += string(newDigit)
		}

		fmt.Println(deshoferText)
	}

	return
}
