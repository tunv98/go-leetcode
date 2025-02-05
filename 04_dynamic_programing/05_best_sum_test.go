package main

import (
	"fmt"
	"testing"
)

/*
Write a function should return an array containing the shortest combination of numbers
that add up to exactly the targetSum
If there is a tie for the shortest combination, you may return any one of the shortest

bestSum(7, [5,3,4,7]) =>[3,4] , [7] => [7]
*/

func bestSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return make([]int, 0)
	}
	if targetSum < 0 {
		return nil
	}
	var shortestCombination []int

	for _, num := range numbers {
		remainder := targetSum - num
		remainderVal := bestSum(remainder, numbers)
		if remainderVal != nil {
			combination := append(remainderVal, num)
			if shortestCombination == nil || len(combination) < len(shortestCombination) {
				shortestCombination = combination
			}
		}
	}
	return shortestCombination
}

func TestBestSum(t *testing.T) {
	fmt.Println(bestSum(7, []int{5, 3, 4, 7}))
}
