package main

// 2816. Double a Number Represented as a Linked List
// You are given the head of a non-empty linked list representing a non-negative integer without leading zeroes.
// Return the head of the linked list after doubling it.

// Example 1:
// 1 -> 8 -> 9   =>  3 -> 7 -> 8
// Input: head = [1,8,9]
// Output: [3,7,8]
// Explanation: The figure above corresponds to the given linked list which represents the number 189. Hence, the returned linked list represents the number 189 * 2 = 378.

// Example 2:
// 9 -> 9 -> 9   =>   1 -> 9 -> 9 -> 8
// Input: head = [9,9,9]
// Output: [1,9,9,8]
// Explanation: The figure above corresponds to the given linked list which represents the number 999. Hence, the returned linked list reprersents the number 999 * 2 = 1998. 
 
// Constraints:
//     The number of nodes in the list is in the range [1, 10^4]
//     0 <= Node.val <= 9
//     The input is generated such that the list represents a number that does not have leading zeros, except the number 0 itself.

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
    l := len(arr) - 1
    head := &ListNode{arr[l], nil}
    for i := l - 1; i >= 0; i-- {
        n := &ListNode{arr[i], head}
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
func doubleIt(head *ListNode) *ListNode {
    p1, p2 := head, &ListNode{Next: head}
    if p1.Val >= 5 { // 需要进位处理 5 * 2 = 10
        head = p2
    }
    for p1 != nil {
        val := p1.Val << 1 // val / 2
        p1.Val = val % 10
        p2.Val += val / 10
        p1 = p1.Next
        p2 = p2.Next
    }
    return head
}

func main() {
    // Example 1:
    // 1 -> 8 -> 9   =>  3 -> 7 -> 8
    // Input: head = [1,8,9]
    // Output: [3,7,8]
    // Explanation: The figure above corresponds to the given linked list which represents the number 189. Hence, the returned linked list represents the number 189 * 2 = 378.
    l1 := makeListNode([]int{1,8,9})
    printListNode(l1) // 1 -> 8 -> 9
    printListNode(doubleIt(l1)) // 3 -> 7 -> 8
    // Example 2:
    // 9 -> 9 -> 9   =>   1 -> 9 -> 9 -> 8
    // Input: head = [9,9,9]
    // Output: [1,9,9,8]
    // Explanation: The figure above corresponds to the given linked list which represents the number 999. Hence, the returned linked list reprersents the number 999 * 2 = 1998. 
    l2 := makeListNode([]int{9,9,9})
    printListNode(l2) // 9 -> 9 -> 9
    printListNode(doubleIt(l2)) // 1 -> 9 -> 9 -> 8
}