package main

// 817. Linked List Components
// You are given the head of a linked list containing unique integer values 
// and an integer array nums that is a subset of the linked list values.

// Return the number of connected components in nums 
// where two values are connected if they appear consecutively in the linked list.

// Example 1:
//     0 -> 1 -> 2 -> 3
// <img src="https://assets.leetcode.com/uploads/2021/07/22/lc-linkedlistcom1.jpg" />
// Input: head = [0,1,2,3], nums = [0,1,3]
// Output: 2
// Explanation: 0 and 1 are connected, so [0, 1] and [3] are the two connected components.

// Example 2:
//     0 -> 1 -> 2 -> 3 -> 4
// <img src="https://assets.leetcode.com/uploads/2021/07/22/lc-linkedlistcom2.jpg" />
// Input: head = [0,1,2,3,4], nums = [0,3,1,4]
// Output: 2
// Explanation: 0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.

// Constraints:
//     The number of nodes in the linked list is n.
//     1 <= n <= 10^4
//     0 <= Node.val < n
//     All the values Node.val are unique.
//     1 <= nums.length <= n
//     0 <= nums[i] < n
//     All the values of nums are unique.

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
func numComponents(head *ListNode, nums []int) int {
    targets := make(map[int]bool)
    for _, v := range nums {
        targets[v] = true
    }
    res := 0
    var prev *ListNode
    for cur := head; cur != nil; cur = cur.Next {
        if prev != nil {
            if !targets[prev.Val] && targets[cur.Val] {
                res++
            }
        } else {
            if targets[cur.Val] {
                res++
            }
        }
        prev = cur
    }
    return res
}

func main() {
    // Example 1:
    //     0 -> 1 -> 2 -> 3
    // <img src="https://assets.leetcode.com/uploads/2021/07/22/lc-linkedlistcom1.jpg" />
    // Input: head = [0,1,2,3], nums = [0,1,3]
    // Output: 2
    // Explanation: 0 and 1 are connected, so [0, 1] and [3] are the two connected components.
    list1 := makeListNode([]int{0,1,2,3})
    printListNode(list1)
    fmt.Println(numComponents(list1,[]int{0,1,3})) // 2
    // Example 2:
    //     0 -> 1 -> 2 -> 3 -> 4
    // <img src="https://assets.leetcode.com/uploads/2021/07/22/lc-linkedlistcom2.jpg" />
    // Input: head = [0,1,2,3,4], nums = [0,3,1,4]
    // Output: 2
    // Explanation: 0 and 1 are connected, 3 and 4 are connected, so [0, 1] and [3, 4] are the two connected components.
    list2 := makeListNode([]int{0,1,2,3,4})
    printListNode(list2)
    fmt.Println(numComponents(list2,[]int{0,3,1,4})) // 2
}