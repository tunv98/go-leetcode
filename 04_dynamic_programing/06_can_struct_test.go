package main

import (
	"fmt"
	"strings"
	"testing"
)

/*
the function should return a boolean indicating whether the `target` can be constructed by
concatenating elements of the `wordbank` array
you may reuse elements of `wordbank` as many times as needed

canStruct(abcdef,[ab,abc,cd,def,abcd] = abc+def => true
canStruct('', [cat,dog,more]) -> true

abcdef
=> -ab -> cdef - cd -> ef (X)
=> -abc -> def - def -> '' (O)
=> -abcd -> ef (X)
*/

func canConstruct(target string, wordBank []string) bool {
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if !strings.HasPrefix(target, word) {
			continue
		}
		remainder := strings.TrimPrefix(target, word)
		remainderVal := canConstruct(remainder, wordBank)
		if remainderVal {
			return true
		}
	}
	return false
}

func canConstructMemoization(target string, wordBank []string, memo map[string]bool) bool {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if !strings.HasPrefix(target, word) {
			continue
		}
		remainder := strings.TrimPrefix(target, word)
		remainderVal := canConstructMemoization(remainder, wordBank, memo)
		if remainderVal {
			memo[target] = true
			return true
		}
	}
	memo[target] = false
	return false
}

func TestCanConstruct(t *testing.T) {
	//fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "cdef", "abcd"}))
	fmt.Println(canConstructMemoization("abcdef", []string{"ab", "abc", "cd", "cdef", "abcd"}, map[string]bool{}))
}
