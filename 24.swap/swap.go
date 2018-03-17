package main

import "fmt"

/* Definition for singly-linked list.a*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	node := head
	if node == nil {
		return head
	}
	if node.Next == nil {
		return head
	}
	head = node.Next
	n := node.Next.Next
	node.Next.Next = node
	node.Next = n
	pre := node
	node = n

	for node != nil && node.Next != nil {
		pre.Next = node.Next
		n := node.Next.Next
		node.Next.Next = node
		node.Next = n
		pre = node
		node = n
	}
	return head
}

func main() {
	fmt.Println("vim-go")
}
