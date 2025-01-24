package random

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

/*
đọc đề bài thấy
s and t consist of lowercase English letters.
=> theo ascii thì chỉ có 26 kí tự bắt đầu từ a
=> nên khi sử dụng map hay array có thể bắt đầu khi kí tự - 'a' để gán
O(26) ~~ O(1)
*/
func sortString(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

/*
Complexity
Time: O(nlogn)
Space: O(1)
*/
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sortedS := sortString(s)
	sortedT := sortString(t)
	if sortedS != sortedT {
		return false
	}
	return true
}

/*
Complexity
Time: O(n)
Space: O(n)
*/
func isAnagram2(s string, t string) bool {
	m := make(map[int32]int)
	for _, i := range s {
		m[i]++
	}
	for _, j := range t {
		m[j]--
		if m[j] == 0 {
			delete(m, j)
		}
	}
	if len(m) != 0 {
		return false
	}
	return true
}

func isAnagram3(s string, t string) bool {
	m := make(map[uint8]int, 26)
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
		m[t[i]-'a']--
		if m[s[i]-'a'] == 0 {
			delete(m, s[i]-'a')
		}
		if m[t[i]-'a'] == 0 {
			delete(m, t[i]-'a')
		}
	}
	if len(m) != 0 {
		return false
	}
	return true
}

func isAnagram4(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	fre := make([]int, 26)
	for i := 0; i < len(s); i++ {
		fre[s[i]-'a']++
		fre[t[i]-'a']--
	}
	for i := 0; i < 26; i++ {
		if fre[i] != 0 {
			return false
		}
	}
	return true
}

func Test_isAnagram(t *testing.T) {
	assert.True(t, isAnagram1("abccba", "aabbcc"))
	assert.False(t, isAnagram1("abccbak", "aabbcc"))

	assert.True(t, isAnagram2("abccba", "aabbcc"))
	assert.False(t, isAnagram2("abccbak", "aabbcc"))

	assert.True(t, isAnagram3("abccba", "aabbcc"))
	assert.False(t, isAnagram3("abccbak", "aabbcc"))

	assert.True(t, isAnagram4("abccba", "aabbcc"))
	assert.False(t, isAnagram4("abccbak", "aabbcc"))
}
