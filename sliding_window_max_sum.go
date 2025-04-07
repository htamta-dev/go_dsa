package main

import "fmt"

func sum[T int](arr []T, start T, end T) T {
	var sum T
	for i := start; i < end; i++ {
		sum += arr[i]
	}
	return sum
}

func max[T int](a T, b T) T {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindowSum(arr []int, windowSize int) int {
	if len(arr) < windowSize{
		return sum(arr, 0, len(arr))
	}
	var windowSum = sum(arr, 0, windowSize)
	var maxSum = windowSum
	for i:=0; i< len(arr)-windowSize; i++{
		windowSum = windowSum - arr[i] + arr[i+windowSize]
		maxSum = max(maxSum, windowSum)
	}
	return maxSum
}

func TestMaxSlidingWindowSum() {
	arr := []int{2, 1,5, 1, 3,2, 8, 4,1, 9}
	windowSize := 3
	result := maxSlidingWindowSum(arr, windowSize)
	fmt.Printf("Max sum of sliding window of size %d is: %d\n", windowSize, result)
}