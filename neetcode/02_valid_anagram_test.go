package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/valid-anagram/description/

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	fre := make([]int32, 26)
	for _, c := range s {
		fre[c-'a']++
	}
	for _, c := range t {
		fre[c-'a']--
	}
	for _, i := range fre {
		if i != 0 {
			return false
		}
	}
	return true
}

/*
Use [26]int{} instead of make([]int, 26)
Tối ưu bộ nhớ: Mảng tĩnh [26]int{} nằm trên stack, nhanh hơn slice make([]int, 26) nằm trên heap.
*/
func isAnagramOptimize(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	freq := [26]int{}
	for i := 0; i < len(s); i++ {
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}
	for _, v := range freq {
		if v != 0 {
			return false
		}
	}
	return true
}

func Test_IsAnagram(t *testing.T) {
	assert.True(t, isAnagram("anagram", "nagaram"))
	assert.False(t, isAnagram("rat", "car"))

	assert.True(t, isAnagramOptimize("anagram", "nagaram"))
	assert.False(t, isAnagramOptimize("rat", "car"))
}
