package main

import (
	"math"
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
