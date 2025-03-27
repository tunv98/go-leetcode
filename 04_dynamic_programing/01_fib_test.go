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

// time: O(n)
// space: O(n)
func fibTabulation(n int) int {
	table := make([]int, n+1)
	table[1] = 1
	for i := 2; i <= n; i++ {
		table[i] = table[i-1] + table[i-2]
	}
	return table[n]
}

// time: O(n)
// space: O(1)
func fibOptimizedSpace(n int) int {
	prev1, prev2 := 1, 0
	for i := 2; i <= n; i++ {
		cur := prev1 + prev2
		prev1, prev2 = cur, prev1
	}
	return prev1
}
func Test_fib(t *testing.T) {
	fmt.Println("result: ", fibMemoization(6, map[int]int{}))
	fmt.Println("result: ", fibTabulation(6))
	fmt.Println("result: ", fibOptimizedSpace(6))
}
