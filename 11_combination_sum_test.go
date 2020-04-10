package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := make([][]int, 0)
	helperCombinationSum(candidates, target, 0, []int{}, &res)
	return res
}

func helperCombinationSum(nums []int, target int, pos int, ans []int, res *[][]int) {
	if target == 0 {
		*res = append(*res, append([]int(nil), ans...))
		return
	}
	for i := pos; i < len(nums); i++ {
		if target-nums[i] >= 0 {
			helperCombinationSum(nums, target-nums[i], i, append(ans, nums[i]), res)
		}
	}
}

func TestCombinationSum(t *testing.T) {
	trans := cmp.Transformer("Sort", func(in [][]int) [][]int {
		out := append([][]int(nil), in...)
		sort.Slice(out, func(i, j int) bool {
			return len(out[i]) < len(out[j])
		})
		return out
	})
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"example1", args{candidates: []int{2, 3, 6, 7}, target: 7}, [][]int{{7}, {2, 2, 3}}},
		{"example2", args{candidates: []int{2, 3, 5}, target: 8}, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum(tt.args.candidates, tt.args.target); !cmp.Equal(got, tt.want, trans) {
				t.Errorf("want %v, but %v", tt.want, got)
			}
		})
	}
}
