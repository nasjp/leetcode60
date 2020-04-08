package main

// 3. Longest Substring Without Repeating Characters
// https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	var max int
	for i, j, list := 0, 0, map[byte]int{}; i < len(s); {
		if prei, ok := list[s[i]]; ok && prei >= j {
			j = prei + 1
		}
		list[s[i]] = i
		if cur := i - j + 1; cur > max {
			max = cur
		}
		i++
	}
	return max
}
