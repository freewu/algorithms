package main

// 203. Remove Linked List Elements
// Given the head of a linked list and an integer val, 
// remove all the nodes of the linked list that has Node.val == val, and return the new head.

// Example 1:
// 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6  =>  1 -> 2  -> 3 -> 4 -> 5 
// <img src="https://assets.leetcode.com/uploads/2021/03/06/removelinked-list.jpg" />
// Input: head = [1,2,6,3,4,5,6], val = 6
// Output: [1,2,3,4,5]

// Example 2:
// Input: head = [], val = 1
// Output: []

// Example 3:
// Input: head = [7,7,7,7], val = 7
// Output: []
 
// Constraints:
//     The number of nodes in the list is in the range [0, 104].
//     1 <= Node.val <= 50
//     0 <= val <= 50

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
func removeElements(head *ListNode, val int) *ListNode {
    if nil == head {
        return nil
    }
    pre, tmp := head, head
    for tmp != nil {
        if tmp.Val == val {
            pre.Next = tmp.Next
        } else {
            pre = tmp
        }
        tmp = tmp.Next
    }
    if head.Val == val { // 处理最后一节点值 为 val 的情况
        head = head.Next 
    }
    return head
}

// 快慢指针
func removeElements1(head *ListNode, val int) *ListNode {
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

func removeElements2(head *ListNode, val int) *ListNode {
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

func main() {
    // Example 1:
    // 1 -> 2 -> 6 -> 3 -> 4 -> 5 -> 6  =>  1 -> 2  -> 3 -> 4 -> 5 
    // <img src="https://assets.leetcode.com/uploads/2021/03/06/removelinked-list.jpg" />
    // Input: head = [1,2,6,3,4,5,6], val = 6
    // Output: [1,2,3,4,5]
    l1 := makeListNode([]int{1,2,6,3,4,5,6})
    printListNode(l1)
    printListNode(removeElements(l1,6))
    // Example 2:
    // Input: head = [], val = 1
    // Output: []
    l2 := makeListNode([]int{})
    printListNode(l2)
    printListNode(removeElements(l2,1))
    // Example 3:
    // Input: head = [7,7,7,7], val = 7
    // Output: []
    l3 := makeListNode([]int{7,7,7,7})
    printListNode(l3)
    printListNode(removeElements(l3,7))

    l11 := makeListNode([]int{1,2,6,3,4,5,6})
    printListNode(l11)
    printListNode(removeElements(l11,6))
    l12 := makeListNode([]int{})
    printListNode(l12)
    printListNode(removeElements(l12,1))
    l13 := makeListNode([]int{7,7,7,7})
    printListNode(l13)
    printListNode(removeElements(l13,7))

    l21 := makeListNode([]int{1,2,6,3,4,5,6})
    printListNode(l21)
    printListNode(removeElements(l21,6))
    l22 := makeListNode([]int{})
    printListNode(l22)
    printListNode(removeElements(l22,1))
    l23 := makeListNode([]int{7,7,7,7})
    printListNode(l23)
    printListNode(removeElements(l23,7))
}