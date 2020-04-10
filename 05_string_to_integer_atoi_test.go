package main

import (
	"math"
	"testing"
)

// 8. String to Integer (atoi)
// https://leetcode.com/problems/string-to-integer-atoi/
func myAtoi(str string) int {
	var (
		ans   int64
		start bool
		sign  int64 = 1
	)

LOOP:
	for _, c := range str {
		switch {
		case '0' <= c && c <= '9':
			start = true
			ans = ans*10 + int64(c-'0')
			if ans > math.MaxInt32 {
				break LOOP
			}
		case !start && c == ' ':
		case !start && c == '+':
			start = true
		case !start && c == '-':
			sign = -1
			start = true
		default:
			break LOOP
		}
	}

	ans = sign * ans
	if ans > math.MaxInt32 {
		ans = math.MaxInt32
	}
	if ans < math.MinInt32 {
		ans = math.MinInt32
	}
	return int(ans)
}

func TestMyAtoi(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{str: "42"}, 42},
		{"example2", args{str: "   -42"}, -42},
		{"example3", args{str: "4193 with words"}, 4193},
		{"example4", args{str: "words and 987"}, 0},
		{"example5", args{str: "-91283472332"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.str); got != tt.want {
				t.Errorf("want %d, but %d", tt.want, got)
			}
		})
	}
}
