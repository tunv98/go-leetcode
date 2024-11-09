package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//https://leetcode.com/problems/first-bad-version/

func isBadVersion(n int) bool {
	return n >= 4
}
func firstBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1 //do the first true will return left
		}
	}
	return left
}

func Test_firstBadVersion(t *testing.T) {
	assert.Equal(t, firstBadVersion(10), 4)
	assert.Equal(t, firstBadVersion(4), 4)
}
