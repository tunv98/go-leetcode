package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Given an array of size n, find the majority element.
//The majority element is the element that appears more than ⌊ n/2 ⌋ times.
//You may assume that the array is non-empty and the majority element always exist in the array.

//Example
//Array: [2, 2, 1, 1, 1, 2, 2] => n=7 => [7/2] = 3 => 2 appears 4 times > [7/2]=3 => 2 is majority

// Time: O(n),Space: O(n)
func majorityElementByMap(nums []int) int {
	counter := make(map[int]int)
	threshold := len(nums) / 2
	for _, i := range nums {
		counter[i]++
		if counter[i] >= threshold {
			return nums[i]
		}
	}
	return -1
}

// Time: O(n),Space: O(1)
func majorityElementByBoyerMooreVoting(nums []int) int {
	candidate, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate, count = nums[i], 1
		} else if nums[i] == candidate {
			count++
		} else {
			count--
		}
	}

	count = 0
	for _, num := range nums {
		if num == candidate {
			count++
		}
	}

	if count > len(nums)/2 {
		return candidate
	}
	return -1
}

func Test_majorityElement(t *testing.T) {
	assert.Equal(t, 2, majorityElementByMap([]int{2, 2, 1, 1, 1, 2, 2}))
	assert.Equal(t, 2, majorityElementByBoyerMooreVoting([]int{2, 2, 1, 1, 1, 2, 2}))
}
