package main

import "fmt"

func fillPrefixSum(arr []int32, n int) []int32 {
	var prefixSum = make([]int32, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	return prefixSum
}

func prefixSum(arr []int32) int64 {
	var n = len(arr)
	if n == 0 {
		return 0
	}
	sums := fillPrefixSum(arr, n)
	fmt.Println(sums)
	return int64(sums[len(sums)-1])
}
