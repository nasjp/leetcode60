package main

// 141. Linked List Cycle
// https://leetcode.com/problems/linked-list-cycle/submissions/
type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	return resolveHasCycle(head)
}

func resolveHasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	fast, slow := head, head

	for slow != nil {
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true
		}
	}
	return false
}
