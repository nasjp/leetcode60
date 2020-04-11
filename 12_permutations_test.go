package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// 46. Permutations
// https://leetcode.com/problems/permutations/
func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	head, tail := nums[0], nums[1:]
	res := make([][]int, 0)
	for _, v := range permute(tail) {
		for j := 0; j <= len(v); j++ {
			s := make([]int, len(v[:j]))
			copy(s, v[:j])
			s = append(append(s, head), v[j:]...)
			res = append(res, s)
		}
	}
	return res
}

func TestPermute(t *testing.T) {
	trans := cmp.Transformer("Sort", func(in [][]int) [][]int {
		out := append([][]int(nil), in...)
		sort.Slice(out, func(i, j int) bool {
			return out[i][0] < out[j][0]
		})
		return out
	})
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"example1", args{nums: []int{1, 2, 3}}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := permute(tt.args.nums); !cmp.Equal(got, tt.want, trans) {
				t.Errorf("want %v, but %v", tt.want, got)
			}
		})
	}
}
