package main

import "testing"

//https://leetcode.com/problems/top-k-frequent-elements/
/*
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
*/
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int)
	for _, num := range nums {
		freq[num]++
	}
	return nil
}

func Test_TopKFrequent(t *testing.T) {

}
