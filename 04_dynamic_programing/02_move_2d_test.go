package main

import (
	"fmt"
	"testing"
)

/*
Say that you area traveler on a 2D grid.
You begin in the top-left corner and your goal is to travel to the bottom-right corner
You may only move down or right

How many can you travel to the goal on a grid with dimensions m*n
*/

/*
Matrix(2,3)
2,3 | 2,2 | 2,1
1,3 | 1,2 | 1,1
*/
/*
Time: O(2^n+m) // tăng theo cấp số nhân => tổ hợp
Space: O(n+m) // stack
*/

func gridTraveler(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	return gridTraveler(m-1, n) + gridTraveler(m, n-1)
}

/*
Time: O(nxm)
Space: O(n+m)
*/
func gridTravelerMemoization(m, n int, memo map[string]int) int {
	key := fmt.Sprintf("%d-%d", m, n)
	if val, ok := memo[key]; ok {
		return val
	}
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	memo[key] = gridTravelerMemoization(m-1, n, memo) + gridTravelerMemoization(m, n-1, memo)
	return memo[key]
}

func Test_GridTraveler(t *testing.T) {
	fmt.Println("Result: ", gridTraveler(2, 3))
	memo := map[string]int{}
	fmt.Println("Result by memoization:", gridTravelerMemoization(2, 3, memo))
}
