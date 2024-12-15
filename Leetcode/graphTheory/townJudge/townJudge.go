package main

import "fmt"

func main() {
	trust := [][]int{{1, 3}, {2, 3}}
	fmt.Println(findJudge(3, trust))
}

func findJudge(n int, trust [][]int) int {
	if len(trust) < 1 {
		if n == 1 {
			return 1
		}
		return -1
	} else if len(trust) < 2 {
		return trust[0][1]
	}

	myMap := make(map[int][]int)
	for i := 0; i < len(trust); i++ {
		myMap[trust[i][1]] = append(myMap[trust[i][1]], trust[i][0])
	}

	for k, v := range myMap {
		if len(v) == n-1 {
			flag := true
			for k2, v2 := range myMap {
				if k2 != k {
					for _, i := range v2 {
						if k == i {
							flag = false
							break
						}
					}
				}
			}
			if flag {
				return k
			}
		}
	}

	// fmt.Println(myMap)

	return -1
}
