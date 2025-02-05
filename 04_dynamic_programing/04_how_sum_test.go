package main

import (
	"fmt"
	"testing"
)

/*
Write a function should return an array containing any combination of elements that add up to exactly
the targetSum. if there is no combination that adds up to the targetSum, that return null
If there are multiple combinations possible, you may return any single one

howSum(7,[5,4,3,7] => [3,4], [7]
howSum(8,[2,3,5] => [3,5]
howSum(7,[2,4] => null
*/
/*
Time: O(n^m *m)
Space:O(m)
*/
func howSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return make([]int, 0)
	}
	if targetSum < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := targetSum - num
		remainderVal := howSum(remainder, numbers)
		if remainderVal != nil {
			return append(remainderVal, num)
		}
	}
	return nil
}

/*
Time: O(n*m^2)
Space:O(m^2)
*/
func howSumMemoization(targetSum int, numbers []int, memo map[int][]int) []int {
	if val, ok := memo[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return make([]int, 0)
	}
	if targetSum < 0 {
		return nil
	}

	for _, num := range numbers {
		remainder := targetSum - num
		remainderVal := howSumMemoization(remainder, numbers, memo)
		if remainderVal != nil {
			memo[targetSum] = append(memo[targetSum], num)
			return memo[targetSum]
		}
	}
	return nil
}
func TestHowSum(t *testing.T) {
	//fmt.Println(howSum(7, []int{5, 4, 3, 7}))
	fmt.Println(howSumMemoization(7, []int{5, 4, 3, 7}, map[int][]int{}))
	//fmt.Println(howSumMemoization(7, []int{2, 2, 2, 7}, map[int][]int{}))
}
