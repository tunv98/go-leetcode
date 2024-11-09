package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 5, removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

func removeDuplicates(nums []int) int {
	iLeft := 0
	for iRight := 1; iRight < len(nums); iRight++ {
		if nums[iLeft] != nums[iRight] {
			iLeft++
			nums[iLeft] = nums[iRight]
		}
	}
	return iLeft + 1
}

//iLeft luôn đứng đầu ở mỗi kí tự đầu tiên unique, và sẽ đợi thay thế kí tự kế tiếp nếu khác kí tự đầu
