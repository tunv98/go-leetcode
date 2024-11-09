package main

import (
	"reflect"
	"testing"
)

// https://leetcode.com/problems/two-sum/description/
func TestTwoSum(t *testing.T) {
	reflect.DeepEqual([]int{0, 1}, twoSum([]int{2, 7, 11, 15}, 9))
	reflect.DeepEqual([]int{0, 1}, twoSumWithTwoPointer([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	trackLogs := make(map[int]int)
	for i, v := range nums {
		if _, exists := trackLogs[v]; exists {
			return []int{i, trackLogs[v]}
		}
		trackLogs[target-v] = i
	}
	return []int{-1, -1}
}

//Time: O(n), Space: O(n)

func twoSumWithTwoPointer(nums []int, target int) []int { // use it sorted
	leftIndex := 0
	rightIndex := len(nums) - 1
	for leftIndex < rightIndex {
		total := nums[leftIndex] + rightIndex
		if total == target {
			return []int{leftIndex, rightIndex}
		}
		if total > target {
			rightIndex--
			continue
		}
		leftIndex++
	}
	return []int{-1, -1}
}

//Time: O(n), Space: O(1)
