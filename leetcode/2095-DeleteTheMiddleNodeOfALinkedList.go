package main

// 2095. Delete the Middle Node of a Linked List
// You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.
// The middle node of a linked list of size n is the ⌊ n / 2 ⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.
// For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.
 
// Example 1:
// 1 -> 3 -> 4 -> [7] -> 1 -> 2 -> 6
// <img src="https://assets.leetcode.com/uploads/2021/11/16/eg1drawio.png" />
// Input: head = [1,3,4,7,1,2,6]
// Output: [1,3,4,1,2,6]
// Explanation:
// The above figure represents the given linked list. The indices of the nodes are written below.
// Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
// We return the new list after removing this node. 

// Example 2:
// 1 -> 2 -> [3] -> 4
// <img src="https://assets.leetcode.com/uploads/2021/11/16/eg2drawio.png" />
// Input: head = [1,2,3,4]
// Output: [1,2,4]
// Explanation:
// The above figure represents the given linked list.
// For n = 4, node 2 with value 3 is the middle node, which is marked in red.

// Example 3:
// 2 -> [1]
// <img src="https://assets.leetcode.com/uploads/2021/11/16/eg3drawio.png" />
// Input: head = [2,1]
// Output: [2]
// Explanation:
// The above figure represents the given linked list.
// For n = 2, node 1 with value 1 is the middle node, which is marked in red.
// Node 0 with value 2 is the only node remaining after removing node 1.
 
// Constraints:
//     The number of nodes in the list is in the range [1, 10^5].
//     1 <= Node.val <= 10^5

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
// 快慢指针
func deleteMiddle(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return nil
    }
    previous, slow, fast := head, head, head
    // Time: O(n/2) = O(n)
    for slow != nil && fast != nil && fast.Next != nil { // slow 走一步 fast 走两步
        previous, slow, fast = slow, slow.Next, fast.Next.Next
    }
    previous.Next = slow.Next // fast 走完退出 slow 刚好在中间
    return head
}

func deleteMiddle1(head *ListNode) *ListNode {
    length := 0
    for i := head; i != nil; i = i.Next { // 得到链表长度
        length++
    }
    if length == 1 {
        return nil
    }
    middle := length / 2
    node := head
    for i := 0; i < middle - 1; i++ { // 循环到中间位置
        node = node.Next
    }
    node.Next = node.Next.Next // 删除中间节点
    return head
}

func main() {
    // Example 1:
    // 1 -> 3 -> 4 -> [7] -> 1 -> 2 -> 6
    // <img src="https://assets.leetcode.com/uploads/2021/11/16/eg1drawio.png" />
    // Input: head = [1,3,4,7,1,2,6]
    // Output: [1,3,4,1,2,6]
    // Explanation:
    // The above figure represents the given linked list. The indices of the nodes are written below.
    // Since n = 7, node 3 with value 7 is the middle node, which is marked in red.
    // We return the new list after removing this node. 
    l1 := makeListNode([]int{1,3,4,7,1,2,6})
    printListNode(l1) // 1 -> 3 -> 4 -> 7 -> 1 -> 2 -> 6
    printListNode(deleteMiddle(l1)) // 1 -> 3 -> 4 -> 1 -> 2 -> 6
    // Example 2:
    // 1 -> 2 -> [3] -> 4
    // <img src="https://assets.leetcode.com/uploads/2021/11/16/eg2drawio.png" />
    // Input: head = [1,2,3,4]
    // Output: [1,2,4]
    // Explanation:
    // The above figure represents the given linked list.
    // For n = 4, node 2 with value 3 is the middle node, which is marked in red.
    l2 := makeListNode([]int{1,2,3,4}) 
    printListNode(l2) // 1 -> 2 -> 3 -> 4
    printListNode(deleteMiddle(l2)) // 1 -> 2 -> 4
    // Example 3:
    // 2 -> [1]
    // <img src="https://assets.leetcode.com/uploads/2021/11/16/eg3drawio.png" />
    // Input: head = [2,1]
    // Output: [2]
    // Explanation:
    // The above figure represents the given linked list.
    // For n = 2, node 1 with value 1 is the middle node, which is marked in red.
    // Node 0 with value 2 is the only node remaining after removing node 1.
    l3 := makeListNode([]int{2,1})
    printListNode(l3) // 2 -> 1
    printListNode(deleteMiddle(l3)) // 2

    l11 := makeListNode([]int{1,3,4,7,1,2,6})
    printListNode(l11) // 1 -> 3 -> 4 -> 7 -> 1 -> 2 -> 6
    printListNode(deleteMiddle(l11)) // 1 -> 3 -> 4 -> 1 -> 2 -> 6
    l12 := makeListNode([]int{1,2,3,4}) 
    printListNode(l12) // 1 -> 2 -> 3 -> 4
    printListNode(deleteMiddle(l12)) // 1 -> 2 -> 4
    l13 := makeListNode([]int{2,1})
    printListNode(l13) // 2 -> 1
    printListNode(deleteMiddle(l13)) // 2
}