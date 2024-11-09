package main

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

//Input: arr[ ] = {3, 8, 9, 15}, K = 2
//Output: 6.5

func TestGetMinMaxDiff(t *testing.T) {
	arr := []int{3, 8, 9, 15}
	k := 2
	fmt.Println(getMinMaxDiff(arr, k))
}

func getMinMaxDiff(arr []int, k int) float64 {
	l := 0
	min, max := math.MaxInt, math.MinInt
	sum := 0
	for r := 0; r < len(arr); r++ {
		sum += arr[r]
		if r >= k-1 {
			min = utils.Min([]int{min, sum})
			max = utils.Max([]int{max, sum})
			sum -= arr[l]
			l++
		}
	}
	return (float64(max) - float64(min)) / float64(k)
}
