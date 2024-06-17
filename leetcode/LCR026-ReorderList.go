package main

// LCR 026. 重排链表
// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//     L0 → L1 → … → Ln-1 → Ln 

// 请将其重新排列后变为：

//     L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …

// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

// 示例 1:
// <img src="https://pic.leetcode-cn.com/1626420311-PkUiGI-image.png" />
// 输入: head = [1,2,3,4]
// 输出: [1,4,2,3]

// 示例 2:
// <img src="https://pic.leetcode-cn.com/1626420320-YUiulT-image.png" />
// 输入: head = [1,2,3,4,5]
// 输出: [1,5,2,4,3]

// 提示：
//     链表的长度范围为 [1, 5 * 10^4]
//     1 <= node.val <= 1000


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
func reorderList(head *ListNode) {
    if head.Next == nil { // 只有一个结点,直接返回
        return
    }
    middleNode := func (head *ListNode) *ListNode { // 查找中间结点
        slow, fast := head, head // 快慢指针
        for fast != nil && fast.Next != nil {
            slow = slow.Next
            fast = fast.Next.Next
        }
        return slow
    }
    reverse := func (l *ListNode) *ListNode { // 反转链表（不带头结点）
        if l == nil {
            return nil
        }
        head := &ListNode{} // 空头结点
        p := l
        for p != nil {
            l = l.Next
            p.Next = head.Next
            head.Next = p
            p = l
        }
        return head.Next
    }

    mid := middleNode(head)
    back := mid.Next // 后半段链表
    mid.Next = nil // 断开前后
    front := head.Next   // 前半段链表
    back = reverse(back) // 后半段链表反转

    // 交替合并前后链表
    p := head
    for front != nil && back != nil {
        p.Next = back
        back = back.Next
        p = p.Next

        p.Next = front
        front = front.Next
        p = p.Next
    }
    if front != nil {
        p.Next = front
    }
    if back != nil {
        p.Next = back
    }
}

func reorderList1(head *ListNode)  {
    middleNode := func(head *ListNode) *ListNode {
        slow, fast := head, head
        for fast != nil && fast.Next != nil {
            slow = slow.Next
            fast = fast.Next.Next
        }
        return slow
    }
    reverseList := func(head *ListNode) *ListNode {
        var pre, cur *ListNode = nil, head
        for cur != nil {
            nxt := cur.Next
            cur.Next = pre
            pre = cur
            cur = nxt
        }
        return pre
    }
    m := middleNode(head)
    head2 := reverseList(m)
    for head2.Next != nil {
        nxt := head.Next
        nxt2 := head2.Next
        head.Next = head2
        head2.Next = nxt
        head = nxt
        head2 = nxt2
    }
}

func main() {
    l1 :=  makeListNode([]int{1,2,3,4})
    printListNode(l1) // (1) -> (2) -> (3) -> (4)
    reorderList(l1)
    fmt.Println("After reorder: ")
    printListNode(l1) // (1) -> (4) -> (2) -> (3)
    
    l2 :=  makeListNode([]int{1,2,3,4,5})
    printListNode(l2) // (1) -> (2) -> (3) -> (4) -> (5)
    reorderList(l2)
    fmt.Println("After reorder: ")
    printListNode(l2) // (1) -> (5) -> (2) -> (4) -> (3)

    l11 :=  makeListNode([]int{1,2,3,4})
    printListNode(l11) // (1) -> (2) -> (3) -> (4)
    reorderList1(l11)
    fmt.Println("After reorder: ")
    printListNode(l11) // (1) -> (4) -> (2) -> (3)
    
    l12 :=  makeListNode([]int{1,2,3,4,5})
    printListNode(l12) // (1) -> (2) -> (3) -> (4) -> (5)
    reorderList1(l12)
    fmt.Println("After reorder: ")
    printListNode(l12) // (1) -> (5) -> (2) -> (4) -> (3)
}