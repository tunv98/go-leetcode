package main

import (
	"reflect"
	"testing"
)

func TestEasy(t *testing.T) {
	// 1. Viết hàm đảo ngược một chuỗi input. Ví dụ "abcdef" => "fedcba"
	if reserve("abcdef") != "fedcba" {
		t.Error("test reserve error")
	}
	// 2. Viết hàm bỏ đi những phần tử lặp lại trong mảng. Ví dụ [1,2,3,4,5,4,4,2] => [1,2,3,4,5]
	if !reflect.DeepEqual(removeDuplicate([]int{1, 2, 3, 4, 5, 4, 4, 2}), []int{1, 2, 3, 4, 5}) {
		t.Errorf("test removeDuplicate error")
	}
	//3. Viết hàm kiểm tra cặp kí tự '(' và ')' trong chuỗi. Ví dụ
	//"()()()()(())" => output: true
	//"(()()()(()" => output: false
	if !checkParentheses("()()()()(())") || checkParentheses("(()()()(()") {
		t.Errorf("test checkParentheses error")
	}
}

func reserve(s string) string {
	r := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func removeDuplicate(arr []int) (uniqueArr []int) {
	encountered := make(map[int]bool)
	for _, i := range arr {
		if encountered[i] {
			continue
		}
		uniqueArr = append(uniqueArr, i)
		encountered[i] = true
	}
	return uniqueArr
}

func checkParentheses(s string) bool {
	stack := make([]rune, 0)
	for _, r := range s {
		if r == '(' {
			stack = append(stack, r)
			continue
		}
		if r == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
