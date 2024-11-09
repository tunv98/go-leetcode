package main

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

//https://leetcode.com/problems/minimum-size-subarray-sum/description/

func TestMinSubArrayLen(t *testing.T) {
	nums := []int{2, 3, 1, 3, 4, 3}
	k := 7
	fmt.Println(minSubArrayLen(k, nums))
}

func minSubArrayLen(target int, nums []int) int {
	l := 0
	min := math.MaxInt
	sum := 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		for sum >= target {
			min = utils.Min([]int{min, r - l + 1})
			//min = int(math.Min(float64(min), float64(r-l+1)))
			sum -= nums[l]
			l++
		}
	}
	if min == math.MaxInt {
		return 0
	}
	return min
}
