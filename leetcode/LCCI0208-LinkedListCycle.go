package main

// 面试题 02.08. Linked List Cycle LCCI
// Given a circular linked list, implement an algorithm that returns the node at the beginning of the loop.

// Circular linked list: A (corrupt) linked list in which a node's next pointer points to an earlier node, so as to make a loop in the linked list.

// Example 1:
// Input: head = [3,2,0,-4], pos = 1
// Output: tail connects to node index 1

// Example 2:
// Input: head = [1,2], pos = 0
// Output: tail connects to node index 0

// Example 3:
// Input: head = [1], pos = -1
// Output: no cycle

// Follow Up:
//      you solve it without using additional space?

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
// hash
func detectCycle(head *ListNode) *ListNode {
    seen := make(map[*ListNode]bool)
    for head != nil {
        if seen[head] {
            return head
        }
        seen[head] = true
        head = head.Next
    }
    return nil
}

// 快慢指针
func detectCycle1(head *ListNode) *ListNode {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
        if slow == fast {
            break
        }
    }
    if fast == nil || fast.Next == nil {
        return nil
    }
    slow = head
    for fast != slow {
        fast, slow = fast.Next, slow.Next
    }
    return fast
}

func main() {
    // Example 1:
    // Input: head = [3,2,0,-4], pos = 1
    // Output: tail connects to node index 1
    list1 := makeListNode([]int{3,2,0,-4})
    printListNode(list1) // 3 -> 2 -> 0 -> -4
    printListNode(detectCycle(list1))
    // Example 2:
    // Input: head = [1,2], pos = 0
    // Output: tail connects to node index 0
    list2 := makeListNode([]int{1,2})
    printListNode(list2) // 1 -> 2
    printListNode(detectCycle(list2))
    // Example 3:
    // Input: head = [1], pos = -1
    // Output: no cycle
    list3 := makeListNode([]int{1})
    printListNode(list3) // 1
    printListNode(detectCycle(list3))
}
