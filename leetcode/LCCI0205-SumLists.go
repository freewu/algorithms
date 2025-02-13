package main

// 面试题 02.05. Sum Lists LCCI
// You have two numbers represented by a linked list, where each node contains a single digit. 
// The digits are stored in reverse order, such that the 1's digit is at the head of the list. 
// Write a function that adds the two numbers and returns the sum as a linked list.

// Example1:
// Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
// Output: 2 -> 1 -> 9. That is, 912.
// Follow Up: Suppose the digits are stored in forward order. Repeat the above problem.

// Example2:
// Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
// Output: 9 -> 1 -> 2. That is, 912.

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
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
    var tail, head *ListNode
    carry := 0
    for l1 != nil || l2 != nil {
        n1, n2 := 0, 0
        if l1 != nil {
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            n2 = l2.Val
            l2 = l2.Next
        }
        sum := n1 + n2 + carry
        sum, carry = sum % 10, sum / 10
        if head == nil {
            head = &ListNode{Val: sum}
            tail = head
        } else {
            tail.Next = &ListNode{ Val: sum }
            tail = tail.Next
        }
    }
    if carry > 0 { // 需要进位
        tail.Next = &ListNode{ Val: carry }
    }
    return head
}

func main() {
    // Example1:
    // Input: (7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
    // Output: 2 -> 1 -> 9. That is, 912.
    // Follow Up: Suppose the digits are stored in forward order. Repeat the above problem.
    list11 := makeListNode([]int{7, 1, 6}) 
    list12 := makeListNode([]int{5, 9, 2}) 
    printListNode(list11) // 7 -> 1 -> 6
    printListNode(list12) // 5 -> 9 -> 2
    printListNode(addTwoNumbers(list11, list12)) // 2 -> 1 -> 9
    // Example2:
    // Input: (6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
    // Output: 9 -> 1 -> 2. That is, 912.
    list21 := makeListNode([]int{6, 1, 7}) 
    list22 := makeListNode([]int{2, 9, 5}) 
    printListNode(list21) // 6 -> 1 -> 7
    printListNode(list22) // 2 -> 9 -> 5
    printListNode(addTwoNumbers(list21, list22)) // 8 -> 0 -> 3 -> 1
}