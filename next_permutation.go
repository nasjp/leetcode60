package main

// 31. Next Permutation
// https://leetcode.com/problems/next-permutation/
func nextPermutation(nums []int) {
	var (
		i = len(nums) - 2
		j int
	)
	for i >= 0 {
		if nums[i] < nums[i+1] {
			j = i + 1
			for j < len(nums) && nums[j] > nums[i] {
				j++
			}
			j--
			nums[i], nums[j] = nums[j], nums[i]
			break
		}
		i--
	}
	i++
	for j = len(nums) - 1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
