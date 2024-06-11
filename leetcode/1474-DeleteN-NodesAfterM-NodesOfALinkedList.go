package main

// 1474. Delete N Nodes After M Nodes of a Linked List
// You are given the head of a linked list and two integers m and n.
// Traverse the linked list and remove some nodes in the following way:
//     Start with the head as the current node.
//     Keep the first m nodes starting with the current node.
//     Remove the next n nodes
//     Keep repeating steps 2 and 3 until you reach the end of the list.

// Return the head of the modified list after removing the mentioned nodes.

// Example 1:
// 1 -> 2 -> (3) -> (4) -> (5) -> 6 -> 7 -> (8) -> (9) -> (10) -> 11 -> 12 -> (13)
// 1 -> 2 ----------------------> 6 -> 7 -----------------------> 11 -> 12
// <img src="https://assets.leetcode.com/uploads/2020/06/06/sample_1_1848.png" />
// Input: head = [1,2,3,4,5,6,7,8,9,10,11,12,13], m = 2, n = 3
// Output: [1,2,6,7,11,12]
// Explanation: Keep the first (m = 2) nodes starting from the head of the linked List  (1 ->2) show in black nodes.
// Delete the next (n = 3) nodes (3 -> 4 -> 5) show in read nodes.
// Continue with the same procedure until reaching the tail of the Linked List.
// Head of the linked list after removing nodes is returned.

// Example 2:
// 1 -> (2) -> (3) -> (4) -> 5 -> (6) -> (7) -> (8) -> 9 -> (10) -> (11)
// 1 ----------------------> 5 ----------------------> 9
// <img src="https://assets.leetcode.com/uploads/2020/06/06/sample_2_1848.png" />
// Input: head = [1,2,3,4,5,6,7,8,9,10,11], m = 1, n = 3
// Output: [1,5,9]
// Explanation: Head of linked list after removing nodes is returned.
 
// Constraints:
//     The number of nodes in the list is in the range [1, 10^4].
//     1 <= Node.val <= 10^6
//     1 <= m, n <= 1000

// Follow up: Could you solve this problem by modifying the list in-place?

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
func deleteNodes(head *ListNode, m int, n int) *ListNode {
    dummy := &ListNode{ Next: head }
    cur := dummy
    for cur != nil {
        for i := 0; i < m && cur != nil; i++ { // 保留 m 个节点
            cur = cur.Next
        }
        if cur == nil {
            break
        }
        // 删除 n 个节点
        del := cur
        for i := 0; i < n && del.Next != nil; i++ {
            del = del.Next
        }
        cur.Next = del.Next
    }
    return dummy.Next
}

func main() {
    // Example 1:
    // 1 -> 2 -> (3) -> (4) -> (5) -> 6 -> 7 -> (8) -> (9) -> (10) -> 11 -> 12 -> (13)
    // 1 -> 2 ----------------------> 6 -> 7 -----------------------> 11 -> 12
    // <img src="https://assets.leetcode.com/uploads/2020/06/06/sample_1_1848.png" />
    // Input: head = [1,2,3,4,5,6,7,8,9,10,11,12,13], m = 2, n = 3
    // Output: [1,2,6,7,11,12]
    // Explanation: Keep the first (m = 2) nodes starting from the head of the linked List  (1 ->2) show in black nodes.
    // Delete the next (n = 3) nodes (3 -> 4 -> 5) show in read nodes.
    // Continue with the same procedure until reaching the tail of the Linked List.
    // Head of the linked list after removing nodes is returned.
    list1 := makeListNode([]int{1,2,3,4,5,6,7,8,9,10,11,12,13}) // 
    printListNode(list1) // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11 -> 12 -> 13
    printListNode(deleteNodes(list1,2,3)) // 1 -> 2 -> 6 -> 7 -> 11 -> 12
    // Example 2:
    // 1 -> (2) -> (3) -> (4) -> 5 -> (6) -> (7) -> (8) -> 9 -> (10) -> (11)
    // 1 ----------------------> 5 ----------------------> 9
    // <img src="https://assets.leetcode.com/uploads/2020/06/06/sample_2_1848.png" />
    // Input: head = [1,2,3,4,5,6,7,8,9,10,11], m = 1, n = 3
    // Output: [1,5,9]
    // Explanation: Head of linked list after removing nodes is returned.
    list2 := makeListNode([]int{1,2,3,4,5,6,7,8,9,10,11})
    printListNode(list2) // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11
    printListNode(deleteNodes(list2,1,3)) // 1 -> 5 -> 9
}