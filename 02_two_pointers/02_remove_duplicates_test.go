package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	assert.Equal(t, 5, removeDuplicates_2([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

// iLeft luôn đứng đầu ở mỗi kí tự đầu tiên unique, và sẽ đợi thay thế kí tự kế tiếp nếu khác kí tự đầu
func removeDuplicates(nums []int) int {
	n := len(nums)
	if n == 0 || n == 1 {
		return n
	}
	iLeft := 0
	for iRight := 1; iRight < n; iRight++ {
		if nums[iLeft] != nums[iRight] {
			iLeft++
			nums[iLeft] = nums[iRight]
		}
	}
	return iLeft + 1
}

// i là vị trí cần đổi chỗ, j là vị trí check
func removeDuplicates_2(nums []int) int {
	i, j := 1, 1
	for i < len(nums) {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
		j++
	}
	return i
}
