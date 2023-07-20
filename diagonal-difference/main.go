package main

import (
	"fmt"
	"math"
)

/*
 * Complete the 'diagonalDifference' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts 2D_INTEGER_ARRAY arr as parameter.
 */

type matrixElement []int32

var matrix [][]int32 = [][]int32{
	matrixElement{1, 2, 3},
	matrixElement{4, 5, 6},
	matrixElement{9, 8, 9},
}

func diagonalDifference(arr [][]int32) int32 {
	var diagonalA, diagonalB int32
	for i := 0; i < len(arr[0]); i++ {
		diagonalA += arr[i][i]
		diagonalB += arr[i][len(arr)-1-i]
		continue
	}
	return int32(math.Abs(float64(diagonalA) - float64(diagonalB)))
}

func main() {
	fmt.Println(diagonalDifference(matrix))
}
