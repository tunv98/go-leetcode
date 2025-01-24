package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

//https://leetcode.com/problems/3sum/

// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
// sum = 0

/*
Sort -> Move from n-1 => 1 to achieve third element
sort: O(nlogn)
loop: O(n)
*/
func thirdMax_1(nums []int) int {
	sort.Ints(nums)
	distinctCount := 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			distinctCount++
			if distinctCount == 3 {
				return nums[i]
			}
		}
	}
	return nums[len(nums)-1]
}

/*
use for to find max element x 3
Time: O(n)
Space: O(1
*/
func thirdMax_2(nums []int) int {
	first, second, third := math.MinInt32, math.MinInt32, math.MinInt32
	count := 0
	for _, num := range nums {
		if (num == first || num == second || num == third) && num != math.MinInt32 {
			continue
		}
		if num >= first {
			first, second, third = num, first, second
			count++
			continue
		}
		if num >= second {
			second, third = num, second
			count++
			continue
		}
		if num >= third {
			count++
			third = num
		}
	}
	if count < 3 {
		return first
	}
	return third
}

func TestThreeSum(t *testing.T) {
	//fmt.Println(thirdMax_1([]int{2, 4, 5, 6, 7, 7, 8, 8}))
	fmt.Println(thirdMax_2([]int{-2147483648, -2147483648, -2147483648, -2147483648, 1, 1, 1}))
}
