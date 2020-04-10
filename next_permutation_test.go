package main

import (
	"reflect"
	"testing"
)

// 31. Next Permutation
// https://leetcode.com/problems/next-permutation/
func nextPermutation(nums []int) {
	var k, l int
	for k = len(nums) - 2; k >= 0; {
		if nums[k] < nums[k+1] {
			l = k + 1
			for l < len(nums) && nums[l] > nums[k] {
				l++
			}
			l--
			nums[k], nums[l] = nums[l], nums[k]
			break
		}
		k--
	}
	k++
	for l = len(nums) - 1; k < l; k, l = k+1, l-1 {
		nums[k], nums[l] = nums[l], nums[k]
	}
}

func TestNextPermutation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"example1", args{nums: []int{1, 2, 3}}, []int{1, 3, 2}},
		{"example2", args{nums: []int{3, 2, 1}}, []int{1, 2, 3}},
		{"example3", args{nums: []int{1, 1, 5}}, []int{1, 5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if nextPermutation(tt.args.nums); !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("want %v, but %v", tt.want, tt.args.nums)
			}
		})
	}
}
