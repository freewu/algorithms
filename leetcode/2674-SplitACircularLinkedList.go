package main

// 2674. Split a Circular Linked List
// Given a circular linked list list of positive integers, 
// your task is to split it into 2 circular linked lists 
// so that the first one contains the first half of the nodes in list (exactly ceil(list.length / 2) nodes) in the same order they appeared in list, 
// and the second one contains the rest of the nodes in list in the same order they appeared in list.

// Return an array answer of length 2 in which the first element is a circular linked list representing the first half 
// and the second element is a circular linked list representing the second half.

// A circular linked list is a normal linked list with the only difference being 
// that the last node's next node, is the first node.

// Example 1:
// Input: nums = [1,5,7]
// Output: [[1,5],[7]]
// Explanation: The initial list has 3 nodes so the first half would be the first 2 elements since ceil(3 / 2) = 2 
// and the rest which is 1 node is in the second half.

// Example 2:
// Input: nums = [2,6,1,5]
// Output: [[2,6],[1,5]]
// Explanation: The initial list has 4 nodes so the first half would be the first 2 elements since ceil(4 / 2) = 
// and the rest which is 2 nodes are in the second half.

// Constraints:
//     The number of nodes in list is in the range [2, 10^5]
//     0 <= Node.val <= 10^9
//     LastNode.next = FirstNode where LastNode is the last node of the list and FirstNode is the first one

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
 func splitCircularLinkedList(list *ListNode) []*ListNode {
    // 快慢指针
    fast, slow := list, list
    for fast.Next != list && fast.Next.Next != list {
        slow, fast = slow.Next, fast.Next.Next 
    }
    t := slow.Next
    if fast.Next == list { fast.Next = t }
    if fast.Next.Next == list { fast.Next.Next = t }
    slow.Next = list
    return []*ListNode{ list, t }
}

func main() {
    // Example 1:
    // Input: nums = [1,5,7]
    // Output: [[1,5],[7]]
    // Explanation: The initial list has 3 nodes so the first half would be the first 2 elements since ceil(3 / 2) = 2 
    // and the rest which is 1 node is in the second half.
    list1 := makeListNode([]int{1,5,7})
    printListNode(list1)
    res1 := splitCircularLinkedList(makeListNode([]int{1,5,7}))
    fmt.Println(res1)
    // printListNode(res1[0])
    // printListNode(res1[1])
    // Example 2:
    // Input: nums = [2,6,1,5]
    // Output: [[2,6],[1,5]]
    // Explanation: The initial list has 4 nodes so the first half would be the first 2 elements since ceil(4 / 2) = 
    // and the rest which is 2 nodes are in the second half.
    list2 := makeListNode([]int{2,6,1,5})
    printListNode(list2)
    res2 := splitCircularLinkedList(makeListNode([]int{2,6,1,5}))
    // printListNode(res2[0])
    // printListNode(res2[1])
    fmt.Println(res2)
}