package main

// LCR 025. 两数相加 II
// 给定两个 非空链表 l1和 l2 来代表两个非负整数。数字最高位位于链表开始位置。
// 它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

// 可以假设除了数字 0 之外，这两个数字都不会以零开头。

// 示例1：
// <img src="https://pic.leetcode-cn.com/1626420025-fZfzMX-image.png" />
// 输入：l1 = [7,2,4,3], l2 = [5,6,4]
// 输出：[7,8,0,7]

// 示例2：
// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[8,0,7]

// 示例3：
// 输入：l1 = [0], l2 = [0]
// 输出：[0]
 
// 提示：
//     链表的长度范围为 [1, 100]
//     0 <= node.val <= 9
//     输入数据保证链表代表的数字无前导 0
 
// 进阶：如果输入链表不能修改该如何处理？换句话说，不能对列表中的节点进行翻转。

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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    getVal := func (l *ListNode) int { if l == nil { return 0; }; return l.Val;}
    nodeNext := func (l *ListNode) *ListNode { if l == nil {  return nil; }; return l.Next; }
    reverse := func (head *ListNode) *ListNode {
        var resp *ListNode
        for head != nil {
            old := resp
            resp, head = head, head.Next
            resp.Next = old
        }
        
        return resp
    }
    l1, l2 = reverse(l1), reverse(l2)
    res := &ListNode{}
    head, memo := res, 0
    for l1 != nil || l2 != nil {
        head.Next = &ListNode{
            Val: getVal(l1) + getVal(l2) + memo,
        }
        l1, l2 = nodeNext(l1), nodeNext(l2)
        memo = 0
        head = head.Next
        if head.Val > 9 {
            memo = 1
            head.Val -= 10
        }
    }
    if memo > 0 {
        head.Next = &ListNode{
            Val: 1,
        }
    }
    return reverse(res.Next)
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
    var prv *ListNode
    for l1 != nil {
        next := l1.Next
        l1.Next = prv
        prv = l1
        l1 = next
    }
    l1 = prv
    prv = nil
    for l2 != nil {
        next := l2.Next
        l2.Next = prv
        prv = l2
        l2 = next
    }
    l2 = prv
    add := 0
    dummy := &ListNode{}
    cur := dummy
    for l1 != nil || l2 != nil {
        res := add
        if l1 != nil {
            res += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            res += l2.Val
            l2 = l2.Next
        }
        add = 0
        if res >= 10 {
            add = res / 10
            res = res % 10
        }
        cur.Next = &ListNode{
            Val: res,
        }
        cur = cur.Next
    }
    if add > 0 {
        cur.Next = &ListNode{
            Val: add,
        }
        cur = cur.Next
    }
    prv = nil
    cur = dummy.Next
    for cur != nil {
        next := cur.Next
        cur.Next = prv
        prv = cur
        cur = next
    }
    return prv
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
    reserve := func(node *ListNode) *ListNode {
        var prev *ListNode
        cur:=node
        for cur !=nil {
            next := cur.Next
            cur.Next = prev
            prev = cur
            cur = next
        }
        return prev
    }
    res :=&ListNode{} 
    list1, list2 := reserve(l1), reserve(l2) // 先反转
    // 3427 
    // 465
    carry := 0
    cur := res
    for list1 != nil || list2 != nil || carry !=0 {
        sum := carry
        if list1 != nil {
            sum +=list1.Val
            list1 = list1.Next
        }
        if list2 !=nil {
            sum += list2.Val
            list2 = list2.Next
        }
        cur.Next = &ListNode{ Val: sum % 10 }
        cur = cur.Next
        carry = sum / 10 // 进位
    }
    return reserve(res.Next)
}

func main() {
    l11 := makeListNode([]int{7,2,4,3})
    l12 := makeListNode([]int{5,6,4})
    fmt.Println("7243 + 564 = 7807")
    printListNode(l11) // 7 -> 2 -> 4 -> 3
    printListNode(l12) // 5 -> 6 -> 4
    printListNode(addTwoNumbers(l11,l12)) // 7 -> 8 -> 0 -> 7

    l21 := makeListNode([]int{2,4,3})
    l22 := makeListNode([]int{5,6,4})
    fmt.Println("243 + 564 = 807")
    printListNode(l21) // 2 -> 4 -> 3
    printListNode(l22) // 5 -> 6 -> 4
    printListNode(addTwoNumbers(l21,l22)) // 8 -> 0 -> 7

    l31 := makeListNode([]int{0})
    l32 := makeListNode([]int{0})
    fmt.Println("0 + 0 = 0")
    printListNode(l31) // 0
    printListNode(l32) // 0
    printListNode(addTwoNumbers(l31,l32)) // 0

    l111 := makeListNode([]int{7,2,4,3})
    l112 := makeListNode([]int{5,6,4})
    fmt.Println("7243 + 564 = 7807")
    printListNode(l111) // 7 -> 2 -> 4 -> 3
    printListNode(l112) // 5 -> 6 -> 4
    printListNode(addTwoNumbers1(l111,l112)) // 7 -> 8 -> 0 -> 7

    l121 := makeListNode([]int{2,4,3})
    l122 := makeListNode([]int{5,6,4})
    fmt.Println("243 + 564 = 807")
    printListNode(l121) // 2 -> 4 -> 3
    printListNode(l122) // 5 -> 6 -> 4
    printListNode(addTwoNumbers1(l121,l122)) // 8 -> 0 -> 7

    l131 := makeListNode([]int{0})
    l132 := makeListNode([]int{0})
    fmt.Println("0 + 0 = 0")
    printListNode(l131) // 0
    printListNode(l132) // 0
    printListNode(addTwoNumbers1(l131,l132)) // 0

    l211 := makeListNode([]int{7,2,4,3})
    l212 := makeListNode([]int{5,6,4})
    fmt.Println("7243 + 564 = 7807")
    printListNode(l211) // 7 -> 2 -> 4 -> 3
    printListNode(l112) // 5 -> 6 -> 4
    printListNode(addTwoNumbers2(l211,l212)) // 7 -> 8 -> 0 -> 7

    l221 := makeListNode([]int{2,4,3})
    l222 := makeListNode([]int{5,6,4})
    fmt.Println("243 + 564 = 807")
    printListNode(l221) // 2 -> 4 -> 3
    printListNode(l222) // 5 -> 6 -> 4
    printListNode(addTwoNumbers2(l221,l222)) // 8 -> 0 -> 7

    l231 := makeListNode([]int{0})
    l232 := makeListNode([]int{0})
    fmt.Println("0 + 0 = 0")
    printListNode(l231) // 0
    printListNode(l232) // 0
    printListNode(addTwoNumbers2(l231,l232)) // 0
}