package main

import "testing"

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

func TestIsValid(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"example1", args{s: "()"}, true},
		{"example2", args{s: "()[]{}"}, true},
		{"example3", args{s: "(]"}, false},
		{"example4", args{s: "([)]"}, false},
		{"example5", args{s: "{[]}"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.s); got != tt.want {
				t.Errorf("want %t, but %t", tt.want, got)
			}
		})
	}
}
