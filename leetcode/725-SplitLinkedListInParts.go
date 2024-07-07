package main

// 725. Split Linked List in Parts
// Given the head of a singly linked list and an integer k, split the linked list into k consecutive linked list parts.
// The length of each part should be as equal as possible: no two parts should have a size differing by more than one. This may lead to some parts being null.
// The parts should be in the order of occurrence in the input list, and parts occurring earlier should always have a size greater than or equal to parts occurring later.
// Return an array of the k parts.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/split1-lc.jpg" />
// Input: head = [1,2,3], k = 5
// Output: [[1],[2],[3],[],[]]
// Explanation:
// The first element output[0] has output[0].val = 1, output[0].next = null.
// The last element output[4] is null, but its string representation as a ListNode is [].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/split2-lc.jpg" />
// Input: head = [1,2,3,4,5,6,7,8,9,10], k = 3
// Output: [[1,2,3,4],[5,6,7],[8,9,10]]
// Explanation:
// The input has been split into consecutive parts with size difference at most 1, and earlier parts are a larger size than the later parts.
 
// Constraints:
//     The number of nodes in the list is in the range [0, 1000].
//     0 <= Node.val <= 1000
//     1 <= k <= 50

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
func splitListToParts(head *ListNode, k int) []*ListNode {
    list := []*ListNode{}
    for head != nil {
        list = append(list, head)
        head = head.Next
    }
    size, div := len(list) / k, len(list) % k
    if size == 0 {
        size = 1
        div = 0
    }
    res := []*ListNode{}
    i := 0
    for i < len(list) {
        if i > 0 {
            list[i-1].Next = nil
        }
        res = append(res, list[i])
        i += size
        if div > 0 {
            i++
            div--
        }
    }
    for len(res) < k {
        res = append(res, nil)
    }
    return res
}

func splitListToParts1(head *ListNode, k int) []*ListNode {
    length := 0 // 链表长度
    cur := head 
    for cur != nil {
        length++
        cur = cur.Next
    }
    quotient, remain := length / k, length % k
    res := make([]*ListNode, k) // 分割成k个链表
    cur = head
    for i := 0; i < k && cur != nil; i++ {
        res[i] = cur
        size := quotient
        if i < remain {
            size++
        }
        // 分配节点
        for j := 1; j < size; j++ {
            cur = cur.Next
        }
        cur, cur.Next = cur.Next, nil
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/split1-lc.jpg" />
    // Input: head = [1,2,3], k = 5
    // Output: [[1],[2],[3],[],[]]
    // Explanation:
    // The first element output[0] has output[0].val = 1, output[0].next = null.
    // The last element output[4] is null, but its string representation as a ListNode is [].
    list1 := makeListNode([]int{1,2,3})
    printListNode(list1) // 1 -> 2 -> 3
    fmt.Println(splitListToParts(list1, 5)) // [0xc000028090 0xc000028080 0xc000028070 <nil> <nil>]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/split2-lc.jpg" />
    // Input: head = [1,2,3,4,5,6,7,8,9,10], k = 3
    // Output: [[1,2,3,4],[5,6,7],[8,9,10]]
    // Explanation:
    // The input has been split into consecutive parts with size difference at most 1, and earlier parts are a larger size than the later parts.
    list2 := makeListNode([]int{1,2,3,4,5,6,7,8,9,10})
    printListNode(list2) // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10
    fmt.Println(splitListToParts(list2, 3)) // [0xc000028150 0xc000028110 0xc0000280e0]

    list11 := makeListNode([]int{1,2,3})
    printListNode(list11) // 1 -> 2 -> 3
    fmt.Println(splitListToParts1(list11, 5)) // [0xc000028090 0xc000028080 0xc000028070 <nil> <nil>]
    list12 := makeListNode([]int{1,2,3,4,5,6,7,8,9,10})
    printListNode(list12) // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9 -> 10
    fmt.Println(splitListToParts1(list12, 3)) // [0xc000028150 0xc000028110 0xc0000280e0]
}