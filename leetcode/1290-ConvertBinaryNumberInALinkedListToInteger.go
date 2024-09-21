package main

// 1290. Convert Binary Number in a Linked List to Integer
// Given head which is a reference node to a singly-linked list. 
// The value of each node in the linked list is either 0 or 1. 
// The linked list holds the binary representation of a number.

// Return the decimal value of the number in the linked list.

// The most significant bit is at the head of the linked list.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/12/05/graph-1.png" />
// Input: head = [1,0,1]
// Output: 5
// Explanation: (101) in base 2 = (5) in base 10

// Example 2:
// Input: head = [0]
// Output: 0

// Constraints:
//     The Linked List is not empty.
//     Number of nodes will not exceed 30.
//     Each node's value is either 0 or 1.

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
func getDecimalValue(head *ListNode) int {
    res:= 0
    for head != nil {
        res = 2 * res + head.Val
        head = head.Next
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/12/05/graph-1.png" />
    // Input: head = [1,0,1]
    // Output: 5
    // Explanation: (101) in base 2 = (5) in base 10
    list1 := makeListNode([]int{1,0,1})
    printListNode(list1)
    fmt.Println(getDecimalValue(list1)) // 5
    // Example 2:
    // Input: head = [0]
    // Output: 0
    list2 := makeListNode([]int{0})
    printListNode(list2)
    fmt.Println(getDecimalValue(list2)) // 0

    list3 := makeListNode([]int{1,1,1,1,1,1,1,1})
    printListNode(list3)
    fmt.Println(getDecimalValue(list3)) // 255

    list4 := makeListNode([]int{1,1,1,1,1,1,1,1,1,1})
    printListNode(list4)
    fmt.Println(getDecimalValue(list4)) // 1023

    list5 := makeListNode([]int{1,0,0,1,0,0,1,1,1,0,0,0,0,0,0})
    printListNode(list5)
    fmt.Println(getDecimalValue(list5)) // 18880
}