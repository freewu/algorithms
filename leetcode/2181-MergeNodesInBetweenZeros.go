package main

// 2181. Merge Nodes in Between Zeros
// You are given the head of a linked list, which contains a series of integers separated by 0's. 
// The beginning and end of the linked list will have Node.val == 0.

// For every two consecutive 0's, merge all the nodes lying in between them into a single node whose value is the sum of all the merged nodes. 
// The modified list should not contain any 0's.

// Return the head of the modified linked list.

// Example 1:
// 0 -> (3) -> (1) -> 0 -> (4) -> (5) -> (2) -> 0
//          |                     |
//          4        ->           11   
// <img src="https://assets.leetcode.com/uploads/2022/02/02/ex1-1.png" />
// Input: head = [0,3,1,0,4,5,2,0]
// Output: [4,11]
// Explanation: 
// The above figure represents the given linked list. The modified list contains
// - The sum of the nodes marked in green: 3 + 1 = 4.
// - The sum of the nodes marked in red: 4 + 5 + 2 = 11.

// Example 2:
// 0 -> (1) -> 0 -> (3) -> 0 -> (2) -> (2) -> 0
//       |           |               |
//       1     ->    3     ->        4
// <img src="https://assets.leetcode.com/uploads/2022/02/02/ex2-1.png" />
// Input: head = [0,1,0,3,0,2,2,0]
// Output: [1,3,4]
// Explanation: 
// The above figure represents the given linked list. The modified list contains
// - The sum of the nodes marked in green: 1 = 1.
// - The sum of the nodes marked in red: 3 = 3.
// - The sum of the nodes marked in yellow: 2 + 2 = 4.

// Constraints:
//     The number of nodes in the list is in the range [3, 2 * 10^5].
//     0 <= Node.val <= 1000
//     There are no two consecutive nodes with Node.val == 0.
//     The beginning and end of the linked list have Node.val == 0.

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
func mergeNodes(head *ListNode) *ListNode {
    var res *ListNode
    sum := 0
    for head != nil{
        if head.Val == 0 { 
            if sum != 0 {
                t := &ListNode{sum, res }
                res = t
            }
            sum = 0
        } else {
            sum += head.Val
        }
        head = head.Next
    }
    reverse := func (head *ListNode) *ListNode {
        var res *ListNode
        for head != nil {
            next := head.Next  // 备份head.Next
            head.Next = res // 更新  head.Next
            res = head      // 移动 new_head
            head = next
        }
        return res
    }
    return reverse(res)
}

func mergeNodes1(head *ListNode) *ListNode {
    head = head.Next
    cur := head
    for cur != nil {
        if cur.Next.Val != 0 {
            cur.Val += cur.Next.Val
            cur.Next = cur.Next.Next
        } else {
            cur.Next = cur.Next.Next
            cur = cur.Next
        }
    }
    return head
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/02/02/ex1-1.png" />
    // Input: head = [0,3,1,0,4,5,2,0]
    // Output: [4,11]
    // Explanation: 
    // The above figure represents the given linked list. The modified list contains
    // - The sum of the nodes marked in green: 3 + 1 = 4.
    // - The sum of the nodes marked in red: 4 + 5 + 2 = 11.
    list1 := makeListNode([]int{0,3,1,0,4,5,2,0})
    printListNode(list1) // 0 -> 3 -> 1 -> 0 -> 4 -> 5 -> 2 -> 0
    printListNode(mergeNodes(list1)) // [4,11]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/02/02/ex2-1.png" />
    // Input: head = [0,1,0,3,0,2,2,0]
    // Output: [1,3,4]
    // Explanation: 
    // The above figure represents the given linked list. The modified list contains
    // - The sum of the nodes marked in green: 1 = 1.
    // - The sum of the nodes marked in red: 3 = 3.
    // - The sum of the nodes marked in yellow: 2 + 2 = 4.
    list2 := makeListNode([]int{0,1,0,3,0,2,2,0})
    printListNode(list2)
    printListNode(mergeNodes(list2)) // [1,3,4]

    list11 := makeListNode([]int{0,3,1,0,4,5,2,0})
    printListNode(list11) // 0 -> 3 -> 1 -> 0 -> 4 -> 5 -> 2 -> 0
    printListNode(mergeNodes1(list11)) // [4,11]
    list12 := makeListNode([]int{0,1,0,3,0,2,2,0})
    printListNode(list12)
    printListNode(mergeNodes1(list12)) // [1,3,4]
}