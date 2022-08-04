package main

import (
	"fmt"
)

func Solution(X int, B []int, Z int) int {
	// write your code in Go 1.4
	var totalDownloaded = sum(B)
	var lastNumbers int
	var averageOfTheLastZMinutes int
	var totalRemainingBites = X - totalDownloaded
	if X == totalDownloaded {
		return 0
	}
	if len(B) <= 0 {
		return -1
	} else {
		lastNumbers = len(B) - Z
		var sum = sum(B[lastNumbers:])
		averageOfTheLastZMinutes = sum / Z
		return round(float32(totalRemainingBites) / float32(averageOfTheLastZMinutes))
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func round(num float32) int {
	return int(num + 0.5)
}

func main() {
	fmt.Println(Solution(100, []int{10, 6, 6, 8}, 2))
}
