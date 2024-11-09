package main

import (
	"github.com/stretchr/testify/assert"
	"leetcode/utils"
	"math"
	"testing"
)

//https://leetcode.com/problems/fruit-into-baskets/description/
//Input: fruits = [0,1,2,2]
//Output: 3
//Explanation: We can pick from trees [1,2,2].
//If we had started at the first tree, we would only pick from trees [0,1]
//=> Find max number of fruits in each basket

func TestTotalFruit(t *testing.T) {
	//fruits := []int{1, 2, 3, 2, 2}
	//assert.Equal(t, 4, totalFruit(fruits))

	fruits := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	assert.Equal(t, 5, totalFruit(fruits))

}

func totalFruit(fruits []int) int {
	start := 0
	max := math.MinInt
	charCount := make(map[int]int)
	for end := 0; end < len(fruits); end++ {
		fEnd := fruits[end]
		charCount[fEnd]++
		for len(charCount) > 2 {
			fStart := fruits[start]
			charCount[fStart]--
			start++
			if charCount[fStart] == 0 {
				delete(charCount, fStart)
			}
		}
		max = utils.Max([]int{max, sumValueOfMap(charCount)})
	}
	return max
}

func sumValueOfMap(counter map[int]int) int {
	sum := 0
	for _, value := range counter {
		sum += value
	}
	return sum
}
