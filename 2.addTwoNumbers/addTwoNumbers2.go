package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNodes(l1 *ListNode, l2 *ListNode, carry int) (*ListNode, int) {
	if l1 == nil && l2 == nil {
		return nil, carry
	}

	if l1 == nil {
		result := l2.Val + carry
		remain := result % 10
		nextCarry := (result - remain) / 10
		return &ListNode{
			Val:  remain,
			Next: nil,
		}, nextCarry
	}

	if l2 == nil {
		result := l1.Val + carry
		remain := result % 10
		nextCarry := (result - remain) / 10
		return &ListNode{
			Val:  remain,
			Next: nil,
		}, nextCarry
	}

	result := l1.Val + l2.Val + carry
	remain := result % 10
	nextCarry := (result - remain) / 10
	return &ListNode{
		Val:  remain,
		Next: nil,
	}, nextCarry
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var head, pre *ListNode
	head = nil
	pre = nil
	carry = 0

	for l1 != nil || l2 != nil {
		var node *ListNode
		node, carry = addTwoNodes(l1, l2, carry)
		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
		if head == nil {
			head = node
		}
		if pre == nil {
			pre = node
			continue
		}
		pre.Next = node
		pre = node

	}
	if carry != 0 {
		pre.Next = &ListNode{
			Val:  carry,
			Next: nil,
		}
	}
	return head
}

func main() {
	fmt.Println("vim-go")
}
