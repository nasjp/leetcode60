package main

// 2. Add Two Numbers
// https://leetcode.com/problems/add-two-numbers/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	root := &ListNode{}
	cur := root
	var beforeIncre int
	for l1 != nil || l2 != nil {
		val, incre := sum(l1, l2)
		curVal := val + beforeIncre
		curIncre := curVal / 10
		curVal = curVal - curIncre*10
		cur.Next = &ListNode{Val: curVal}
		beforeIncre = incre + curIncre
		cur = cur.Next
		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}
		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}
	if beforeIncre != 0 {
		cur.Next = &ListNode{Val: beforeIncre}
	}
	return root.Next
}

func sum(l1 *ListNode, l2 *ListNode) (int, int) {
	var val1, val2 int
	if l1 != nil {
		val1 = l1.Val
	}
	if l2 != nil {
		val2 = l2.Val
	}
	sum := val1 + val2
	increments := sum / 10
	val := sum - increments*10
	return val, increments
}
