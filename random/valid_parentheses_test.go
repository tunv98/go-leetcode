package random

import (
	"fmt"
	"testing"
)

func isValid(s string) bool {
	var stack []rune
	matching := map[rune]rune{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	for _, c := range s {
		if closeBracket, existed := matching[c]; existed { // if c is open
			stack = append(stack, closeBracket)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func Test_isValid(t *testing.T) {
	//fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("([])[]{}")) // true
	//fmt.Println(isValid("(]"))       // false
	//fmt.Println(isValid("([)]"))     // false
	//fmt.Println(isValid("{[]}"))     // true
}

/*


 */
