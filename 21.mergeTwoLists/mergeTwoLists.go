package main

import "fmt"

/**
 * Definition for singly-linked list.
  * type ListNode struct {
	   *     Val int
	    *     Next *ListNode
		 * }
*/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	head = nil
	node := head
	for l1 != nil || l2 != nil {
		if l1 == nil {
			if head == nil {
				return l2
			} else {
				node.Next = l2
				return head
			}
		}
		if l2 == nil {
			if head == nil {
				return l1
			} else {
				node.Next = l1
				return head
			}
		}

		if l1.Val <= l2.Val {
			if head == nil {
				head = l1
				node = l1
				l1 = l1.Next
				continue
			}
			node.Next = l1
			node = node.Next
			l1 = l1.Next
		} else {
			if head == nil {
				head = l2
				node = l2
				l2 = l2.Next
				continue
			}
			node.Next = l2
			node = node.Next
			l2 = l2.Next
		}
	}
	return head

}
