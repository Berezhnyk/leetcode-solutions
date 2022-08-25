package longestpalindrome

func longestPalindrome(s string) string {
	slen := len(s)
	if len(s) == 0 {
		return ""
	}

	begin := 0
	end := 0
	for i := 0; i < slen; i++ {
		l1 := findMaxPalidromeFromSides(s, i, i)
		l2 := findMaxPalidromeFromSides(s, i, i+1)

		len := l1
		if l2 > l1 {
			len = l2
		}

		if len > end-begin {
			begin = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[begin : end+1]
}

func findMaxPalidromeFromSides(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		right++
		left--
	}

	return right - left - 1
}
