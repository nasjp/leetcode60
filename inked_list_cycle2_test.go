package main

// 142. Linked List Cycle II
// https://leetcode.com/problems/linked-list-cycle-ii/
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var (
		slow = head
		fast = head
	)

	for slow != nil {
		if slow.Next == nil || fast.Next == nil || fast.Next.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for {
				if head == slow {
					return head
				}
				head = head.Next
				slow = slow.Next
			}
		}
	}
	return nil
}
