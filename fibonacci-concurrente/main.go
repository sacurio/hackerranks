package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(15))
	fmt.Println(fibonacciRecursive(15))
}

func fibonacci(n int) []int {
	seq := make([]int, n)
	seq[0] = 0
	seq[1] = 1
	for i := 2; i < n; i++ {
		seq[i] = seq[i-1] + seq[i-2]
	}

	return seq
}

func fibonacciRecursive(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacciRecursive(n-1) + fibonacciRecursive(n-2)
}
