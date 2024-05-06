package main 

// 109. Convert Sorted List to Binary Search Tree
// Given the head of a singly linked list where elements are sorted in ascending order, 
// convert it to a height-balanced binary search tree.

// Example 1:
//                                         0
// -10 -> -3 -> 0 -> 5 -> 9   =>          /  \
//                                      -3    9
//                                      /    /
//                                   -10    5
// <img src="https://assets.leetcode.com/uploads/2020/08/17/linked.jpg" / >
// Input: head = [-10,-3,0,5,9]
// Output: [0,-3,9,-10,null,5]
// Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.

// Example 2:
// Input: head = []
// Output: []
 
// Constraints:
//     The number of nodes in head is in the range [0, 2 * 10^4].
//     -10^5 <= Node.val <= 10^5

import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
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
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 递归
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil { // 空节点
        return nil
    }
    if head.Next == nil { // 只有一个节点
        return &TreeNode{head.Val, nil, nil}
    }
    prevSlow, slow, fast := head, head, head
    for fast != nil && fast.Next != nil { // 快慢指针找到中间位置
        prevSlow = slow
        slow = slow.Next
        fast = fast.Next.Next
    }
    node := new(TreeNode)
    node.Val = slow.Val
    prevSlow.Next = nil // 断开
    node.Left = sortedListToBST(head)
    node.Right = sortedListToBST(slow.Next)
    return node
}

func main() {
    // Example 1:
    //                                         0
    // -10 -> -3 -> 0 -> 5 -> 9   =>          /  \
    //                                      -3    9
    //                                      /    /
    //                                   -10    5
    // Input: head = [-10,-3,0,5,9]
    // Output: [0,-3,9,-10,null,5]
    // Explanation: One possible answer is [0,-3,9,-10,null,5], which represents the shown height balanced BST.
    l1 := makeListNode([]int{-10,-3,0,5,9})
    printListNode(l1)
    t1 := sortedListToBST(l1)
    fmt.Println(t1.Val) // 0
    fmt.Println(t1.Left.Val) // -3
    fmt.Println(t1.Right.Val) // 9
    // Example 2:
    // Input: head = []
    // Output: []
    l2 := makeListNode([]int{})
    printListNode(l2)
    t2 := sortedListToBST(l2)
    fmt.Println(t2) // nil
}