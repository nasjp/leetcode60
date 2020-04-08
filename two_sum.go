package main

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
