package twosum

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
