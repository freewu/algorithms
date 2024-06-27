package main

// LCR 136. 删除链表的节点
// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// 返回删除后的链表的头节点。

// 示例 1:
// 输入: head = [4,5,1,9], val = 5
// 输出: [4,1,9]
// 解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.

// 示例 2:
// 输入: head = [4,5,1,9], val = 1
// 输出: [4,5,9]
// 解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

// 说明：
//     题目保证链表中节点的值互不相同
//     若使用 C 或 C++ 语言，你不需要 free 或 delete 被删除的节点

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
// 快慢指针
func deleteNode(head *ListNode, val int) *ListNode {
    if head == nil{
        return head
    }
    // 思路:双指针
    slow := head
    for head.Val == val && head.Next != nil {
        head = head.Next
        slow.Next = nil
        slow = head
    }
    if slow.Next == nil {
        if slow.Val == val {
            return nil
        } else {
            return head
        }
    }
    fast := slow.Next
    for {
        if fast.Val == val {
            slow.Next = nil
            if fast.Next == nil {
                break
            }
            fast = fast.Next
            continue
        }
        slow.Next = fast
        slow = slow.Next
        if fast.Next == nil {
            break
        }
        fast = fast.Next
    }
    return head
}

func deleteNode1(head *ListNode, val int) *ListNode {
    dummy := &ListNode{0, head}
    cur, pre := head, dummy
    for cur != nil {
        if cur.Val == val {
            tmp := cur.Next
            pre.Next = tmp
            cur.Next = nil
            cur = tmp
        } else {
            pre = cur
            cur = cur.Next
        }
    }
    return dummy.Next
}

func deleteNode2(head *ListNode, val int) *ListNode {
    if head.Val == val {
        return head.Next
    }
    dummy := &ListNode{Next: head}
    pre, cur := head, head
    for cur != nil {
        if cur.Val == val {
            pre.Next = cur.Next
            return dummy.Next
        }
        pre = cur
        cur = cur.Next
    }
    return dummy.Next
}

func main() {
    // 示例 1:
    // 输入: head = [4,5,1,9], val = 5
    // 输出: [4,1,9]
    // 解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
    list1 := makeListNode([]int{4,5,1,9})
    printListNode(list1) // 4 -> 5 -> 1 -> 9
    printListNode(deleteNode(list1, 5)) // 4 -> 1 -> 9
    // 示例 2:
    // 输入: head = [4,5,1,9], val = 1
    // 输出: [4,5,9]
    // 解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.
    list2 := makeListNode([]int{4,5,1,9})
    printListNode(list2) // 4 -> 5 -> 1 -> 9
    printListNode(deleteNode(list2, 1)) // 4 -> 5 -> 9

    list11 := makeListNode([]int{4,5,1,9})
    printListNode(list11) // 4 -> 5 -> 1 -> 9
    printListNode(deleteNode1(list11, 5)) // 4 -> 1 -> 9
    list12 := makeListNode([]int{4,5,1,9}) 
    printListNode(list12) // 4 -> 5 -> 1 -> 9
    printListNode(deleteNode1(list12, 1)) // 4 -> 5 -> 9

    list21 := makeListNode([]int{4,5,1,9})
    printListNode(list21) // 4 -> 5 -> 1 -> 9
    printListNode(deleteNode2(list21, 5)) // 4 -> 1 -> 9
    list22 := makeListNode([]int{4,5,1,9}) 
    printListNode(list22) // 4 -> 5 -> 1 -> 9
    printListNode(deleteNode2(list22, 1)) // 4 -> 5 -> 9
}
