package main

// LCR 171. 训练计划 V
// 某教练同时带教两位学员，分别以链表 l1、l2 记录了两套核心肌群训练计划，节点值为训练项目编号。
// 两套计划仅有前半部分热身项目不同，后续正式训练项目相同。
// 请设计一个程序找出并返回第一个正式训练项目编号。如果两个链表不存在相交节点，返回 null 。

// 如下面的两个链表：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_statement.png" />
// 在节点 c1 开始相交。

// 输入说明：
//     intersectVal - 相交的起始节点的值。如果不存在相交节点，这一值为 0
//     l1 - 第一个训练计划链表
//     l2 - 第二个训练计划链表
//     skip1 - 在 l1 中（从头节点开始）跳到交叉节点的节点数
//     skip2 - 在 l2 中（从头节点开始）跳到交叉节点的节点数

// 程序将根据这些输入创建链式数据结构，并将两个头节点 head1 和 head2 传递给你的程序。
// 如果程序能够正确返回相交节点，那么你的解决方案将被视作正确答案 。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2018/12/13/160_example_1.png" />
// 输入：intersectVal = 8, listA = [4,1,8,4,5], listB = [5,0,1,8,4,5], skipA = 2, skipB = 3
// 输出：Reference of the node with value = 8
// 解释：第一个正式训练项目编号为 8 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [4,1,8,4,5]，链表 B 为 [5,0,1,8,4,5]。在 A 中，相交节点前有 2 个节点；在 B 中，相交节点前有 3 个节点。

// 示例 2：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_2.png" />
// 输入：intersectVal = 2, listA = [0,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
// 输出：Reference of the node with value = 2
// 解释：第一个正式训练项目编号为 2 （注意，如果两个列表相交则不能为 0）。从各自的表头开始算起，链表 A 为 [0,9,1,2,4]，链表 B 为 [3,2,4]。在 A 中，相交节点前有 3 个节点；在 B 中，相交节点前有 1 个节点。

// 示例 3：
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/160_example_3.png" />
// 输入：intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
// 输出：null
// 解释：两套计划完全不同，返回 null。从各自的表头开始算起，链表 A 为 [2,6,4]，链表 B 为 [1,5]。由于这两个链表不相交，所以 intersectVal 必须为 0，而 skipA 和 skipB 可以是任意值。

// 注意：
//     如果两个链表没有交点，返回 null.
//     在返回结果后，两个链表仍须保持原有的结构。
//     可假定整个链表结构中没有循环。
//     程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

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