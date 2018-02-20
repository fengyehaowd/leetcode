package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNodes(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil {
		if carry == 0 {
			return nil
		} else {
			return &ListNode{
				Val:  carry,
				Next: nil,
			}
		}
	}

	if l1 == nil {
		result := l2.Val + carry
		remain := result % 10
		nextCarry := (result - remain) / 10
		node := addTwoNodes(l1, l2.Next, nextCarry)
		return &ListNode{
			Val:  remain,
			Next: node,
		}
	}

	if l2 == nil {
		result := l1.Val + carry
		remain := result % 10
		nextCarry := (result - remain) / 10
		node := addTwoNodes(l1.Next, l2, nextCarry)
		return &ListNode{
			Val:  remain,
			Next: node,
		}
	}

	result := l1.Val + l2.Val + carry
	remain := result % 10
	nextCarry := (result - remain) / 10
	node := addTwoNodes(l1.Next, l2.Next, nextCarry)
	return &ListNode{
		Val:  remain,
		Next: node,
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNodes(l1, l2, 0)
}
func main() {
	fmt.Println("vim-go")
}
