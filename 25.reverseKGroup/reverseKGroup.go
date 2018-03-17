package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	node := head

	if k == 1 {
		return head
	}

	//对于开头进行处理
	var reserve []*ListNode

	for i := 1; i <= k; i++ {
		node = node.Next
		reserve = append(reserve, node)
		if node == nil && i != k {
			return head
		}
	}
	for i := k - 2; i > 0; i-- {
		reserve[i].Next = reserve[i-1]
	}
	reserve[0].Next = head
	head.Next = reserve[k-1]
	pre := head
	head = reserve[k-2]
	//同样的逻辑，对于剩下的段进行处理

	for {
		reserve = reserve[0:0]
		node = pre
		for i := 0; i < k; i++ {
			node = node.Next
			if node == nil {
				return head
			}
			reserve = append(reserve, node)
		}
		pre.Next = node
		next := node.Next

		for i := k - 1; i > 0; i-- {
			reserve[i].Next = reserve[i-1]
		}
		reserve[0].Next = next
		pre = reserve[0]
	}

	return head

}

func main() {
	fmt.Println("vim-go")
}
