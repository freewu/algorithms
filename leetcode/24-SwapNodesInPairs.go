package main

// 24. Swap Nodes in Pairs
// Given a linked list, swap every two adjacent nodes and return its head. 
// You must solve the problem without modifying the values in the list's nodes 
// (i.e., only nodes themselves may be changed.)

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/03/swap_ex1.jpg" />
// Input: head = [1,2,3,4]
// Output: [2,1,4,3]

// Example 2:
// Input: head = []
// Output: []

// Example 3:
// Input: head = [1]
// Output: [1]
 
// Constraints:
//     The number of nodes in the list is in the range [0, 100].
//     0 <= Node.val <= 100

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
// 递归
func swapPairs(head *ListNode) *ListNode {
    var swap func (p *ListNode) *ListNode
    swap = func (p *ListNode) *ListNode {
        if p == nil || p.Next == nil {
            return p
        }
        q := p.Next
        // 1 交换头两个节点
        tail := swap(q.Next) // 2 交换头两个节点之后的链表(这就是重复的子问题, 符合递归)
        q.Next = p
        p.Next = tail // 3 将 1 的节点链接到 2的结果
        return q
    }
    return swap(head)
}

func swapPairs1(head *ListNode) *ListNode {
    
    if head == nil || head.Next == nil {
        return head 
    }
    newHead := head.Next
    head.Next = swapPairs1(newHead.Next)
    newHead.Next = head
    return newHead
}

func swapPairs2(head *ListNode) *ListNode {
    var swap func (head *ListNode) *ListNode
    swap = func (head *ListNode) *ListNode {
        if head == nil || head.Next == nil {
            return head 
        }
        newHead := head.Next
        head.Next = swap(newHead.Next)
        newHead.Next = head
        return newHead
    }
    return swap(head)
}


func main() {
    // Input: head = [1,2,3,4]
    // Output: [2,1,4,3]
    l1 := makeListNode([]int{1,2,3,4}) // 1 -> 2 -> 3 -> 4
    fmt.Println("before: ")
    printListNode(l1)
    fmt.Println("after: ")
    printListNode(swapPairs(l1)) // [2,1,4,3]  2 -> 1 -> 4 -> 3

    l2 := makeListNode([]int{1})
    fmt.Println("before: ")
    printListNode(l2)
    fmt.Println("after: ")
    printListNode(swapPairs(l2)) // [1]

    l11 := makeListNode([]int{1,2,3,4}) // 1 -> 2 -> 3 -> 4
    fmt.Println("before: ")
    printListNode(l11)
    fmt.Println("after: ")
    printListNode(swapPairs1(l11)) // [2,1,4,3]  2 -> 1 -> 4 -> 3

    l12 := makeListNode([]int{1})
    fmt.Println("before: ")
    printListNode(l12)
    fmt.Println("after: ")
    printListNode(swapPairs1(l12)) // [1]

    l21 := makeListNode([]int{1,2,3,4}) // 1 -> 2 -> 3 -> 4
    fmt.Println("before: ")
    printListNode(l21)
    fmt.Println("after: ")
    printListNode(swapPairs2(l21)) // [2,1,4,3]  2 -> 1 -> 4 -> 3

    l22 := makeListNode([]int{1})
    fmt.Println("before: ")
    printListNode(l22)
    fmt.Println("after: ")
    printListNode(swapPairs2(l22)) // [1]
}