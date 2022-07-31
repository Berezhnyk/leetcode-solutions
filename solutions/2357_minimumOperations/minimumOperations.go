package minimumoperations

func minimumOperations(nums []int) int {
	return len(uniqueNonZero(nums))
}

func uniqueNonZero(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value && entry != 0 {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
