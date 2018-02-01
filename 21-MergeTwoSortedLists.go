package main

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
*/

import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if nil == l1 && nil == l2 {
		return nil
	}
	
	var t = &ListNode{-1,nil}
	var l3 = &ListNode{-1,t}

	for {
		if nil == l1 || nil == l2 {
			break
		}
		if l1.Val < l2.Val {
			t.Next = l1
			l1 = l1.Next
		} else {
			t.Next = l2
			l2 = l2.Next
		}
		t = t.Next
	}

	if nil == l1 {
		t.Next = l2
	} else {
		t.Next = l1
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
			fmt.Print(l.Val," -> ")
		}
		l = l.Next
	}
	fmt.Println("")
}

func main() {
	// Input: 1->2->4, 1->3->4
	//var l1 *ListNode;
	var l11 = &ListNode{4,nil}
	var l12 = &ListNode{2,l11}
	var l13 = &ListNode{1,l12}

	var l21 = &ListNode{4,nil}
	var l22 = &ListNode{3,l21}
	var l23 = &ListNode{1,l22}

	printListNode(l11)
	printListNode(l12)
	printListNode(l13)
	printListNode(l23)

	printListNode(mergeTwoLists(l13,l23))
}