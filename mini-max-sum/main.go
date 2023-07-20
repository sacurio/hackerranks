package main

import "fmt"

/*
 * Complete the 'miniMaxSum' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */
var test []int32 = []int32{1, 3, 5, 7, 9}

func miniMaxSum(arr []int32) {
	var min, max int64
	var sum int64
	founded := 0
	for i := 0; i < len(arr); i++ {
		sum = 0
		founded = 0
		for j := 0; j < len(arr); j++ {
			if founded == 0 {
				if arr[i] == arr[j] {
					founded = 1
					continue
				}
			}
			sum += int64(arr[j])
		}

		switch {
		case i == 0:
			min = sum
			max = sum
		case sum > max:
			max = sum
		case sum < min:
			min = sum
		}
	}

	fmt.Println(min, max)
}

func main() {
	miniMaxSum(test)
}
