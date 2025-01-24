package random

import (
	"fmt"
	"sort"
	"testing"
)

// https://leetcode.com/problems/merge-intervals/submissions/1499275413/
/*
Time: O(nlogn)
Space: O(n)
*/
func mergeIntervals(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	result := make([][]int, 0)
	merged := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if merged[1] >= intervals[i][0] {
			merged = []int{merged[0], max(merged[1], intervals[i][1])}
			continue
		}
		result = append(result, merged)
		merged = intervals[i]
	}
	result = append(result, merged)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_MergeInterval(t *testing.T) {
	//fmt.Println(mergeIntervals([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	//fmt.Println(mergeIntervals([][]int{{1, 4}, {4, 5}}))
	//fmt.Println(mergeIntervals([][]int{{1, 4}, {0, 4}}))
	fmt.Println(mergeIntervals([][]int{{1, 4}, {2, 3}}))
}
