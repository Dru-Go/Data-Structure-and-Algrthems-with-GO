package main

func equilibriumIndex(arr []int32) int {
	if len(arr) == 0 {
		return -1
	}
	prefix := fillPrefixSum(arr, len(arr))
	sum := prefix[len(arr)-1]
	for i := 0; i < len(arr); i++ {
		leftSum := prefix[i] - arr[i]
		rightSum := sum - prefix[i]
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}

func efficientEquilibriumIndex(arr []int32) int {
	var totalSum, leftSum int32 = 0, 0
	if len(arr) == 0 {
		return -1
	}

	for _, v := range arr {
		totalSum += v
	}

	for i := 0; i < len(arr); i++ {
		rightSum := totalSum - leftSum - arr[i]
		if leftSum == rightSum {
			return i
		}
		leftSum += arr[i]
	}
	return -1
}

func fillPrefixSum(arr []int32, n int) []int32 {
	var prefixSum = make([]int32, n)
	prefixSum[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}
	return prefixSum
}
