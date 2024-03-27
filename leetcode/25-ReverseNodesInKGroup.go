package main

// 25. Reverse Nodes in k-Group
// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
// k is a positive integer and is less than or equal to the length of the linked list. 
// If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.

// You may not alter the values in the list's nodes, only nodes themselves may be changed.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/03/reverse_ex1.jpg" />
// Input: head = [1,2,3,4,5], k = 2
// Output: [2,1,4,3,5]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/03/reverse_ex2.jpg" />
// Input: head = [1,2,3,4,5], k = 3
// Output: [3,2,1,4,5]
 
// Constraints:
//     The number of nodes in the list is n.
//     1 <= k <= n <= 5000
//     0 <= Node.val <= 1000

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
func reverseKGroup(head *ListNode, k int) *ListNode {
    cnt := 0
    // 统计节点个数
    for p := head; p != nil; p = p.Next {
        cnt++
    }
    dummy := &ListNode{-1, head}
    reversedListTail := dummy
    pre, cur := (*ListNode)(nil), head 
    for cnt >= k {
        for i := 0; i < k; i++ {
            nxt := cur.Next 
            cur.Next = pre
            pre = cur 
            cur = nxt 
        }
        tail := reversedListTail.Next 
        tail.Next = cur 
        reversedListTail.Next = pre 
        reversedListTail = tail 
        cnt -= k 
    }
    return dummy.Next
}

func reverseKGroup1(head *ListNode, k int) *ListNode {
    cur := head
    index:=0
    var start,lastEnd,res *ListNode
    reverse := func (start,end *ListNode) *ListNode {
        cur := start 
        var last *ListNode
        for cur != end {
            next := cur.Next
            if last == nil {
                cur.Next = end
            } else {
                cur.Next = last
            }
            last = cur
            cur = next
        }
        return last
    }
    for cur != nil {
        index++
        if index == 1 {
            start = cur
        }
        cur = cur.Next

        if index == k {
            // tmp := start
            h := reverse(start,cur)
            if res == nil {
                res = h
            } else {
                lastEnd.Next = h
            }
            lastEnd = start
            start = cur
            index=0
        }
    }
    if res == nil {
        res = head
    }
    return res
}

func main() {
    l1 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("before: ")
    printListNode(l1) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("after: ")
    printListNode(reverseKGroup(l1, 2)) // [2,1,4,3,5] [2 -> 1] -> [4 -> 3] -> 5

    l2 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("before: ")
    printListNode(l2) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("after: ")
    printListNode(reverseKGroup(l2, 3)) // [3,2,1,4,5] [3 -> 2 -> 1] -> 4 -> 5

    l3 := makeListNode([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15})
    fmt.Println("before: ")
    printListNode(l3) // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11 -> 12 -> 13 -> 14 -> 15
    fmt.Println("after: ")
    printListNode(reverseKGroup(l3, 3)) // [3 -> 2 -> 1] -> [6 -> 5 -> 4] -> [9 -> 8 -> 7] -> [12 -> 11 -> 10] -> [15 -> 14 -> 13]

    l11 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("before: ")
    printListNode(l11) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("after: ")
    printListNode(reverseKGroup1(l11, 2)) // [2,1,4,3,5] [2 -> 1] -> [4 -> 3] -> 5

    l12 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("before: ")
    printListNode(l12) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("after: ")
    printListNode(reverseKGroup1(l12, 3)) // [3,2,1,4,5] [3 -> 2 -> 1] -> 4 -> 5

    l13 := makeListNode([]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15})
    fmt.Println("before: ")
    printListNode(l13) // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10 -> 11 -> 12 -> 13 -> 14 -> 15
    fmt.Println("after: ")
    printListNode(reverseKGroup1(l13, 3)) // [3 -> 2 -> 1] -> [6 -> 5 -> 4] -> [9 -> 8 -> 7] -> [12 -> 11 -> 10] -> [15 -> 14 -> 13]

}