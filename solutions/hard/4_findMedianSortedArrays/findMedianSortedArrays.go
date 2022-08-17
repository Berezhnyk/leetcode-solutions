package findmediansortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	max1 := len(nums1) - 1
	max2 := len(nums2) - 1
	resultArray := make([]int, 0)
	i1 := 0
	i2 := 0
	for {
		if i1 > max1 || max1 == -1 {
			resultArray = append(resultArray, nums2[0:]...)
			resultArray = append(resultArray, nums1...)
			break
		} else if i2 > max2 || max2 == -1 {
			resultArray = append(resultArray, nums1[0:]...)
			resultArray = append(resultArray, nums2...)
			break
		}

		n1 := nums1[0]
		n2 := nums2[0]

		if n1 < n2 {
			resultArray = append(resultArray, n1)
			i1 += 1
			nums1 = nums1[1:]
		} else if n1 > n2 {
			resultArray = append(resultArray, n2)
			i2 += 1
			nums2 = nums2[1:]
		} else {
			resultArray = append(resultArray, n1, n2)
			i1 += 1
			i2 += 1
			nums1 = nums1[1:]
			nums2 = nums2[1:]
		}

	}

	resLen := len(resultArray)
	if resLen == 1 {
		return float64(resultArray[0])
	}

	if resLen%2 == 0 {
		middleLeft := resLen/2 - 1
		middleRight := middleLeft + 1
		return float64(resultArray[middleLeft]+resultArray[middleRight]) / 2
	} else {
		return float64(resultArray[resLen/2])
	}
}
