package main

// 86. Partition List
// Given the head of a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
// You should preserve the original relative order of the nodes in each of the two partitions.

// Example 1:
// Input: head = [1,4,3,2,5,2], x = 3
// Output: [1,2,2,4,3,5]

// Example 2:
// Input: head = [2,1], x = 2
// Output: [1,2]
 
// Constraints:
//     The number of nodes in the list is in the range [0, 200].
//     -100 <= Node.val <= 100
//     -200 <= x <= 200

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