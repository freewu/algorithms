package main

import "fmt"

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

type ListNode struct {
	Val  int
	Next *ListNode
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

func makeNodeList(nums []int) *ListNode {
	var n = &ListNode{-1, nil}
	var b = &ListNode{-1, n}
	for i := 0; i < len(nums); i++ {
		n.Next = &ListNode{nums[i], nil}
		n = n.Next
	}
	return b.Next.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if nil == head {
		return nil
	}
	if nil == head.Next {
		return head
	}
	var t = head.Val
	var n = &ListNode{t, nil}
	var n1 = &ListNode{-1, n}

	for {
		if nil == head {
			break
		}

		if head.Val > t {
			t = head.Val
			n.Next = &ListNode{head.Val, nil}
			n = n.Next
		}

		head = head.Next
	}
	return n1.Next
}

func main() {
	var l21 = &ListNode{9, nil}
	var l22 = &ListNode{2, l21}
	var l23 = &ListNode{1, l22}

	printListNode(l23)
	printListNode(deleteDuplicates(l23))

	printListNode(deleteDuplicates(makeNodeList([]int{1, 2, 2, 3, 9, 9, 10, 11})))
}
