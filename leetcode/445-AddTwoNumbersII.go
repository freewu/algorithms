package main

// 445. Add Two Numbers II
// You are given two non-empty linked lists representing two non-negative integers.
// The most significant digit comes first and each of their nodes contains a single digit. 
// Add the two numbers and return the sum as a linked list.

// You may assume the two numbers do not contain any leading zero, except the number 0 itself.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/sumii-linked-list.jpg" />
// Input: l1 = [7,2,4,3], l2 = [5,6,4]
// Output: [7,8,0,7]

// Example 2:
// Input: l1 = [2,4,3], l2 = [5,6,4]
// Output: [8,0,7]

// Example 3:
// Input: l1 = [0], l2 = [0]
// Output: [0]
 
// Constraints:
//     The number of nodes in each linked list is in the range [1, 100].
//     0 <= Node.val <= 9
//     It is guaranteed that the list represents a number that does not have leading zeros.

// Follow up: Could you solve it without reversing the input lists?

import "fmt"

type ListNode struct {
    Val  int
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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    getVal := func (l *ListNode) int { if l == nil { return 0; }; return l.Val;}
    nodeNext := func (l *ListNode) *ListNode { if l == nil {  return nil; }; return l.Next; }
    reverse := func (head *ListNode) *ListNode {
        var resp *ListNode
        for head != nil {
            old := resp
            resp, head = head, head.Next
            resp.Next = old
        }
        
        return resp
    }
    l1 = reverse(l1)
    l2 = reverse(l2)

    res := &ListNode{}
    head := res
    memo := 0
    for l1 != nil || l2 != nil {
        head.Next = &ListNode{
            Val: getVal(l1) + getVal(l2) + memo,
        }

        l1 = nodeNext(l1)
        l2 = nodeNext(l2)
        memo = 0
        head = head.Next
        if head.Val > 9 {
            memo = 1
            head.Val -= 10
        }
    }

    if memo > 0 {
        head.Next = &ListNode{
            Val: 1,
        }
    }
    return reverse(res.Next)
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
    var prv *ListNode
    for l1 != nil {
        next := l1.Next
        l1.Next = prv
        prv = l1
        l1 = next
    }
    l1 = prv
    prv = nil
    for l2 != nil {
        next := l2.Next
        l2.Next = prv
        prv = l2
        l2 = next
    }
    l2 = prv
    add := 0
    dummy := &ListNode{}
    cur := dummy
    for l1 != nil || l2 != nil {
        res := add
        if l1 != nil {
            res += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            res += l2.Val
            l2 = l2.Next
        }
        add = 0
        if res >= 10 {
            add = res / 10
            res = res % 10
        }
        cur.Next = &ListNode{
            Val: res,
        }
        cur = cur.Next
    }
    if add > 0 {
        cur.Next = &ListNode{
            Val: add,
        }
        cur = cur.Next
    }
    prv = nil
    cur = dummy.Next
    for cur != nil {
        next := cur.Next
        cur.Next = prv
        prv = cur
        cur = next
    }
    return prv
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
    // 遍历两个list 获得数组
    // 数组计算
    // 通过数组重新生成链表
    return makeListNode([]int{})
}

func main() {
    l11 := makeListNode([]int{7,2,4,3})
    l12 := makeListNode([]int{5,6,4})
    fmt.Println("7243 + 564 = 7807")
    printListNode(l11) // 7 -> 2 -> 4 -> 3
    printListNode(l12) // 5 -> 6 -> 4
    printListNode(addTwoNumbers(l11,l12)) // 7 -> 8 -> 0 -> 7

    l21 := makeListNode([]int{2,4,3})
    l22 := makeListNode([]int{5,6,4})
    fmt.Println("243 + 564 = 807")
    printListNode(l21) // 2 -> 4 -> 3
    printListNode(l22) // 5 -> 6 -> 4
    printListNode(addTwoNumbers(l21,l22)) // 8 -> 0 -> 7

    l31 := makeListNode([]int{0})
    l32 := makeListNode([]int{0})
    fmt.Println("0 + 0 = 0")
    printListNode(l31) // 0
    printListNode(l32) // 0
    printListNode(addTwoNumbers(l31,l32)) // 0

    l111 := makeListNode([]int{7,2,4,3})
    l112 := makeListNode([]int{5,6,4})
    fmt.Println("7243 + 564 = 7807")
    printListNode(l111) // 7 -> 2 -> 4 -> 3
    printListNode(l112) // 5 -> 6 -> 4
    printListNode(addTwoNumbers1(l111,l112)) // 7 -> 8 -> 0 -> 7

    l121 := makeListNode([]int{2,4,3})
    l122 := makeListNode([]int{5,6,4})
    fmt.Println("243 + 564 = 807")
    printListNode(l121) // 2 -> 4 -> 3
    printListNode(l122) // 5 -> 6 -> 4
    printListNode(addTwoNumbers1(l121,l122)) // 8 -> 0 -> 7

    l131 := makeListNode([]int{0})
    l132 := makeListNode([]int{0})
    fmt.Println("0 + 0 = 0")
    printListNode(l131) // 0
    printListNode(l132) // 0
    printListNode(addTwoNumbers1(l131,l132)) // 0
}