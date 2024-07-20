package main

// LCR 077. 排序链表
// 给定链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/09/14/sort_list_1.jpg" />
// 输入：head = [4,2,1,3]
// 输出：[1,2,3,4]

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2020/09/14/sort_list_2.jpg" />
// 输入：head = [-1,5,3,4,0]
// 输出：[-1,0,3,4,5]

// 示例 3：
// 输入：head = []
// 输出：[]

// 提示：
//     链表中节点的数目在范围 [0, 5 * 10^4] 内
//     -10^5 <= Node.val <= 10^5
 
// 进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

import "fmt"
import "slices"

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
func sortList(head *ListNode) *ListNode {
    curr := head
    arr := []*ListNode{}
    // 添加到一个 slice 中
    for curr != nil {
        arr = append(arr, curr)
        curr = curr.Next
    }
    // 使用 slices 排序
    slices.SortFunc(arr, func(l1, l2 *ListNode) int {
        return l1.Val - l2.Val
    })
    // 通过 arr 重链
    dummy := &ListNode{}
    curr = dummy
    for _, node := range arr {
        curr.Next = node
        curr = node
    }
    curr.Next = nil
    return dummy.Next
}

func sortList1(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    ma, mi := head.Val, head.Val
    t := head.Next
    for t != nil {
        if t.Val > ma {
            ma = t.Val
        }
        if t.Val < mi {
            mi = t.Val
        }
        t = t.Next
    }
    counts := make([]uint, ma - mi + 1)
    t = head
    for t != nil {
        counts[t.Val - mi] += 1
        t = t.Next
    }
    t = head
    for i, v := range counts {
        for v > 0 {
            t.Val = mi + i
            t = t.Next
            v--
        }
    }
    return head
}

func main() {
    l1 :=  makeListNode([]int{4,2,1,3})
    fmt.Println("Before sortList: ")
    printListNode(l1) // 4 -> 2 -> 1 -> 3
    fmt.Println("After sortList: ")
    printListNode(sortList(l1)) // 1 -> 2 -> 3 -> 4

    l2 :=  makeListNode([]int{-1,5,3,4,0})
    fmt.Println("Before sortList: ")
    printListNode(l2) // -1 -> 5 -> 3 -> 4 -> 0
    fmt.Println("After sortList: ")
    printListNode(sortList(l2)) // -1 -> 0 -> 3 -> 4 -> 5

    l3 :=  makeListNode([]int{})
    fmt.Println("Before sortList: ")
    printListNode(l3) // 
    fmt.Println("After sortList: ")
    printListNode(sortList(l3)) // 

    l11 :=  makeListNode([]int{4,2,1,3})
    fmt.Println("Before sortList: ")
    printListNode(l11) // 4 -> 2 -> 1 -> 3
    fmt.Println("After sortList: ")
    printListNode(sortList1(l11)) // 1 -> 2 -> 3 -> 4

    l12 :=  makeListNode([]int{-1,5,3,4,0})
    fmt.Println("Before sortList: ")
    printListNode(l12) // -1 -> 5 -> 3 -> 4 -> 0
    fmt.Println("After sortList: ")
    printListNode(sortList1(l12)) // -1 -> 0 -> 3 -> 4 -> 5

    l13 :=  makeListNode([]int{})
    fmt.Println("Before sortList: ")
    printListNode(l13) // 
    fmt.Println("After sortList: ")
    printListNode(sortList1(l13)) // 
}