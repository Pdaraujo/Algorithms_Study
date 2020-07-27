package fibonacci

//Naive approach
func fibonacci(n int64) int64 {
	if n <= 1 {return n}
	return fibonacci(n -1) + fibonacci(n - 2)
}

