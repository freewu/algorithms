package main

// 面试题 02.04. Partition List LCCI
// Write code to partition a linked list around a value x, 
// such that all nodes less than x come before all nodes greater than or equal to x. 
// If x is contained within the list, the values of x only need to be after the elements less than x (see below). 
// The partition element x can appear anywhere in the "right partition"; 
// it does not need to appear between the left and right partitions.

// Example:
// Input: head = 3->5->8->5->10->2->1, x = 5
// Output: 3->1->2->10->5->5->8

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
func partition(head *ListNode, x int) *ListNode {
    beforeHead := &ListNode{Val: 0, Next: nil}
    before := beforeHead
    afterHead := &ListNode{Val: 0, Next: nil}
    after := afterHead
    for head != nil {
        if head.Val < x {
            before.Next = head
            before = before.Next
        } else {
            after.Next = head
            after = after.Next
        }
        head = head.Next
    }
    after.Next = nil
    before.Next = afterHead.Next
    return beforeHead.Next
}

func partition1(head *ListNode, x int) *ListNode {
    // 判断边界
    if head == nil || head.Next == nil {
        return head
    }
    // 之前的链表 初始化
    beforeHead := &ListNode{Val: 0, Next: nil}
    before := beforeHead
    // 之后的链表节点 初始化
    afterHead := &ListNode{Val: 0, Next: nil}
    after := afterHead
    for head != nil {
        if head.Val < x { // n1
            before.Next = head
            before = before.Next
        } else { // n2
            after.Next = head
            after = after.Next
        }
        head = head.Next
    }
    after.Next = nil
    before.Next = afterHead.Next
    return beforeHead.Next
} 

func main() {
    // Example 1:
    // Input: head = [1,4,3,2,5,2], x = 3
    // Output: [1,2,2,4,3,5]
    l1 := makeListNode([]int{1,4,3,2,5,2})
    printListNode(l1) // 1 -> 4 -> 3 -> 2 -> 5 -> 2
    printListNode(partition(l1, 3)) // 1 -> 2 -> 2 -> 4 -> 3 -> 5
    // Example 2:
    // Input: head = [2,1], x = 2
    // Output: [1,2]
    l2 := makeListNode([]int{2,1})
    printListNode(l2) // 2 -> 1
    printListNode(partition(l2, 2)) // 1 -> 2

    l11 := makeListNode([]int{1,4,3,2,5,2})
    printListNode(l11) // 1 -> 4 -> 3 -> 2 -> 5 -> 2
    printListNode(partition1(l11, 3)) // 1 -> 2 -> 2 -> 4 -> 3 -> 5
    l12 := makeListNode([]int{2,1})
    printListNode(l12) // 2 -> 1
    printListNode(partition1(l12, 2)) // 1 -> 2
}