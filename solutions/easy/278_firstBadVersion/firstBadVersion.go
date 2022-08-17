package firstbadversion

func isBadVersionFunc(version int) bool {
	return false
}

var isBadVersion = isBadVersionFunc

func firstBadVersion(n int) int {
	left := 0
	right := n
	for {
		version := (right-left)/2 + left
		if isBadVersion(version) {
			right = version
		} else {
			left = version
		}

		if right-left == 1 {
			return right
		}
	}
}
