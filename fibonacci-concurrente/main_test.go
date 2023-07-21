package main

import "testing"

func BenchmarkFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci(20)
	}
}

func BenchmarkFibonacciRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciRecursive(20)
	}
}

func BenchmarkFibonacciConcurrent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciConcurrent(20)
	}
}
