package main

import "fmt"

var (
	ar       = []int32{1, 2, 3, 4, 5, 6}
	k  int32 = 5
	n  int32 = 6
)

func divisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var counter int32
	for i := 0; i < len(ar); i++ {
		if i+1 <= len(ar) {
			subset := ar[i+1:]
			for j := 0; j < len(subset); j++ {
				v := (ar[i] + subset[j]) % k
				if v == 0 {
					counter++
				}
			}
		}
	}
	return counter
}

func main() {
	fmt.Println(divisibleSumPairs(n, k, ar))
}
