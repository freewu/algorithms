package main

// 61. Rotate List
// Given the head of a linked list, rotate the list to the right by k places.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/13/rotate1.jpg" />
// Input: head = [1,2,3,4,5], k = 2
// Output: [4,5,1,2,3]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/13/roate2.jpg" />
// Input: head = [0,1,2], k = 4
// Output: [2,0,1]
 
// Constraints:
//     The number of nodes in the list is in the range [0, 500].
//     -100 <= Node.val <= 100
//     0 <= k <= 2 * 10^9

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
func rotateRight(head *ListNode, k int) *ListNode {
    if k == 0 || head == nil{
        return head
    }
    var dummy  = &ListNode{Next: head}
    count, tail, sp := 0, head, head
    
    for tail.Next != nil {
        count++
        tail = tail.Next
    }
    count++
    j := k % count
    for i := 1; i < count-j; i++{
        sp = sp.Next
    }
    tail.Next = head
    dummy.Next = sp.Next
    sp.Next = nil
    return dummy.Next
}

func main() {
    l11 := makeListNode([]int{1,2,3,4,5})
    printListNode(l11)
    printListNode(rotateRight(l11,2)) //  [4,5,1,2,3]

    l12 := makeListNode([]int{0,1,2})
    printListNode(l12)
    printListNode(rotateRight(l12,4)) // [2,0,1]
 
}