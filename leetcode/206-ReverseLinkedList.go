package main

// 206. Reverse Linked List
// Given the head of a singly linked list, reverse the list, and return the reversed list.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg" />
// Input: head = [1,2,3,4,5] 1->2->3->4->5->NULL
// Output: [5,4,3,2,1] 5->4->3->2->1->NULL

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg" />
// Input: head = [1,2]  1->2->NULL
// Output: [2,1] 2->1->NULL

// Example 3:
// Input: head = []
// Output: []
 
// Constraints:
//     The number of nodes in the list is the range [0, 5000].
//     -5000 <= Node.val <= 5000

// Example:
// Input: 
// Output: 

// 1->2->3->4->5

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 打印链表
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

// 数组创建链表
func makeListNode(arr []int) *ListNode {
    if (len(arr) == 0) {
        return nil
    }
    var l = (len(arr) - 1)
    var head = &ListNode{arr[l], nil}
    for i := l - 1; i >= 0; i--  {
        var n = &ListNode{arr[i], head}
        head = n
    }
    return head
}


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    var res *ListNode
    for head != nil {
        next := head.Next  // 备份head.Next
        head.Next = res // 更新  head.Next
        res = head      // 移动 new_head
        head = next
        //fmt.Println("new_head: ",new_head,"head: ",head)
    }
    return res
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

func main() {
    l1 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
	printListNode(l1)
    fmt.Println("reverseList(l1): ")
	printListNode(reverseList(l1))

    l2 := makeListNode([]int{1,2})
    fmt.Println("l2: ")
	printListNode(l2)
    fmt.Println("reverseList(l2): ")
	printListNode(reverseList(l2))

    l3 := makeListNode([]int{})
    fmt.Println("l3: ")
	printListNode(l3)
    fmt.Println("reverseList(l3): ")
	printListNode(reverseList(l3))


    l1 = makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
	printListNode(l1)
    fmt.Println("reverseList1(l1): ")
	printListNode(reverseList1(l1))

    l2 = makeListNode([]int{1,2})
    fmt.Println("l2: ")
	printListNode(l2)
    fmt.Println("reverseList1(l2): ")
	printListNode(reverseList1(l2))

    l3 = makeListNode([]int{})
    fmt.Println("l3: ")
	printListNode(l3)
    fmt.Println("reverseList1(l3): ")
	printListNode(reverseList1(l3))
}
