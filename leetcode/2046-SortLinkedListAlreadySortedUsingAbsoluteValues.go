package main

// 2046. Sort Linked List Already Sorted Using Absolute Values
// Given the head of a singly linked list that is sorted in non-decreasing order using the absolute values of its nodes,
// return the list sorted in non-decreasing order using the actual values of its nodes.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/17/image-20211017201240-3.png">
// Input: head = [0,2,-5,5,10,-10]
// Output: [-10,-5,0,2,5,10]
// Explanation:
// The list sorted in non-descending order using the absolute values of the nodes is [0,2,-5,5,10,-10].
// The list sorted in non-descending order using the actual values is [-10,-5,0,2,5,10].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/17/image-20211017201318-4.png">
// Input: head = [0,1,2]
// Output: [0,1,2]
// Explanation:
// The linked list is already sorted in non-decreasing order.

// Example 3:
// Input: head = [1]
// Output: [1]
// Explanation:
// The linked list is already sorted in non-decreasing order.

// Constraints:
//     The number of nodes in the list is the range [1, 10^5].
//     -5000 <= Node.val <= 5000
//     head is sorted in non-decreasing order using the absolute value of its nodes.

// Follow up: Can you think of a solution with O(n) time complexity?

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
func sortLinkedList(head *ListNode) *ListNode {
    dummy := &ListNode{-1, head}
    a, b := dummy, head
    if b.Val < 0 {
        a, b = b, b.Next
    }
    for b != nil {
        if b.Val < 0 {
            a.Next = b.Next
            b.Next = dummy.Next
            dummy.Next = b
            b = a.Next
        } else {
            a, b = b, b.Next
        }
    }
    return dummy.Next
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/17/image-20211017201240-3.png">
    // Input: head = [0,2,-5,5,10,-10]
    // Output: [-10,-5,0,2,5,10]
    // Explanation:
    // The list sorted in non-descending order using the absolute values of the nodes is [0,2,-5,5,10,-10].
    // The list sorted in non-descending order using the actual values is [-10,-5,0,2,5,10].
    list1 := makeListNode([]int{0,2,-5,5,10,-10})
    printListNode(list1) // 0 -> 2 -> -5 -> 5 -> 10 -> -10
    printListNode(sortLinkedList(list1)) // -10 -> -5 -> 0 -> 2 -> 5 -> 10
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/17/image-20211017201318-4.png">
    // Input: head = [0,1,2]
    // Output: [0,1,2]
    // Explanation:
    // The linked list is already sorted in non-decreasing order.
    list2 := makeListNode([]int{0,1,2})
    printListNode(list2) // 0 -> 1 -> 2
    printListNode(sortLinkedList(list2)) // 0 -> 1 -> 2
    // Example 3:
    // Input: head = [1]
    // Output: [1]
    // Explanation:
    // The linked list is already sorted in non-decreasing order.
    list3 := makeListNode([]int{1}) 
    printListNode(list3) // 1
    printListNode(sortLinkedList(list3)) // 1
}