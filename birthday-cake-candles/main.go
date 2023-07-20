package main

/*
 * Complete the 'birthdayCakeCandles' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY candles as parameter.
 */
var test []int32 = []int32{4, 4, 1, 3}

func birthdayCakeCandles(candles []int32) int32 {
	var max, sum int32

	for _, height := range candles {
		if height > max {
			max = height
			sum = 1
		} else if height == max {
			sum++
		}
	}

	return sum
}

func main() {
	birthdayCakeCandles(test)
}
