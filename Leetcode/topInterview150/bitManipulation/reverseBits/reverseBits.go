package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverseBits(4294967293))
}

func reverseBits(num uint32) uint32 {
	stringNum := strconv.FormatUint(uint64(num), 2)
	stringNum32bits := ""
	for len(stringNum)+len(stringNum32bits) < 32 {
		stringNum32bits += "0"
	}
	stringNum32bits += stringNum

	var result uint32 = 0
	for i := 0; i < len(stringNum32bits); i++ {

		if string(stringNum32bits[i]) == "1" {

			result += uint32(math.Pow(2, float64(i)))

		}

	}

	return result

}
