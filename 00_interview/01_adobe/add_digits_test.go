package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Given a non-negative integer num, repeatedly add all its digits until the result has only one digit.

//For example:
//Given num = 38, the process is like: 3 + 8 = 11, 1 + 1 = 2. Since 2 has only one digit, return it.

//Follow up:
//Could you do it without any loop/recursion in O(1) runtime?

func addDigits(num int) int {
	for num >= 10 {
		sum := 0
		for num > 0 {
			sum += num % 10 // put end
			num /= 10       // remove end
		}
		num = sum
	}
	return num
}

func addDigitsOptimize(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}

/*
Tính chất digital root
num % 9 = 0 -> return 9 => bội của 9 thì tổng của nó luôn bằng 9
khi cộng tất cả các chữ số của 1 số => kết quả của tổng đó mod 9 = số ban đầu mod 9
38 => 3 + 8 = 11 => 1 + 1 = 2
38 % 9 = 2
*/

func Test_addDigits(t *testing.T) {
	//assert.Equal(t, 1, addDigits(19))
	assert.Equal(t, 2, addDigitsOptimize(1910))
}
