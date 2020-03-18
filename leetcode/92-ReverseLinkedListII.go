package main

import "fmt"

/**
Reverse a linked list from position m to n. Do it in one-pass.

Note: 1 ≤ m ≤ n ≤ length of list.

Example:

Input: 1->2->3->4->5->NULL, m = 2, n = 4
Output: 1->4->3->2->5->NULL
 */

type ListNode struct {
    Val  int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
    changeLen := n - m + 1 // 计算需要逆置的节点数
    var preHead *ListNode  // 初始开始逆置的节点前驱
    result := head         // 最终转换后的表头节点
    for head !=nil {
        m--
        if m <= 0 {
            break
        }
        preHead = head
        head = head.Next
    }
    //var modify_list_tail *ListNode
    modifyListTail := head
    var newHead *ListNode
    for head != nil && changeLen > 0 {
        next := head.Next
        head.Next = newHead
        newHead = head
        head = next
        changeLen--
    }
    // 不需要逆置尾部的加
    modifyListTail.Next = head
    if nil != preHead {
        preHead.Next = newHead
    } else {
        result = newHead
    }
    return result
}

// best
func reverseBetween1(head *ListNode, m int, n int) *ListNode {

    if m == n || head == nil || head.Next ==nil  {
        return head
    }


    root := &ListNode{
        Val:-1,
        Next:head,
    }
    // slow := root.Next
    // fast := slow.Next
    slow,fast := head,head.Next

    for i:=0 ; i < n-1 ; i++  {

        if i < m-1 {
            root, slow, fast = root.Next, slow.Next, fast.Next
            continue
        }

        slow.Next = fast.Next
        fast.Next = root.Next
        root.Next = fast
        fast = slow.Next

    }
    if m == 1{
        head = root.Next
    }
    return head
}

func printListNode(a *ListNode) {
    head := a
    for {
        fmt.Print(head.Val," ")
        head = head.Next
        if nil == head {
            break
        }
    }
}

func main() {
    var a, b, c, d, e ListNode
    a.Val = 1
    b.Val = 2
    c.Val = 3
    d.Val = 4
    e.Val = 5
    a.Next = &b
    b.Next = &c
    c.Next = &d
    d.Next = &e
    e.Next = nil

    printListNode(&a)
    fmt.Println()
    printListNode(reverseBetween(&a,2,4))
    fmt.Println()
    printListNode(reverseBetween(&a,1,4))
}
