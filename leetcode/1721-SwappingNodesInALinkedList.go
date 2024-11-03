package main

// 1721. Swapping Nodes in a Linked List
// You are given the head of a linked list, and an integer k.

// Return the head of the linked list after swapping the values of the kth node from the beginning 
// and the kth node from the end (the list is 1-indexed).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/21/linked1.jpg" />
// Input: head = [1,2,3,4,5], k = 2
// Output: [1,4,3,2,5]

// Example 2:
// Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
// Output: [7,9,6,6,8,7,3,0,9,5]

// Constraints:
//     The number of nodes in the list is n.
//     1 <= k <= n <= 10^5
//     0 <= Node.val <= 100

import "fmt"

type ListNode struct {
    Val int
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
func swapNodes(head *ListNode, k int) *ListNode {
    slow, fast := head, head
    for fast != nil && k > 1 { 
        fast = fast.Next
        k-- 
    }
    start := fast
    for fast.Next != nil { 
        slow = slow.Next
        fast = fast.Next 
    }
    end := slow
    start.Val, end.Val = end.Val, start.Val // swap
    return head
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/linked1.jpg" />
    // Input: head = [1,2,3,4,5], k = 2
    // Output: [1,4,3,2,5]
    list1 := makeListNode([]int{1,2,3,4,5})
    printListNode(list1) // 1 -> (2) -> 3 -> [4] -> 5
    printListNode(swapNodes(list1, 2)) // 1 -> [4] -> 3 -> (2) -> 5
    // Example 2:
    // Input: head = [7,9,6,6,7,8,3,0,9,5], k = 5
    // Output: [7,9,6,6,8,7,3,0,9,5]
    list2 := makeListNode([]int{7,9,6,6,7,8,3,0,9,5})
    printListNode(list2) // 7 -> 9 -> 6 -> 6 -> (7) -> [8] -> 3 -> 0 -> 9 -> 5
    printListNode(swapNodes(list2, 5)) // 7 -> 9 -> 6 -> 6 -> [8] -> (7) -> 3 -> 0 -> 9 -> 5
}