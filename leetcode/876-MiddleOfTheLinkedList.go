package main

// 876. Middle of the Linked List
// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.

// Example 1:
// [1] -> [2] -> (3) -> [4] -> [5]
//               (3) -> [4] -> [5]
// Input: head = [1,2,3,4,5]
// Output: [3,4,5]
// Explanation: The middle node of the list is node 3.

// Example 2:
// [1] -> [2] -> [3] -> (4) -> [5] -> [6] 
//                      (4) -> [5] -> [6] 
// Input: head = [1,2,3,4,5,6]
// Output: [4,5,6]
// Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.
 
// Constraints:
//         The number of nodes in the list is in the range [1, 100].
//         1 <= Node.val <= 100

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
func middleNode(head *ListNode) *ListNode {
    // 使用快慢指针
    // 快 1次2步 fast.Next.Next
    // 慢 1次2步 slow.Next
    slow, fast := head, head
    for(  fast != nil && fast.Next != nil ) { // 当 fast 到了最后一个位置 
        fast = fast.Next.Next
        slow = slow.Next
    }
    return slow // slow 就到中间位置了
}

func main() {
    printListNode(makeListNode([]int{1,2,3,4,5})) // [1] -> [2] -> (3) -> [4] -> [5]
    printListNode(middleNode(makeListNode([]int{1,2,3,4,5}))) // (3) -> [4] -> [5]

    printListNode(makeListNode([]int{1,2,3,4,5,6})) // [1] -> [2] -> [3] -> (4) -> [5] -> [6]
    printListNode(middleNode(makeListNode([]int{1,2,3,4,5,6}))) // (4) -> [5] -> [6]
}