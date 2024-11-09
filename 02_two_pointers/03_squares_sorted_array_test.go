package main

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	reflect.DeepEqual([]int{0, 1, 4, 9, 25, 100}, sortedSquares([]int{-5, -2 - 1, 0, 3, 10}))
}

func sortedSquares(nums []int) []int {
	iLeft, iRight := 0, len(nums)-1
	result := make([]int, len(nums), len(nums))
	for i := 0; iLeft != iRight; i++ {
		leftSquare := nums[iLeft] * nums[iLeft]
		rightSquare := nums[iRight] * nums[iRight]
		if leftSquare > rightSquare {
			result[len(nums)-1-i] = leftSquare
			iLeft++
			continue
		}
		result[len(nums)-1-i] = rightSquare
		iRight--
	}
	result[0] = nums[iRight] * nums[iRight]
	return result
}
