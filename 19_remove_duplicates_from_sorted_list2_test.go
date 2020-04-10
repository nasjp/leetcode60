package main

// 82. Remove Duplicates from Sorted List II
// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	current := head
	pre := head
	var isFirstDup bool
	max := current.Val
	for current.Next != nil {
		if max == current.Next.Val {
			if max == head.Val {
				isFirstDup = true
			}
			pre.Next = current.Next.Next
			current = pre
			continue
		}
		max = current.Next.Val
		pre = current
		current = current.Next
	}
	if isFirstDup {
		return head.Next
	}
	return head
}
