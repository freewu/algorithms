package main

// 面试题 02.03. Delete Middle Node LCCI
// Implement an algorithm to delete a node in the middle 
// (i.e., any node but the first and last node, not necessarily the exact middle) of a singly linked list, 
// given only access to that node.

// Example:
// Input: the node c from the linked list a->b->c->d->e->f
// Output: nothing is returned, but the new linked list looks like a->b->d->e->f

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
func deleteNode(node *ListNode) {
    *node = *node.Next
}

func deleteNode1(node *ListNode) {
    if node == nil { return }
    cur := node
    var prev *ListNode
    for cur.Next != nil {
        cur.Val = cur.Next.Val
        prev = cur
        cur = cur.Next
    }
    prev.Next = nil
}

func main() {
    l1 := makeListNode([]int{4,5,1,9})
    lh1 := l1
    printListNode(l1) // 4 -> 5 -> 1 -> 9
    for nil != l1.Next {
        if l1.Val == 5 {
            deleteNode(l1)
        }
        l1 = l1.Next
    }
    printListNode(lh1) // 4 -> 1 -> 9


    l2 := makeListNode([]int{4,5,1,9})
    lh2 := l2
    printListNode(l2) // 4 -> 5 -> 1 -> 9
    // for nil != l2.Next{
    //     if l2.Val == 1 {
    //         deleteNode(l2)
    //     }
    //     l2 = l2.Next
    // }
    printListNode(lh2) // 4 -> 5 -> 9
}