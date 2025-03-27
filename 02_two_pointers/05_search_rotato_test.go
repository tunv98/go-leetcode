package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Description:
You are given an integer array nums sorted in ascending order,
but it has been rotated at some unknown pivot. You need to search for a target element in this array.
The requirement is that your algorithm should run in O(log n) time complexity.

Example:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Solution: The brute-force solution would involve iterating over the entire array, but that's O(n) time complexity.
Instead, I used a modified binary search approach to solve it in O(log n) time.
*/

// Khi được xoay thì trong mảng luôn tồn tại 1 phần được sx, và 1 phần ko được sx
func searchInRotatedSortedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// if left part is sorted
		if nums[mid] >= nums[left] {
			if nums[left] <= target && target < nums[mid] { // if target in [left, mid]
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // if right part is sorted
			if nums[mid] < target && target <= nums[right] { // if target in [mid, right]
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func Test_SearchInRotatedSortedArray(t *testing.T) {
	assert.Equal(t, 4, searchInRotatedSortedArray([]int{4, 5, 6, 7, 0, 1, 2}, 0))
}
