package main

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
