package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func transfer2Num(l1 *ListNode) int {
	result := 0
	exp := 1
	for l1 != nil {
		result = result + exp*l1.Val
		exp = exp * 10
		l1 = l1.Next
	}
	return result
}

func transfer2Node(num int) *ListNode {
	var pre, head *ListNode
	pre = nil
	head = nil
	for true {
		remain := num % 10
		node := &ListNode{
			Val:  remain,
			Next: nil,
		}

		if pre == nil {
			pre = node
			head = node
		} else {
			pre.Next = node
			pre = node
		}
		num = (num - remain) / 10
		if num == 0 {
			return head
		}
	}
	return head
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := transfer2Num(l1)
	num2 := transfer2Num(l2)
	num := num1 + num2
	return transfer2Node(num)
}
func main() {
	fmt.Println("vim-go")
}
