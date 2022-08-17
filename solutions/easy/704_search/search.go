package search

func search(nums []int, target int) int {
	left := 0
	right := len(nums)
	var index int
	var current int
	for {
		index = ((right - left) / 2) + left
		current = nums[index]
		if current == target {
			return index
		}

		if left == right || left == index || right == index {
			return -1
		}

		if current >= target {
			right = index
		}

		if current <= target {
			left = index
		}
	}
}
