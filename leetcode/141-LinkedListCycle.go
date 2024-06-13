package main

// 141. Linked List Cycle
// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. 
// Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png">
// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png">
// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

// Example 3:
// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.
 
// Constraints:
//         The number of the nodes in the list is in the range [0, 104].
//         -10^5 <= Node.val <= 10^5
//         pos is -1 or a valid index in the linked-list.

// Follow up: Can you solve it using O(1) (i.e. constant) memory?

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
func hasCycle(head *ListNode) bool {
    // 给 2 个指针，一个指针是另外一个指针的下一个指针。
    fast, slow := head, head
    for fast != nil && fast.Next != nil {
        fast = fast.Next.Next // 快指针一次走 2 格
        slow = slow.Next // 慢指针一次走 1 格
        if fast == slow { // 如果存在环，那么前一个指针一定会经过若干圈之后追上慢的指针
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png">
    // Input: head = [3,2,0,-4], pos = 1
    // Output: true
    // Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
    fmt.Println(hasCycle(makeListNode([]int{3,2,0,-4}))) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png">
    // Input: head = [1,2], pos = 0
    // Output: true
    // Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.
    fmt.Println(hasCycle(makeListNode([]int{1,2}))) // true
    // Example 3:
    // Input: head = [1], pos = -1
    // Output: false
    // Explanation: There is no cycle in the linked list.
    fmt.Println(hasCycle(makeListNode([]int{1}))) // false
}