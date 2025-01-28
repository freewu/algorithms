package main

// 面试题 02.01. Remove Duplicate Node LCCI
// Write code to remove duplicates from an unsorted linked list.

// Example1:
// Input: [1, 2, 3, 3, 2, 1]
// Output: [1, 2, 3]

// Example2:
// Input: [1, 1, 1, 1, 2]
// Output: [1, 2]

// Note:
//     The length of the list is within the range[0, 20000].
//     The values of the list elements are within the range [0, 20000].

// Follow Up:
//     How would you solve this problem if a temporary buffer is not allowed?

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
func removeDuplicateNodes(head *ListNode) *ListNode {
    ob := head
    for ob != nil {
        oc := ob
        for oc.Next != nil {
            if oc.Next.Val == ob.Val {
                oc.Next = oc.Next.Next
            } else {
                oc = oc.Next
            }
        }
        ob = ob.Next
    }
    return head
}

func removeDuplicateNodes1(head *ListNode) *ListNode {
    if head == nil { return head }
    mp := map[int]bool{ head.Val: true }
    pos := head
    for pos.Next != nil {
        cur := pos.Next
        if !mp[cur.Val] {
            mp[cur.Val], pos  = true, pos.Next
        } else {
            pos.Next = pos.Next.Next
        }
    }
    pos.Next = nil
    return head
}

func main() {
    // Example1:
    // Input: [1, 2, 3, 3, 2, 1]
    // Output: [1, 2, 3]
    list1 := makeListNode([]int{1, 2, 3, 3, 2, 1}) 
    printListNode(list1) // 1 -> 2 -> 3 -> 3 -> 2 -> 1
    printListNode(removeDuplicateNodes(list1)) // 1 -> 2 -> 3
    // Example2:
    // Input: [1, 1, 1, 1, 2]
    // Output: [1, 2]
    list2 := makeListNode([]int{1, 1, 1, 1, 2})
    printListNode(list2) // 1 -> 1 -> 1 -> 1 -> 2
    printListNode(removeDuplicateNodes(list2)) // 1 -> 2

    list11 := makeListNode([]int{1, 2, 3, 3, 2, 1}) 
    printListNode(list11) // 1 -> 2 -> 3 -> 3 -> 2 -> 1
    printListNode(removeDuplicateNodes1(list11)) // 1 -> 2 -> 3
    list12 := makeListNode([]int{1, 1, 1, 1, 2})
    printListNode(list12) // 1 -> 1 -> 1 -> 1 -> 2
    printListNode(removeDuplicateNodes1(list12)) // 1 -> 2
}