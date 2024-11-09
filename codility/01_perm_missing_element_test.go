package codility

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/*
An array A consisting of N different integers is given. The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.
Your goal is to find that missing element.
Write a function:
func Solution(A []int) int
that, given an array A, returns the value of the missing element.

For example, given array A such that:

	A[0] = 2
	A[1] = 3
	A[2] = 1
	A[3] = 5

the function should return 4, as it is the missing element.
Write an efficient algorithm for the following assumptions:
N is an integer within the range [0..100,000];
the elements of A are all distinct;
each element of array A is an integer within the range [1..(N + 1)].
*/
func permMissingElement(arr []int) int {
	length := len(arr)
	expectedSum := (length + 1) * (length + 2) / 2
	var sum int
	for _, e := range arr {
		sum += e
	}
	return expectedSum - sum
}

func Test_PermMissingElement(t *testing.T) {
	arr := []int{1, 3, 4, 5, 6}
	assert.Equal(t, permMissingElement(arr), 2)
}

/*
1+2+...+n = n*(n+1)/2
=> if missing, sum = (n+1)*(n+2)/2
*/
