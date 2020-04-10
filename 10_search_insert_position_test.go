package main

import "testing"

// 35. Search Insert Position
// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	var ret int
	var insert bool
	for i, v := range nums {
		if target == v {
			ret = i
			insert = true
			break
		}
		if target < v {
			ret = i
			insert = true
			break
		}
	}
	if !insert {
		ret = len(nums)
	}
	if ret < 0 {
		ret = 0
	}
	return ret
}

func TestSerchIntert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example1", args{nums: []int{1, 3, 5, 6}, target: 5}, 2},
		{"example2", args{nums: []int{1, 3, 5, 6}, target: 2}, 1},
		{"example3", args{nums: []int{1, 3, 5, 6}, target: 7}, 4},
		{"example4", args{nums: []int{1, 3, 5, 6}, target: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("want %d, but %d", tt.want, got)
			}
		})
	}
}
