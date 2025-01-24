package random

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
Input: s = ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
*/
/*
time: O(n)
space: O(1)
*/
func reverseString(s []byte) {
	iLeft, iRight := 0, len(s)-1
	for iLeft < iRight {
		s[iLeft], s[iRight] = s[iRight], s[iLeft]
		iLeft++
		iRight--
	}
}

func reverseString1(s []byte) {
	stack := make([]byte, 0, len(s))
	for _, v := range s {
		stack = append(stack, v)
	}
	for i := range s {
		s[i] = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}

func Test_reserveString(t *testing.T) {
	data := []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(data)
	assert.Equal(t, data, []byte{'o', 'l', 'l', 'e', 'h'})

	reverseString1(data)
	assert.Equal(t, data, []byte{'h', 'e', 'l', 'l', 'o'})
}
