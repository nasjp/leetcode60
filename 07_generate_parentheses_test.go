package main

import (
	"reflect"
	"testing"
)

// 22. Generate Parentheses
// https://leetcode.com/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	return helperGenerateParenthesis("", 0, 0, n, make([]string, 0))
}

func helperGenerateParenthesis(str string, opens, closes, n int, result []string) []string {
	if len(str) == n*2 {
		return append(result, str)
	}

	if opens < n {
		result = append(helperGenerateParenthesis(str+"(", opens+1, closes, n, result))
	}
	if closes < opens {
		result = append(helperGenerateParenthesis(str+")", opens, closes+1, n, result))
	}
	return result
}

func TestGenerateParenthesis(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"example1", args{n: 3}, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, but %v", tt.want, got)
			}
		})
	}
}
