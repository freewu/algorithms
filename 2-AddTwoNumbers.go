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
		if nil == l1 || nil == l2 {
			break
		}

		var s = l1.Val + l2.Val + flag
		if s >= 10 {
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
	if 1 == flag {
		t.Next = &ListNode{1, nil}
	}
	return l3.Next.Next
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
	fmt.Println("")
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
	printListNode(addTwoNumbers(l23, l13))
}
