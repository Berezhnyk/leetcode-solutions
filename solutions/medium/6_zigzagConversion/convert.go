package zigzagconversion

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var transformed [][]rune = make([][]rune, numRows)
	for row := range transformed {
		transformed[row] = make([]rune, len(s)/(numRows-1)+numRows)
	}

	i := 0
	j := 0

	direction := 1 // 1 is down, -1 is up
	for _, r := range s {
		transformed[i][j] = r
		i += direction
		if i < 0 || i == numRows {
			if direction == 1 {
				j++
			}
			direction *= -1
			i += direction * 2
		}
		if direction == -1 && i != numRows-2 {
			j++
		}
	}

	var result = make([]rune, 0)
	for i, raw := range transformed {
		for j := range raw {
			if transformed[i][j] != rune(0) {
				result = append(result, transformed[i][j])
			}
		}
	}

	return string(result)
}
