package lengthoflongestsubstring

import "strings"

func lengthOfLongestSubstring(s string) int {
	substring := ""
	maxSubstringLen := 0

	for i := 0; i < len(s); i++ {
		r := string(s[i])
		if ind := strings.Index(substring, r); ind >= 0 {
			substring = substring[ind+1:] + r
		} else {
			substring += r
		}

		substringLen := len(substring)
		if substringLen > maxSubstringLen {
			maxSubstringLen = substringLen
		}
	}

	return maxSubstringLen
}
