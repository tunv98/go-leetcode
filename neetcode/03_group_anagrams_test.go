package main

import (
	"fmt"
	"sort"
	"testing"
)

/*
Given an array of strings strs, group all anagrams together into sublist => You may return the output in any order
*/

/*
Time complexity:
Sort: O(nlogn)
Iterating over m strings: O(m)
=> Total: O(m*nlogn)
*/
func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string) // use hashmap
	for _, str := range strs {
		runes := []rune(str)
		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})
		sortedStr := string(runes)
		anagramMap[sortedStr] = append(anagramMap[sortedStr], str)
	}
	result := make([][]string, 0, len(anagramMap))
	for _, gr := range anagramMap {
		result = append(result, gr)
	}
	return result
}

func groupAnagramsOptimized(strs []string) [][]string {
	anagramMap := make(map[[26]int][]string)
	for _, str := range strs {
		var count [26]int
		for _, c := range str {
			count[c-'a']++
		}
		anagramMap[count] = append(anagramMap[count], str)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, anagramSlice := range anagramMap {
		result = append(result, anagramSlice)
	}
	return result
}
func Test_GroupAnagrams(t *testing.T) {
	fmt.Println(groupAnagrams([]string{"act", "pots", "tops", "cat", "stop", "hat"}))
	fmt.Println(groupAnagramsOptimized([]string{"act", "pots", "tops", "cat", "stop", "hat"}))
}

// Strings are immutable in Go → You cannot sort str directly because str is a string.
// Fix: Convert str to a []rune (or []byte), sort it, then convert it back.
