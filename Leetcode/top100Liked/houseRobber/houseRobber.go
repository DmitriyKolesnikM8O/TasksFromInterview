package main

import "fmt"

func main() {
	nums := []int{2, 1, 1, 2}
	fmt.Println(rob(nums))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	myList := make([]int, len(nums))
	myList[0] = nums[0]
	myList[1] = myMax(nums[0], nums[1])
	for i := 2; i < len(myList); i++ {
		myList[i] = myMax(myList[i-2]+nums[i], myList[i-1])
	}

	return myList[len(myList)-1]

}

func myMax(a int, b int) int {
	if a > b {
		return a
	}

	return b
}
