package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func firstUniqChar(s string) int {
	freq := make(map[int32]int)
	for _, v := range s {
		freq[v]++
	}
	for i, v := range s {
		if freq[v] == 1 {
			return i
		}
	}
	return -1
}

func firstUniq1(s string) int {
	freq := make([]int32, 26)
	for _, v := range s {
		freq[v-'a']++
	}
	for i, v := range s {
		if freq[v-'a'] == 1 {
			return i
		}
	}
	return -1
}

func Test_firstUniqChar(t *testing.T) {
	assert.Equal(t, firstUniqChar("loveleetcode"), 2)
	assert.Equal(t, firstUniq1("loveleetcode"), 2)
}
