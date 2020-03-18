package main

import "fmt"

/**
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL


1->2->3->4->5


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
func reverseList(head *ListNode) *ListNode {
	var new_head *ListNode
	for {
        if head == nil {
            break
        }
	    next := head.Next  // 备份head.Next
	    head.Next = new_head // 更新  head.Next
	    new_head = head      // 移动 new_head
	    head = next
	    fmt.Println("new_head: ",new_head,"head: ",head)
    }
	return new_head
}

// best
func reverseList1(head *ListNode) *ListNode {
    var prev *ListNode = nil
    cur := head
    var next *ListNode = nil

    for cur != nil {
        next = cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }

    return prev
}

func printListNode(a *ListNode) {
	head := a
	for {
		fmt.Println(head.Val)
		head = head.Next
		if nil == head {
			break
		}
	}
}

func main() {
	var a, b, c, d, e ListNode
	a.Val = 1
	b.Val = 2
	c.Val = 3
	d.Val = 4
	e.Val = 5
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &e
	e.Next = nil

	printListNode(&a)
	printListNode(reverseList(&a))
}
