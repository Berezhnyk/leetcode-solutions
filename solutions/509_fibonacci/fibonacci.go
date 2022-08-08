package fib

func fib(n int) int {
	if n == 0 {
		return 0
	}

	prev := 0
	current := 1
	result := 1

	for i := 1; i < n; i++ {
		result = prev + current
		prev = current
		current = result
	}

	return result
}
