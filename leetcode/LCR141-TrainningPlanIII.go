package main

// LCR 141. 训练计划 III
// 给定一个头节点为 head 的单链表用于记录一系列核心肌群训练编号，请将该系列训练编号 倒序 记录于链表并返回。

// 示例 1：
// 输入：head = [1,2,3,4,5]
// 输出：[5,4,3,2,1]

// 示例 2：
// 输入：head = [1,2]
// 输出：[2,1]

// 示例 3：
// 输入：head = []
// 输出：[]

// 提示：
//     链表中节点的数目范围是 [0, 5000]
//     -5000 <= Node.val <= 5000

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
func trainningPlan(head *ListNode) *ListNode {
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
func trainningPlan1(head *ListNode) *ListNode {
    var prev *ListNode = nil
    cur := head
    var next *ListNode = nil

    for cur != nil {
        next = cur.Next
        cur.Next = prev
        prev = cur
        cur = next
    }
    return prev
}

func trainningPlan2(head *ListNode) *ListNode {
    var recur func(q *ListNode, w *ListNode) *ListNode
    recur = func (q *ListNode, w *ListNode) *ListNode {
        if q == nil {
            return w
        }
        res := recur(q.Next, q)
        q.Next = w
        return res
    }
    return recur(head,nil)
}


func main() {
    l1 := makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
    printListNode(l1)
    fmt.Println("trainningPlan(l1): ")
    printListNode(trainningPlan(l1))

    l2 := makeListNode([]int{1,2})
    fmt.Println("l2: ")
    printListNode(l2)
    fmt.Println("trainningPlan(l2): ")
    printListNode(trainningPlan(l2))

    l3 := makeListNode([]int{})
    fmt.Println("l3: ")
    printListNode(l3)
    fmt.Println("trainningPlan(l3): ")
    printListNode(trainningPlan(l3))


    l1 = makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
    printListNode(l1)
    fmt.Println("trainningPlan1(l1): ")
    printListNode(trainningPlan1(l1))

    l2 = makeListNode([]int{1,2})
    fmt.Println("l2: ")
    printListNode(l2)
    fmt.Println("trainningPlan1(l2): ")
    printListNode(trainningPlan1(l2))

    l3 = makeListNode([]int{})
    fmt.Println("l3: ")
    printListNode(l3)
    fmt.Println("trainningPlan1(l3): ")
    printListNode(trainningPlan1(l3))

    l1 = makeListNode([]int{1,2,3,4,5})
    fmt.Println("l1: ")
    printListNode(l1)
    fmt.Println("trainningPlan2(l1): ")
    printListNode(trainningPlan2(l1))

    l2 = makeListNode([]int{1,2})
    fmt.Println("l2: ")
    printListNode(l2)
    fmt.Println("trainningPlan2(l2): ")
    printListNode(trainningPlan2(l2))

    l3 = makeListNode([]int{})
    fmt.Println("l3: ")
    printListNode(l3)
    fmt.Println("trainningPlan2(l3): ")
    printListNode(trainningPlan2(l3))
}
