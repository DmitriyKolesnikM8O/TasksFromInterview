package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	htmlCode, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("error read input data")
	}

	counter := 0
	words := strings.Split(htmlCode, "</")
	for _, line := range words {
		if strings.Contains(line, "href") {
			indexLast := strings.LastIndex(line, ">") + 1
			fmt.Println(line[indexLast:])
			counter++
		}
	}
	fmt.Println(counter)
}
