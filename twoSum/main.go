package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	numsArr := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		numsArr[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		if val, ok := numsArr[target-nums[i]]; ok && i != val {
			return []int{i, val}
		}
	}

	return nil
}
