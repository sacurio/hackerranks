package main

import "fmt"

func CharFrequency(str string) map[rune]int {
	freqMap := make(map[rune]int)
	for _, char := range str {
		freqMap[char]++
	}
	return freqMap
}

func main() {
	fmt.Println(CharFrequency("lorem ipsum ad dolorem"))
}
