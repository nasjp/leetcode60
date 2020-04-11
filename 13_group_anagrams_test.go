package main

import (
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

// 49. Group Anagrams
// https://leetcode.com/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, s := range strs {
		var table [26]int
		for _, char := range s {
			table[char-'a']++
		}
		m[table] = append(m[table], s)
	}
	var res [][]string
	for _, val := range m {
		res = append(res, val)
	}
	return res
}

func TestGroupAnagrams(t *testing.T) {
	transMat := cmp.Transformer("Sort", func(in [][]string) [][]string {
		out := append([][]string(nil), in...)
		sort.Slice(out, func(i, j int) bool {
			return len(out[i]) < len(out[j])
		})
		return out
	})

	transSli := cmp.Transformer("Sort", func(in []string) []string {
		out := append([]string(nil), in...)
		sort.Slice(out, func(i, j int) bool {
			return out[i] < out[j]
		})
		return out
	})
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"example1", args{strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"}}, [][]string{
			{"ate", "eat", "tea"},
			{"nat", "tan"},
			{"bat"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !cmp.Equal(got, tt.want, transMat, transSli) {
				t.Errorf("want %v, but %v", tt.want, got)
			}
		})
	}
}
