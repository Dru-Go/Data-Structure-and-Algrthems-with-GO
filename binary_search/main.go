package main

func binarySearch(arr []int32, target int32) int {
	var low, high = 0, len(arr) - 1
	for i := low; i <= high; i++ {

		var mid = (low + high) / 2

		if arr[i] == target {
			return mid
		} else if target < arr[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func recursiveBinarySearch(array []int32, target int32, lowIndex int, highIndex int) int {
	startIndex := 0
	endIndex := len(array)

	var mid int

	for startIndex < endIndex {
		mid = int((lowIndex + highIndex) / 2)
		if array[mid] > target {
			return recursiveBinarySearch(array, target, lowIndex, mid)
		} else if array[mid] < target {
			return recursiveBinarySearch(array, target, mid+1, highIndex)
		} else {
			return mid
		}
	}
	return -1
}
