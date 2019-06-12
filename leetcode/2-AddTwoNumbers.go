package main

import (
	"fmt"
)

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example 1
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

Eample 2
Input:[2,4,3][5,6,7]
Output: [7,0,1,1]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var t = &ListNode{-1, nil}
	var l3 = &ListNode{-1, t}

	var flag = 0 // 进位符
	for {
		// 如果循环到任意节点为空直接跳出
		if nil == l1 || nil == l2 {
			break
		}

		var s = l1.Val + l2.Val + flag
		if s >= 10 {
			// 如果和大于10 取模型进位
			t.Next = &ListNode{s % 10, nil}
			flag = 1
		} else {
			t.Next = &ListNode{s, nil}
			flag = 0
		}

		t = t.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	// 循环 l1 的剩余节点
	for {
		if nil == l1 {
			break
		}
		if flag == 1 {
			if (l1.Val + 1) >= 10 {
				flag = 1
				l1.Val = (l1.Val + 1) % 10
			} else {
				flag = 0
				l1.Val = l1.Val + 1
			}
		}
		t.Next = l1
		l1 = l1.Next
		t = t.Next
	}
	// 循环 l2 的剩余节点
	for {
		if nil == l2 {
			break
		}
		if flag == 1 {
			if (l2.Val + 1) >= 10 {
				flag = 1
				l2.Val = (l2.Val + 1) % 10
			} else {
				flag = 0
				l2.Val = l2.Val + 1
			}
		}
		t.Next = l2
		l2 = l2.Next
		t = t.Next
	}
	// 如果还存在进位
	if 1 == flag {
		t.Next = &ListNode{1, nil}
	}
	return l3.Next.Next
}

// best speed solution 这个方案好像上有问题的
func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	var head *ListNode
	var currentNode *ListNode
	// 都为空时跳出循环
	for (l1 != nil) || (l2 != nil) || (carry != 0) {
		l1_value := 0
		if l1 != nil {
			l1_value = l1.Val
			l1 = l1.Next
		}

		l2_value := 0
		if l2 != nil {
			l2_value = l2.Val
			l2 = l2.Next
		}

		fmt.Println(l1_value, l2_value, carry)
		current_value := carry + l1_value + l2_value
		carry = current_value / 10         // int
		current_value = current_value % 10 // 取余

		if head == nil {
			head = &ListNode{Val: current_value, Next: nil}
			currentNode = head
		} else {
			next := ListNode{Val: current_value, Next: nil}
			currentNode.Next = &next
			currentNode = &next
		}
	}

	return head
}

func printListNode(l *ListNode) {
	if nil == l {
		return
	}
	for {
		if nil == l.Next {
			fmt.Print(l.Val)
			break
		} else {
			fmt.Print(l.Val, " -> ")
		}
		l = l.Next
	}
	fmt.Println()
}

func main() {
	// var l11 = &ListNode{3, nil}
	// var l12 = &ListNode{4, l11}
	// var l13 = &ListNode{2, l12}

	var l21 = &ListNode{9, nil}
	var l22 = &ListNode{9, l21}
	var l23 = &ListNode{9, l22}

	var l13 = &ListNode{1, nil}

	printListNode(l23)
	printListNode(l13)
	printListNode(addTwoNumbers1(l23, l13))
	printListNode(addTwoNumbers(l23, l13))
}
