package main

// LCR 023. 相交链表
// 给定两个单链表的头节点 headA 和 headB ，请找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
// 图示两个链表在节点 c1 开始相交：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_statement.png" />

// 题目数据 保证 整个链式结构中不存在环。
// 注意，函数返回结果后，链表必须 保持其原始结构 。


// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2018/12/13/160_example_1.png" />
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
// 输出：Intersected at '8'
// 解释：相交节点的值为 8 （注意，如果两个链表相交则不能为 0）。
// 从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。
// 在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2018/12/13/160_example_2.png" />
// 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// 输出：Intersected at '2'
// 解释：相交节点的值为 2 （注意，如果两个链表相交则不能为 0）。
// 从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。
// 在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

// 示例 3：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_3.png" />
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null
// 解释：从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。
// 由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。
// 这两个链表不相交，因此返回 null 。
 
// 提示：
//     listA 中节点数目为 m
//     listB 中节点数目为 n
//     0 <= m, n <= 3 * 10^4
//     1 <= Node.val <= 10^5
//     0 <= skipA <= m
//     0 <= skipB <= n
//     如果 listA 和 listB 没有交点，intersectVal 为 0
//     如果 listA 和 listB 有交点，intersectVal == listA[skipA + 1] == listB[skipB + 1]

// 进阶：能否设计一个时间复杂度 O(n) 、仅用 O(1) 内存的解决方案？

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
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    a, b := headA, headB
    // AB不相交时，AB走的路径都是两条链表。
    // AB相交时，转换为环形链表的入环点。
    for a != b {
        if a != nil { a = a.Next } else { a = headB }
        if b != nil { b = b.Next } else { b = headA }
    }
    return a
 }

func main() {
    fmt.Println("Example 1:")
    l11 := makeListNode([]int{4,1,8,4,5})
    l12 := makeListNode([]int{5,6,1,8,4,5})
    printListNode(l11)
    printListNode(l12)
    printListNode(getIntersectionNode(l11,l12))

    fmt.Println("Example 2:")
    l21 := makeListNode([]int{1,9,1,2,4})
    l22 := makeListNode([]int{3,2,4})
    printListNode(l21)
    printListNode(l22)
    printListNode(getIntersectionNode(l21,l22))

    fmt.Println("Example 3:")
    l31 := makeListNode([]int{2,6,4})
    l32 := makeListNode([]int{1,5})
    printListNode(l31)
    printListNode(l32)
    printListNode(getIntersectionNode(l31, l32))
}