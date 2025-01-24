package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

/* //brute force
func lengthOfLongestSubstring(s string) int {
	maxLength := -1
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if isUnique(s, i, j) {
				maxLength = maxFn(maxLength, j-i+1)
			}
		}
	}
	return maxLength
}
func maxFn(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func isUnique(s string, start, end int) bool {
	seen := make(map[byte]struct{})
	for i := start; i <= end; i++ {
		if _, ok := seen[s[i]]; ok {
			return false
		}
		seen[s[i]] = struct{}{}
	}
	return true
}
*/

func lengthOfLongestSubstring(s string) int {
	l := 0
	maxLength := 0
	seen := make(map[byte]int) // lưu vị trí xuất hiện gần nhất của mỗi ký tự
	for r := 0; r < len(s); r++ {
		if dup, existed := seen[s[r]]; existed && l <= dup {
			l = dup + 1
		}
		seen[s[r]] = r
		maxLength = max(maxLength, r-l+1)
	}
	return maxLength
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, lengthOfLongestSubstring("pwwkew"), 3)
	assert.Equal(t, lengthOfLongestSubstring(""), 0)
	assert.Equal(t, lengthOfLongestSubstring("dvdf"), 3)
	assert.Equal(t, lengthOfLongestSubstring("abba"), 2) // l <= dup to avoid removing duplicate before
}
