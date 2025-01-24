package random

import (
	"fmt"
	"sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	//merge1([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	//merge2([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	//merge2([]int{2, 0}, 1, []int{1}, 1)
	merge3([]int{2, 0}, 1, []int{1}, 1)
}

/*
Time: sort: O(nlogn)
*/
func merge1(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

/*
Time: O(n)
Space: O(n) => make new array
*/
func merge2(nums1 []int, m int, nums2 []int, n int) {
	i, j := 0, 0
	newArr := make([]int, 0, m+n)
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			newArr = append(newArr, nums1[i])
			i++
		} else {
			newArr = append(newArr, nums2[j])
			j++
		}
	}
	newArr = append(newArr, nums2[j:]...)
	for i < m {
		newArr = append(newArr, nums1[i])
		i++
	}
	copy(nums1, newArr)
	fmt.Println(nums1)
}

/*
Time: O(n)
Space: O(1)
*/
func merge3(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for j >= 0 {
		if i >= 0 && nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}
	fmt.Println(nums1)
}
