package fibonacci

//Naive approach
func fibonacci(n int64) int64 {
	if n <= 1 {return n}
	return fibonacci(n -1) + fibonacci(n - 2)
}

//faster approach
func fasterFibonacci(n int64) int64 {
	if n == 0 {return 0}
	var arr = make([]int64, n + 1)
	arr[0], arr[1] = 0, 1

	for i := int64(2); i <= n; i++ {
		arr[i] = arr[i - 1] + arr[i - 2]
	}
	return arr[n]
}

