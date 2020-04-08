package main

// 6. ZigZag Conversion
// https://leetcode.com/problems/zigzag-conversion/
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	matrix := make([][]rune, numRows)
	var row, col int
	incre := numRows - 2
	for _, c := range s {
		switch col % (numRows - 1) {
		case 0:
			matrix[row] = append(matrix[row], c)
			if row+1 == numRows {
				col++
				if incre == numRows-1 {
					incre = 0
				}
				row = incre
				continue
			}
			row++
		default:
			for i := 0; i < numRows; i++ {
				if i != row {
					continue
				}
				matrix[i] = append(matrix[i], c)
			}
			col++
			row--
			if row < 0 {
				row = 0
			}
		}
	}
	var ret string
	for _, arr := range matrix {
		ret += string(arr)
	}
	return ret
}
