package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(34543))
}

func isPalindrome(x int) bool {
	return reverse(x) == x
}

func reverse(x int) int {
	var result int
	for x > 0 {
		result = result*10 + x%10
		x /= 10
	}
	return result
}
