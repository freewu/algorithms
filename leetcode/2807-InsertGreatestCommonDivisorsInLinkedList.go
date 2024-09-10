package main

// 2807. Insert Greatest Common Divisors in Linked List
// Given the head of a linked list head, in which each node contains an integer value.

// Between every pair of adjacent nodes, insert a new node with a value equal to the greatest common divisor of them.

// Return the linked list after insertion.

// The greatest common divisor of two numbers is the largest positive integer that evenly divides both numbers.

// Example 1:
// 18 -> 6 -> 10 -> 3
//         |
// 18 -> 6 -> 6 -> 2 -> 10 -> 1 -> 3
// <img src="https://assets.leetcode.com/uploads/2023/07/18/ex1_copy.png" />
// Input: head = [18,6,10,3]
// Output: [18,6,6,2,10,1,3]
// Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram denotes the linked list after inserting the new nodes (nodes in blue are the inserted nodes).
// - We insert the greatest common divisor of 18 and 6 = 6 between the 1st and the 2nd nodes.
// - We insert the greatest common divisor of 6 and 10 = 2 between the 2nd and the 3rd nodes.
// - We insert the greatest common divisor of 10 and 3 = 1 between the 3rd and the 4th nodes.
// There are no more adjacent nodes, so we return the linked list.

// Example 2:
//      7
//      |
//      7
// <img src="https://assets.leetcode.com/uploads/2023/07/18/ex2_copy1.png" />
// Input: head = [7]
// Output: [7]
// Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram denotes the linked list after inserting the new nodes.
// There are no pairs of adjacent nodes, so we return the initial linked list.

// Constraints:
//     The number of nodes in the list is in the range [1, 5000].
//     1 <= Node.val <= 1000

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
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    gcd := func (x, y int) int { for y != 0 { x, y = y, x % y; }; return x; }
    prev, curr := head, head.Next
    for curr != nil {
        prev.Next = &ListNode{ gcd(prev.Val, curr.Val), curr }
        prev, curr = curr, curr.Next
    }
    return head
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2023/07/18/ex1_copy.png" />
    // Input: head = [18,6,10,3]
    // Output: [18,6,6,2,10,1,3]
    // Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram denotes the linked list after inserting the new nodes (nodes in blue are the inserted nodes).
    // - We insert the greatest common divisor of 18 and 6 = 6 between the 1st and the 2nd nodes.
    // - We insert the greatest common divisor of 6 and 10 = 2 between the 2nd and the 3rd nodes.
    // - We insert the greatest common divisor of 10 and 3 = 1 between the 3rd and the 4th nodes.
    // There are no more adjacent nodes, so we return the linked list.
    list1 := makeListNode([]int{18,6,10,3})
    printListNode(list1) // 18 -> 6 -> 10 -> 3
    printListNode(insertGreatestCommonDivisors(list1)) // 18 -> 6 -> 6 -> 2 -> 10 -> 1 -> 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2023/07/18/ex2_copy1.png" />
    // Input: head = [7]
    // Output: [7]
    // Explanation: The 1st diagram denotes the initial linked list and the 2nd diagram denotes the linked list after inserting the new nodes.
    // There are no pairs of adjacent nodes, so we return the initial linked list.
    list2 := makeListNode([]int{7})
    printListNode(list2) // 7
    printListNode(insertGreatestCommonDivisors(list2)) // 7
}