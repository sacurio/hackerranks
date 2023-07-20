package main

import "fmt"

/*
 * Complete the 'aVeryBigSum' function below.
 *
 * The function is expected to return a LONG_INTEGER.
 * The function accepts LONG_INTEGER_ARRAY ar as parameter.
 */

var test []int64 = []int64{
	1000000001, 1000000002, 1000000003, 1000000004, 1000000005,
}

func aVeryBigSum(ar []int64) int64 {
	var total int64
	for i := 0; i < len(ar); i++ {
		total += ar[i]
	}

	return total
}

func main() {
	fmt.Println(aVeryBigSum(test))
}
