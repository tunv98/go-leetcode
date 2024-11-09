package main

import (
	"fmt"
	"leetcode/utils"
	"math"
	"testing"
)

// Input: nums = [1,12,-5,-6,50,3], k = 4
// Output: 12.75000

func TestFindMaxAverage(t *testing.T) {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	//fmt.Println(ByBruteForce(nums, k))
	fmt.Println(BySlidingWindow(nums, k))
}

func ByBruteForce(nums []int, k int) float64 {
	max := 0.0
	for i := 0; i <= len(nums)-k; i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += nums[j]
		}
		max = math.Max(float64(sum)/float64(k), max)
	}
	return max
}

//O(n * k), n is the number of elements in the input array
//Sự kém hiệu quả khi cần phải loop lại dữ liệu, phần overlap (3 phần tử) được tính lại 2 lần
//=> Để tận dụng lại tổng từ mảng con trước đó, chúng ta sẽ trừ phần tử trước và thêm vào phần tử sau đưa vào siding window
//=> Độ phức tạp chỉ là O(n)

func BySlidingWindow(nums []int, k int) float64 {
	sum := 0
	max := 0
	l := 0
	for r := 0; r < len(nums); r++ {
		sum += nums[r]
		if r >= k-1 {
			max = utils.Max([]int{max, sum})
			sum -= nums[l]
			l++
		}
	}
	return float64(max) / float64(k)
}
