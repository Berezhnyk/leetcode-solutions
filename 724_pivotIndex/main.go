package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}

func pivotIndex(nums []int) int {
	rightSum := 0
	for _, num := range nums {
		rightSum += num
	}

	leftSum := 0

	for i, num := range nums {
		rightSum -= num
		if rightSum == leftSum {
			return i
		}

		leftSum += num
	}

	return -1
}
