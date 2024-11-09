package main

import "testing"

/*
https://leetcode.com/problems/permutation-in-string
// permutation: hoán vị
*/

func checkInclusion(s1 string, s2 string) bool {
	l := 0
	for r := 0; r < len(s2); r++ {
		if len(s1) > r+1 {
			continue
		}
		if checkPermutation(s1, s2[l:r+1]) {
			return true
		}
		l++
	}
	return false
}

// s1 and s2 consist of lowercase English letters
func checkPermutation(a, b string) bool {
	if len(a) != len(b) {
		return false
	}
	code := make([]int, 26)
	for _, c := range a {
		code[c-'a']++
	}

	for _, c := range b {
		if code[c-'a'] == 0 {
			return false
		}
		code[c-'a']--
	}
	return true
}

func TestCheckInclusion(t *testing.T) {
	s1 := "ab"
	s2 := "eidbaooo"
	if checkInclusion(s1, s2) != true {
		t.Errorf("Expected true")
	}
	s1 = "ab"
	s2 = "eidboaoo"
	if checkInclusion(s1, s2) != false {
		t.Errorf("Expected false")
	}

	s1 = "hello"
	s2 = "ooolleoooleh"
	if checkInclusion(s1, s2) != false {
		t.Errorf("Expected false")
	}
}
