package runningsum

func runningSum(nums []int) []int {
	result := make([]int, 0)
	sum := 0
	for _, num := range nums {
		sum += num
		result = append(result, sum)
	}

	return result
}
