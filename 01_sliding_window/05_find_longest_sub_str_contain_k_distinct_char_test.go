package main

import (
	"fmt"
	"testing"
)

//Input: s = 'abcbdbdbbdcdabd' , k = 2
//Output: bdbdbbd

func TestGetLongest(t *testing.T) {
	s := "abcbdbdbbdcdabd"
	k := 2
	fmt.Println(getLongest(s, k))
}

func getLongest(s string, k int) string {
	if len(s) == 0 || k <= 0 {
		return ""
	}
	charCount := make(map[byte]int)
	start := 0
	maxLength := 0
	maxStart := 0
	for end := 0; end < len(s); end++ {
		charCount[s[end]]++
		for len(charCount) > k {
			charCount[s[start]]--
			if charCount[s[start]] == 0 {
				delete(charCount, s[start])
			}
			start++
		}
		if end-start+1 > maxLength {
			maxLength = end - start + 1
			maxStart = start
		}
	}
	return s[maxStart : maxStart+maxLength]
}

/*
Độ phức tạp của thuật toán trượt cửa sổ trong bài toán này là O(n), trong đó n là độ dài của chuỗi đầu vào.
Thuật toán duyệt qua chuỗi một lần, mỗi lần thêm một ký tự mới vào cửa sổ
và di chuyển cửa sổ nếu số ký tự khác nhau trong cửa sổ vượt quá k.
Mỗi ký tự được thêm hoặc loại bỏ từ cửa sổ tối đa một lần.
Do đó, mỗi ký tự được xử lý tối đa hai lần (một khi thêm vào và một khi loại bỏ khỏi cửa sổ).
Do đó, thời gian thực hiện của thuật toán là O(n), trong đó n là độ dài của chuỗi đầu vào.
*/
