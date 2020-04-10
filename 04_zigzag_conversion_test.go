package main

import "testing"

// 6. ZigZag Conversion
// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	matrix := make([][]byte, numRows)
	var row int
	step := 1
	for i := 0; i < len(s); {
		matrix[row] = append(matrix[row], s[i])
		if (step == 1 && row == numRows-1) || (step == -1 && row == 0) {
			step = -step
		}
		row += step
		i++
	}
	var ret string
	for _, arr := range matrix {
		ret += string(arr)
	}
	return ret
}

func TestConvert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"example1", args{s: "PAYPALISHIRING", numRows: 3}, "PAHNAPLSIIGYIR"},
		{"example2", args{s: "PAYPALISHIRING", numRows: 4}, "PINALSIGYAHRPI"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("want %s, but %s", tt.want, got)
			}
		})
	}
}
