package main

import (
	"reflect"
	"testing"
)

// 1. Two Sum
// https://leetcode.com/problems/two-sum/submissions/
func twoSum(nums []int, target int) []int {
	list := map[int]int{}
	for i, v := range nums {
		list[v] = i
	}
	for ix, v := range nums {
		if iy, ok := list[target-v]; ok && ix != iy {
			return []int{ix, iy}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{nums: []int{2, 7, 11, 15}, target: 9}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("want %v, but %v", tt.want, got)
			}
		})
	}
}
