package main

import (
	"fmt"
	"testing"
)

// time: O(n^2)
// space: O(n) -> for stack function
func fib(n int) int {
	if n <= 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// time: O(n)
// space: O(n)
func fibMemoization(n int, memo map[int]int) int {
	if val, ok := memo[n]; ok {
		return val
	}
	var result int
	if n <= 2 {
		result = 1
	} else {
		result = fibMemoization(n-1, memo) + fibMemoization(n-2, memo)
	}
	memo[n] = result
	return result
}

func Test_fib(t *testing.T) {
	memo := map[int]int{}
	fmt.Println("result: ", fibMemoization(3, memo))
}
