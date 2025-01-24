package random

import (
	"fmt"
	"sort"
	"testing"
)

/*
space : O(n1)
time: O(n1+n2)
*/
func intersect(nums1 []int, nums2 []int) []int {
	counter := make(map[int]int, len(nums1))
	var rs []int
	for _, i := range nums1 {
		counter[i]++
	}
	for _, i := range nums2 {
		if counter[i] != 0 {
			counter[i]--
			rs = append(rs, i)
		}
	}
	return rs
}

// What if the given array is already sorted? How would you optimize your algorithm?
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	var result []int
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] == nums2[j] {
			result = append(result, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}
	return result
}

func Test_intersect(t *testing.T) {
	//fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
	fmt.Println(intersect2([]int{1, 2, 2, 1}, []int{2, 2}))
}
