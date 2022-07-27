package main

import "strings"

func main() {
	s := "abcde"
	words := []string{"a", "bb", "acd", "ace"}
	println(numMatchingSubseq(s, words))
}

func numMatchingSubseq(s string, words []string) int {
	num := 0
	for _, word := range words {
		searchString := s
		for j, letter := range word {
			position := strings.Index(searchString, string(letter))
			if position == -1 {
				break
			} else {
				searchString = searchString[position+1:]
			}
			if j == len(word)-1 {
				num++
			}
		}
	}

	return num
}
