package main

import "fmt"

/*
 * Complete the 'plusMinus' function below.
 *
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

var test []int32 = []int32{-4, 3, -9, 0, 4, 1}

func plusMinus(arr []int32) {
	var positives, negatives, zeros int
	l := float32(len(arr))

	for _, num := range arr {
		switch {
		case num < 0:
			negatives++
		case num == 0:
			zeros++
		case num > 0:
			positives++
		}
	}

	fmt.Printf("%.6f\n", float32(positives)/l)
	fmt.Printf("%.6f\n", float32(negatives)/l)
	fmt.Printf("%.6f\n", float32(zeros)/l)
}

func main() {
	plusMinus(test)
}
