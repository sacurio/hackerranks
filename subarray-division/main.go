package main

import "fmt"

var (
	data       = []int32{1, 2, 1, 3, 2}
	d    int32 = 3
	m    int32 = 2
)

func main() {
	fmt.Println(birthday(data, d, m))
}

func birthday(s []int32, d int32, m int32) int32 {
	var counter int32

	var sumSubset = func(s []int32) int32 {
		var total int32
		for _, v := range s {
			total += v
		}
		return total
	}

	for i := 0; i < len(s); i++ {
		if (int(m)-1)+i <= len(s)-1 {
			if d == sumSubset(s[i:int(m)+i]) {
				counter++
			}
		}
	}

	return counter
}
