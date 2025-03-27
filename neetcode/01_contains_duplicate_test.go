package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// https://leetcode.com/problems/contains-duplicate/description/

/*
if nums has many int
eg: 10^9
key-value contain: sizeof(int) + sizeof(bool) = 8 + 1 = 9 bytes
=> 10^9 (num) * 9(bytes) = 9GB
=> consume too much memory => Bitset
*/
func containsDuplicateMap(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	existedMap := make(map[int]bool, len(nums))
	for _, num := range nums {
		if existedMap[num] {
			return true
		}
		existedMap[num] = true
	}
	return false
}

/*
uint 64 -> 64bit => Each element in bitset can save 64 integer
Store 10^9 integer => number of element in bitset
-> (10^9)/64 =15,625,000 element (uint64)
=> 15,625,000 * 8 (bytes) ~ 125MB
*/
func containDuplicateBitSet(nums []int) bool {
	const MaxValue = 1_000_000_000
	bitset := make([]uint64, MaxValue/64+1) // Tạo mảng bitset

	for _, num := range nums {
		index, bit := num/64, uint(num%64)
		if bitset[index]&(1<<bit) != 0 {
			return true
		}
		bitset[index] |= 1 << bit
	}
	return false
}

func Test_ContainDuplicate(t *testing.T) {
	assert.True(t, containsDuplicateMap([]int{1, 2, 3, 4, 3}))
	assert.False(t, containsDuplicateMap([]int{1, 2, 3, 4}))

	assert.True(t, containDuplicateBitSet([]int{1, 2, 3, 4, 3}))
	assert.False(t, containDuplicateBitSet([]int{1, 1, 3, 4}))
}
