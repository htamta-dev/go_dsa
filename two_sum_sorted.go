package main

import (
	"fmt"
	"sort"
)

func twoSumSorted(arr []int, target int) []int {
	type IndexedNum struct{
		Value int
		Index int
	}

	indexedNums := make([]IndexedNum, len(arr))

	for i, num := range(arr){
        indexedNums[i] = IndexedNum{ Value: num, Index: i}
	}

	sort.Slice(indexedNums, func(i, j int) bool{
		return indexedNums[i].Value < indexedNums[j].Value
	})

	left, right := 0, len(arr)-1

	for left < right {
		curr_sum := indexedNums[left].Value + indexedNums[right].Value

		if curr_sum == target {
			return []int{ indexedNums[left].Index, indexedNums[right].Index}
		} else if curr_sum < target {
			left += 1
		} else {
			right -= 1
		}
	}
	return []int{}
}

func TestTwoSumSorted(){
	arr := []int{2, 7, 11, 15}
	target := 9
	result := twoSumSorted(arr, target)
	fmt.Println("Two sum sorted result : ", result);
}