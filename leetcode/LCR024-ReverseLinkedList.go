package main

// LCR 024. 反转链表
// 给定单链表的头节点 head ，请反转链表，并返回反转后的链表的头节点。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2021/02/19/rev1ex1.jpg" />
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2021/02/19/rev1ex2.jpg" />
// 输入：head = [1,2]
// 输出：[2,1]

// 示例 3：
// 输入：head = []
// 输出：[]

// 提示：
//     链表中节点的数目范围是 [0, 5000]
//     -5000 <= Node.val <= 5000
 
// 进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？

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
func reverseList(head *ListNode) *ListNode {
    var res *ListNode
    for head != nil {
        next := head.Next  // 备份head.Next
        head.Next = res // 更新  head.Next
        res = head      // 移动 new_head
        head = next
        //fmt.Println("new_head: ",new_head,"head: ",head)
    }
    return res
}

// best
func reverseList1(head *ListNode) *ListNode {
    var prev *ListNode
    cur := head
    var next *ListNode
    for cur != nil {
        next = cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}

func main() {
    l1 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
    printListNode(l1)
    fmt.Println("reverseList(l1): ")
    printListNode(reverseList(l1))

    l2 := makeListNode([]int{1,2})
    fmt.Println("l2: ")
    printListNode(l2)
    fmt.Println("reverseList(l2): ")
    printListNode(reverseList(l2))

    l3 := makeListNode([]int{})
    fmt.Println("l3: ")
    printListNode(l3)
    fmt.Println("reverseList(l3): ")
    printListNode(reverseList(l3))


    l1 = makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
    printListNode(l1)
    fmt.Println("reverseList1(l1): ")
    printListNode(reverseList1(l1))

    l2 = makeListNode([]int{1,2})
    fmt.Println("l2: ")
    printListNode(l2)
    fmt.Println("reverseList1(l2): ")
    printListNode(reverseList1(l2))

    l3 = makeListNode([]int{})
    fmt.Println("l3: ")
    printListNode(l3)
    fmt.Println("reverseList1(l3): ")
    printListNode(reverseList1(l3))
}