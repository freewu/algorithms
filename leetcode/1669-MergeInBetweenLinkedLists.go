package main

// 1669. Merge In Between Linked Lists
// You are given two linked lists: list1 and list2 of sizes n and m respectively.
// Remove list1's nodes from the ath node to the bth node, and put list2 in their place.
// The blue edges and nodes in the following figure indicate the result:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/fig1.png" />
// Build the result list and return its head.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2024/03/01/ll.png" />
// Input: list1 = [10,1,13,6,9,5], a = 3, b = 4, list2 = [1000000,1000001,1000002]
// Output: [10,1,13,1000000,1000001,1000002,5]
// Explanation: We remove the nodes 3 and 4 and put the entire list2 in their place. The blue edges and nodes in the above figure indicate the result.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/11/05/merge_linked_list_ex2.png" />
// Input: list1 = [0,1,2,3,4,5,6], a = 2, b = 5, list2 = [1000000,1000001,1000002,1000003,1000004]
// Output: [0,1,1000000,1000001,1000002,1000003,1000004,6]
// Explanation: The blue edges and nodes in the above figure indicate the result.
 
// Constraints:
//     3 <= list1.length <= 10^4
//     1 <= a <= b < list1.length - 1
//     1 <= list2.length <= 10^4

import "fmt"

type ListNode struct {
    Val int
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
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    if list1 == nil || list2 == nil{
        return nil
    }
    node, temp, pre := list1, list1, list1
    count := 0
    for temp != nil {
        // 需要找到删除的开始位置(也是 list2 的插入点)
        if count == a - 1 {
            pre = temp
        }
        // 到达插入点位置
        if count == a {
            pre.Next = list2
            node = list2
            // 把 list2 的节点插入到 list1 列表中
            for node.Next != nil{
                node = node.Next
            }
        }
        // 到达 list1 删除的结束点
        if count == b {
            // 接上 list1 剩余的,退出即可
            node.Next = temp.Next
            // 这句可以不需要
            // temp.Next = nil
            break
        }
        temp = temp.Next
        count++
    }
    return list1
}

func mergeInBetween1(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
    p, q := list1, list1
    // 取 list1 的前半部分[0 - a)
    for ; a > 1; a-- {
        p = p.Next
    }
    // 取 list1 的后边半部分[b - len(list1))
    for ; b > 0; b-- {
        q = q.Next
    }
    // list1 的前半部分关联上 list2
    p.Next = list2
    for p.Next != nil {
        p = p.Next
    }
    // 关联上 p.Next 的后半部份
    p.Next = q.Next
    q.Next = nil
    return list1
}

func main() {
    l1 := makeListNode([]int{10,1,13,6,9,5})
    l2 := makeListNode([]int{1000000,1000001,1000002})
    fmt.Println("list1: ")
    printListNode(l1)
    fmt.Println("list2: ")
    printListNode(l2)
    fmt.Println("mergeInBetween: ")
    printListNode(mergeInBetween(l1,3,4,l2))

    l1 = makeListNode([]int{0,1,2,3,4,5,6})
    l2 = makeListNode([]int{1000000,1000001,1000002,1000003,1000004})
    fmt.Println("list1: ")
    printListNode(l1)
    fmt.Println("list2: ")
    printListNode(l2)
    fmt.Println("mergeInBetween: ")
    printListNode(mergeInBetween(l1,2,5,l2))

    l1 = makeListNode([]int{10,1,13,6,9,5})
    l2 = makeListNode([]int{1000000,1000001,1000002})
    fmt.Println("list1: ")
    printListNode(l1)
    fmt.Println("list2: ")
    printListNode(l2)
    fmt.Println("mergeInBetween1: ")
    printListNode(mergeInBetween1(l1,3,4,l2))

    l1 = makeListNode([]int{0,1,2,3,4,5,6})
    l2 = makeListNode([]int{1000000,1000001,1000002,1000003,1000004})
    fmt.Println("list1: ")
    printListNode(l1)
    fmt.Println("list2: ")
    printListNode(l2)
    fmt.Println("mergeInBetween1: ")
    printListNode(mergeInBetween1(l1,2,5,l2))
}