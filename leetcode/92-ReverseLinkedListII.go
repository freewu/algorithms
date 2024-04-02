package main

// 92. Reverse Linked List II
// Given the head of a singly linked list and two integers left and right where left <= right, 
// reverse the nodes of the list from position left to position right, and return the reversed list.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/rev2ex2.jpg" />
// Input: head = [1,2,3,4,5], left = 2, right = 4
// Output: [1,4,3,2,5]

// Example 2:
// Input: head = [5], left = 1, right = 1
// Output: [5]
 
// Constraints:
//     The number of nodes in the list is n.
//     1 <= n <= 500
//     -500 <= Node.val <= 500
//     1 <= left <= right <= n

// Follow up: Could you do it in one pass?

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

// best 快慢指针
func reverseBetween1(head *ListNode, m int, n int) *ListNode {
    if m == n || head == nil || head.Next ==nil  {
        return head
    }
    root := &ListNode {
        Val:-1,
        Next:head,
    }
    slow,fast := head, head.Next

    for i := 0; i < n-1 ; i++  {
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

func reverseBetween2(head *ListNode, left int, right int) *ListNode {
    tl := &ListNode{}
    tl.Next = head
    pre := tl
    for i := 0 ; i < left - 1; i++ {
        pre = pre.Next
    }
    cur := pre.Next
    for i := 0; i < right - left; i++ {
        t := cur.Next
        cur.Next = t.Next
        t.Next = pre.Next
        pre.Next = t
    }
    return tl.Next
}

func main() {
    l11 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("before: ")
    printListNode(l11) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("after: ")
    printListNode(reverseBetween(l11,2,4)) // 1 -> 4 -> 3 -> 2 -> 5

    l12 := makeListNode([]int{5})
    fmt.Println("before: ")
    printListNode(l12) // 5
    fmt.Println("after: ")
    printListNode(reverseBetween(l12,1,1)) // 5

    l21 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("before: ")
    printListNode(l21) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("after: ")
    printListNode(reverseBetween1(l21,2,4))  // 1 -> 4 -> 3 -> 2 -> 5

    l22 := makeListNode([]int{5})
    fmt.Println("before: ")
    printListNode(l22) // 5
    fmt.Println("after: ")
    printListNode(reverseBetween1(l22,1,1)) // 5

    l31 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("before: ")
    printListNode(l31) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println("after: ")
    printListNode(reverseBetween2(l31,2,4))  // 1 -> 4 -> 3 -> 2 -> 5

    l32 := makeListNode([]int{5})
    fmt.Println("before: ")
    printListNode(l32) // 5
    fmt.Println("after: ")
    printListNode(reverseBetween2(l32,1,1)) // 5
}
