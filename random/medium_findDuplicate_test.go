package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	// Giai đoạn 1: Tìm điểm gặp nhau trong chu trình
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	// Giai đoạn 2: Tìm điểm bắt đầu của chu trình
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

func Test_findDuplicate(t *testing.T) {
	//assert.Equal(t, findDuplicate([]int{3, 1, 4, 2, 3}), 3)
	//assert.Equal(t, findDuplicate([]int{1, 3, 4, 2, 2}), 2)
	assert.Equal(t, findDuplicate([]int{3, 1, 4, 2, 2}), 2)
}
