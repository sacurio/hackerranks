package main

import (
	"fmt"
	"strings"
)

/*
 * Complete the 'staircase' function below.
 *
 * The function accepts INTEGER n as parameter.
 */

func staircase(n int32) {
	nint := int(n)
	for i := 0; i < nint; i++ {
		fmt.Printf("%s%s\n", strings.Repeat(" ", nint-i-1), strings.Repeat("#", i+1))
	}
}

func main() {
	staircase(6)
}
