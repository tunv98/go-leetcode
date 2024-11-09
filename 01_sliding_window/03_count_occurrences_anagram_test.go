package main

import (
	"fmt"
	"testing"
)

//Given a word and a text, return the count of occurrences of the anagrams of the word in the text.
//Input: text = gotxxotgxdogt, word = got
//Output: 3 (got, otg, ogt)

func TestCountAnagram(t *testing.T) {
	text := "gotxxotgxdogt"
	word := "got"
	fmt.Println(countAnagram(text, word))
}

func countAnagram(text, word string) int {
	count := 0
	comparedWord := ""
	for r := 0; r < len(text); r++ {
		comparedWord += string(text[r])
		if len(comparedWord) == len(word) {
			if isAnagram(comparedWord, word) {
				fmt.Println(comparedWord)
				count++
			}
			comparedWord = comparedWord[1:]
		}
	}
	return count
}

func isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	code := make([]int, 26)
	for _, c := range a {
		code[c-'a'] = 1
	}
	for _, c := range b {
		if code[c-'a'] == 0 {
			return false
		}
	}
	return true
}
