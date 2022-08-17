package isvalid

import "strings"

func isValid(s string) bool {
	open := "([{"

	brackets := make(map[rune]rune, 3)
	brackets['}'] = '{'
	brackets[']'] = '['
	brackets[')'] = '('

	stack := []rune{}

	for _, current := range s {
		if strings.ContainsRune(open, current) {
			stack = append(stack, current)
		} else {
			lastIndex := len(stack) - 1
			if lastIndex == -1 {
				return false
			}
			last := stack[lastIndex]
			if brackets[current] == last {
				stack[lastIndex] = 0
				stack = stack[:lastIndex]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}
