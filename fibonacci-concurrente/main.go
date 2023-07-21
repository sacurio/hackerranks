package main

import (
	"fmt"
)

func main() {
	fmt.Println(fibonacci(15))
	fmt.Println(fibonacciRecursive(15))
	fmt.Println(fibonacciConcurrent(15))
}

func fibonacci(n int) []int {
	seq := make([]int, n)
	seq[0] = 0
	if n == 1 {
		return seq
	}
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

func fibonacciConcurrent(n int) int {
	if n <= 1 {
		return n
	}

	ch := make(chan int)
	go concurrentHelper(n, ch)
	return <-ch
}

func concurrentHelper(n int, ch chan int) {
	if n <= 1 {
		ch <- n
		return
	}

	ch1 := make(chan int)
	ch2 := make(chan int)

	go concurrentHelper(n-1, ch1)
	go concurrentHelper(n-2, ch2)

	x, y := <-ch1, <-ch2

	ch <- x + y
}
