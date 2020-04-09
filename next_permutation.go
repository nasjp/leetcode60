package main

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
