package main

// 1019. Next Greater Node In Linked List
// You are given the head of a linked list with n nodes.

// For each node in the list, find the value of the next greater node. 
// That is, for each node, find the value of the first node that is next to it and has a strictly larger value than it.

// Return an integer array answer where answer[i] is the value of the next greater node of the ith node (1-indexed). 
// If the ith node does not have a next greater node, set answer[i] = 0.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext1.jpg" />
// Input: head = [2,1,5]
// Output: [5,5,0]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext2.jpg" />
// Input: head = [2,7,4,3,5]
// Output: [7,0,5,5,0]

// Constraints:
//     The number of nodes in the list is n.
//     1 <= n <= 10^4
//     1 <= Node.val <= 10^9

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
func nextLargerNodes(head *ListNode) []int {
    nodes := []int{}
    for head != nil {
        nodes = append(nodes, head.Val)
        head = head.Next
    }
    larger, stack := make([]int, len(nodes)), [][]int{}
    for i := 0; i < len(nodes); {
        if len(stack) == 0 || nodes[i] <= stack[len(stack)-1][0] {
            stack = append(stack, []int{nodes[i], i})
            i++
        } else {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            larger[top[1]] = nodes[i]
        }
    }
    for _, v := range stack {
        larger[v[1]] = 0
    }
    return larger
}

func nextLargerNodes1(head *ListNode) []int {
    nodes := []int{}
    for n := head; n != nil; n = n.Next { // 取出链表中的值
        nodes = append(nodes, n.Val)
    }
    res, stack := make([]int, len(nodes)), []int{}
    for i, n := range nodes {
        for len(stack) > 0 && nodes[stack[len(stack)-1]] < n {
            res[stack[len(stack)-1]] = n
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, i)
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext1.jpg" />
    // Input: head = [2,1,5]
    // Output: [5,5,0]
    list1 := makeListNode([]int{2,1,5})
    printListNode(list1) // 2 -> 1 -> 5
    fmt.Println(nextLargerNodes(list1)) // [5,5,0]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/08/05/linkedlistnext2.jpg" />
    // Input: head = [2,7,4,3,5]
    // Output: [7,0,5,5,0]
    list2 := makeListNode([]int{2,7,4,3,5})
    printListNode(list2) // 2 -> 7 -> 4 -> 3 -> 5
    fmt.Println(nextLargerNodes(list2)) // [7,0,5,5,0]

    list11 := makeListNode([]int{2,1,5})
    printListNode(list11) // 2 -> 1 -> 5
    fmt.Println(nextLargerNodes1(list11)) // [5,5,0]
    list12 := makeListNode([]int{2,7,4,3,5})
    printListNode(list12) // 2 -> 7 -> 4 -> 3 -> 5
    fmt.Println(nextLargerNodes1(list12)) // [7,0,5,5,0]
}