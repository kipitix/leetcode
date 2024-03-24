package p206

import "testing"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev *ListNode
	cur := head
	next := head.Next
	for next != nil {
		cur.Next = prev
		prev = cur
		cur = next
		next = next.Next
	}
	cur.Next = prev

	return cur
}

func TestReverseList(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = reverseList(head)
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}
