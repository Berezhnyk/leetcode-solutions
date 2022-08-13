package longestcommonprefix

import "strings"

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		str := strs[i]
		for j := len(str); j >= 0; j-- {
			current := str[:j]
			if strings.HasPrefix(prefix, current) {
				prefix = current
				break
			}
			if current == "" {
				return ""
			}
		}

	}

	return prefix
}
