package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2 //avoid overflow
		if nums[mid] == target {
			return mid
		}
		if target >= nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func Test_BinarySearch(t *testing.T) {
	assert.Equal(t, binarySearch([]int{1, 3, 5, 6, 8, 11}, 5), 2)
	assert.Equal(t, binarySearch([]int{1, 3, 5, 6, 8, 11}, 11), 5)
	assert.Equal(t, binarySearch([]int{1, 3, 5, 6, 8, 11, 12}, 11), 5)
}
