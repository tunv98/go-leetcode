package random

import (
	"fmt"
	"testing"
)

func intersection(nums1 []int, nums2 []int) []int {
	mySet := make(map[int]struct{})
	for _, num := range nums1 {
		mySet[num] = struct{}{}
	}
	resultSet := make(map[int]struct{})
	for _, num := range nums2 {
		if _, exists := mySet[num]; exists {
			resultSet[num] = struct{}{}
		}
	}
	result := make([]int, 0, len(resultSet))
	for num := range resultSet {
		result = append(result, num)
	}
	return result
}

func Test_intersection(t *testing.T) {
	fmt.Println(intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
