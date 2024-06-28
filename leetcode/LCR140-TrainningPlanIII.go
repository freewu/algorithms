package main

// LCR 140. 训练计划 II
// 给定一个头节点为 head 的链表用于记录一系列核心肌群训练项目编号，请查找并返回倒数第 cnt 个训练项目编号。

// 示例 1：
// 输入：head = [2,4,7,8], cnt = 1
// 输出：8

// 提示：
//     1 <= head.length <= 100
//     0 <= head[i] <= 100
//     1 <= cnt <= head.length

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
func trainingPlan(head *ListNode, cnt int) *ListNode {
    fast, slow := head, head
    for fast != nil && cnt > 0 { // fast走到了cnt步
        fast = fast.Next
        cnt--
    }
    for fast != nil {
        fast = fast.Next
        slow = slow.Next
    }
    return slow
}

func main() {
    // 示例 1：
    // 输入：head = [2,4,7,8], cnt = 1
    // 输出：8
    list1 := makeListNode([]int{2,4,7,8})
    printListNode(list1)
    fmt.Println(trainingPlan(list1, 1)) // &{8 <nil>}

    list2 := makeListNode([]int{2,4,7,8})
    printListNode(list2)
    fmt.Println(trainingPlan(list2, 2)) // &{7 0xc00008a080}
}