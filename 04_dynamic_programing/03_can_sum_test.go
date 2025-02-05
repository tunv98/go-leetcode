package main

import (
	"fmt"
	"testing"
)

/*
Write a function `canSum(targetSum, numbers)` that takes in a targetSum and an array of numbers as arguments
The function should return a boolean indicating whether it is possible to generate the targetSum
using numbers from the array
You may use an element of the array as many times as needed
You may assume that all input numbers are non-negative
eg: canSum(7,[5,3,4,7]) -> true
*/

/*
Time: O(n^m)
Space:O(m)
n as array length, m as target sum
*/
func canSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, num := range numbers {
		remainder := targetSum - num
		if canSum(remainder, numbers) {
			return true
		}
	}
	return false
}

/*
Time: O(m*n)
Time: O(m)
*/
func canSumMemoization(targetSum int, numbers []int, memo map[int]bool) bool {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}
	for _, num := range numbers {
		remainder := targetSum - num
		if canSumMemoization(remainder, numbers, memo) {
			memo[targetSum] = true
			return true
		}
	}
	memo[targetSum] = false
	return false
}

func Test_CanSum(t *testing.T) {
	//fmt.Println(canSum(7, []int{5, 3, 5, 7}))
	//fmt.Println(canSum(7, []int{5, 3, 5, 3}))

	fmt.Println(canSumMemoization(7, []int{5, 3, 5, 7}, map[int]bool{}))
	fmt.Println(canSumMemoization(7, []int{5, 3, 5, 3}, map[int]bool{}))
	fmt.Println(canSumMemoization(7, []int{2, 3, 5, 3}, map[int]bool{}))
}
