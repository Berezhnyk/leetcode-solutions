package issubsequence

import "strings"

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	for _, c := range s {
		index := strings.Index(t, string(c))
		if index == -1 {
			return false
		}
		t = t[index+1:]
	}

	return true
}
