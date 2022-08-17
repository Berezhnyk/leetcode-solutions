package longestpalindrome

func longestPalindrome(s string) int {
	m := make(map[string]int)
	i := 0
	for i < len(s) {
		char := string(s[i])
		if v, ok := m[char]; ok {
			v++
			m[char] = v
		} else {
			m[char] = 1
		}
		i++
	}

	palindromeLength := 0
	hasOddCharsCount := false
	for _, charsCount := range m {
		if charsCount%2 == 0 {
			palindromeLength += charsCount
		} else {
			palindromeLength += charsCount - 1
			if !hasOddCharsCount {
				hasOddCharsCount = true
			}
		}
	}

	if hasOddCharsCount {
		palindromeLength++
	}

	return palindromeLength
}
