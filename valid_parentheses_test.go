package main

// 20. Valid Parentheses
// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}
	if s == "" {
		return true
	}
	list := map[rune]rune{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	stack := make([]rune, 0)
	for _, c := range s {
		if clo, ok := list[c]; ok {
			stack = append(stack, clo)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		pop := stack[len(stack)-1]
		if c != pop {
			return false
		}
		stack = stack[:len(stack)-1]
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
