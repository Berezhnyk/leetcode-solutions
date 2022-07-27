package main

import "fmt"

type fn func(nums []int, k int) int

func main() {
	fmt.Println("Hello, playground")
	shouldBe(waysToPartition, []int{2, -1, 2}, 3, 1)
	shouldBe(waysToPartition, []int{0, 0, 0}, 1, 2)
	shouldBe(waysToPartition, []int{22, 4, -25, -20, -15, 15, -16, 7, 19, -10, 0, -13, -14}, -33, 4)
}

func waysToPartition(nums []int, k int) int {
	return k
}

func shouldBe(f fn, nums []int, k int, expected int) {
	r := f(nums, k)
	if r != expected {
		fmt.Printf("Expected %d, got %d\n", expected, r)
	} else {
		fmt.Println("OK")
	}
}
