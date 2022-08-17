package pivotindex

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
