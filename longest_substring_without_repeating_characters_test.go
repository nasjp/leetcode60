package main

import (
	"testing"
)

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	var max int
	for i, j, list := 0, 0, map[byte]int{}; i < len(s); {
		if prei, ok := list[s[i]]; ok && prei >= j {
			j = prei + 1
		}
		list[s[i]] = i
		if cur := i - j + 1; cur > max {
			max = cur
		}
		i++
	}
	return max
}

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{s: "abcabcbb"}, 3},
		{"example2", args{s: "bbbbb"}, 1},
		{"example3", args{s: "pwwkew"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("want %d, but %d", tt.want, got)
			}
		})
	}
}
