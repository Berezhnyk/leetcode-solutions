package main

import "fmt"

func main() {
	fmt.Println(runningSum([]int{1, 2, 3, 4}))
}

func runningSum(nums []int) []int {
	result := make([]int, 0)
	sum := 0
	for _, num := range nums {
		sum += num
		result = append(result, sum)
	}

	return result
}
