package main

func main() {
	println(fibonacci(14))
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
	return 0
}
