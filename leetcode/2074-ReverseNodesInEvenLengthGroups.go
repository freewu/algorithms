package main

// 2074. Reverse Nodes in Even Length Groups
// You are given the head of a linked list.

// The nodes in the linked list are sequentially assigned to non-empty groups whose lengths form the sequence of the natural numbers (1, 2, 3, 4, ...). 
// The length of a group is the number of nodes assigned to it. In other words,
//     The 1st node is assigned to the first group.
//     The 2nd and the 3rd nodes are assigned to the second group.
//     The 4th, 5th, and 6th nodes are assigned to the third group, and so on.

// Note that the length of the last group may be less than or equal to 1 + the length of the second to last group.

// Reverse the nodes in each group with an even length, and return the head of the modified linked list.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/10/25/eg1.png" />
// Input: head = [5,2,6,3,9,1,7,3,8,4]
// Output: [5,6,2,3,9,1,4,8,3,7]
// Explanation:
// - The length of the first group is 1, which is odd, hence no reversal occurs.
// - The length of the second group is 2, which is even, hence the nodes are reversed.
// - The length of the third group is 3, which is odd, hence no reversal occurs.
// - The length of the last group is 4, which is even, hence the nodes are reversed.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/10/25/eg2.png" />
// Input: head = [1,1,0,6]
// Output: [1,0,1,6]
// Explanation:
// - The length of the first group is 1. No reversal occurs.
// - The length of the second group is 2. The nodes are reversed.
// - The length of the last group is 1. No reversal occurs.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/11/17/ex3.png" />
// Input: head = [1,1,0,6,5]
// Output: [1,0,1,5,6]
// Explanation:
// - The length of the first group is 1. No reversal occurs.
// - The length of the second group is 2. The nodes are reversed.
// - The length of the last group is 2. The nodes are reversed.

// Constraints:
//     The number of nodes in the list is in the range [1, 10^5].
//     0 <= Node.val <= 10^5

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
    l := len(arr) - 1
    head := &ListNode{arr[l], nil}
    for i := l - 1; i >= 0; i-- {
        n := &ListNode{arr[i], head}
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
func reverseEvenLengthGroups(head *ListNode) *ListNode {
    prevTail, node := head, head
    var prev, curr, next *ListNode 
    groupLen, nodes := 2, 0
    for prevTail.Next != nil {
        for nodes = 0; nodes < groupLen && node.Next != nil; nodes++ {
            node = node.Next
        }
        if nodes % 2 == 0 {
            prev, curr = node.Next, prevTail.Next
            for nodes = 0; nodes < groupLen && curr != nil; nodes++ { // reverse
                if curr == nil { break }
                next = curr.Next
                curr.Next = prev
                prev = curr
                curr = next
            }
            prevTail, prevTail.Next = prevTail.Next, node
        } else {
            prevTail = node
        }
        node = prevTail 
        groupLen++
    }
    return head
}

func reverseEvenLengthGroups1(head *ListNode) *ListNode {
    i, cur := 1, head // 组编号，从1开始 当前节点
    var pre *ListNode     // 前一个组的最后一个节点
    reverse := func (a, b *ListNode) *ListNode { // reverse 反转[a, b)区间的链表，并返回新的头节点
        var pre *ListNode
        cur := a
        for cur != b {
            nxt := cur.Next
            cur.Next = pre
            pre = cur
            cur = nxt
        }
        return pre
    }
    for cur != nil { // 遍历链表
        n, now := 0, cur
        for ; n < i && now != nil; n++ { // 计算当前组的长度
            now = now.Next
        }
        if n % 2 == 1 { // 如果长度为奇数，无需反转，移动指针到下一组的开始
            for j := 0; j < n; j++ {
                pre = cur
                cur = cur.Next
            }
        } else {
            newHead := reverse(cur, now) // 如果长度为偶数，需要反转
            if pre != nil {
                pre.Next = newHead
            } else {
                head = newHead
            }
            cur.Next = now
            pre = cur
            cur = now
        }
        i++  // 组编号递增
    }
    return head
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/10/25/eg1.png" />
    // Input: head = [5,2,6,3,9,1,7,3,8,4]
    // Output: [5,6,2,3,9,1,4,8,3,7]
    // Explanation:
    // - The length of the first group is 1, which is odd, hence no reversal occurs.
    // - The length of the second group is 2, which is even, hence the nodes are reversed.
    // - The length of the third group is 3, which is odd, hence no reversal occurs.
    // - The length of the last group is 4, which is even, hence the nodes are reversed.
    list1 := makeListNode([]int{5,2,6,3,9,1,7,3,8,4})
    printListNode(list1) // 5 -> 2 -> 6 -> 3 -> 9 -> 1 -> 7 -> 3 -> 8 -> 4
    printListNode(reverseEvenLengthGroups(list1)) // 5 -> 6 -> 2 -> 3 -> 9 -> 1 -> 4 -> 8 -> 3 -> 7
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/10/25/eg2.png" />
    // Input: head = [1,1,0,6]
    // Output: [1,0,1,6]
    // Explanation:
    // - The length of the first group is 1. No reversal occurs.
    // - The length of the second group is 2. The nodes are reversed.
    // - The length of the last group is 1. No reversal occurs.
    list2 := makeListNode([]int{1,1,0,6})
    printListNode(list2) // 1 -> 1 -> 0 -> 6
    printListNode(reverseEvenLengthGroups(list2)) // 1 -> 0 -> 1 -> 6
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/11/17/ex3.png" />
    // Input: head = [1,1,0,6,5]
    // Output: [1,0,1,5,6]
    // Explanation:
    // - The length of the first group is 1. No reversal occurs.
    // - The length of the second group is 2. The nodes are reversed.
    // - The length of the last group is 2. The nodes are reversed.
    list3 := makeListNode([]int{1,1,0,6,5})
    printListNode(list3) // 1 -> 1 -> 0 -> 6 -> 5
    printListNode(reverseEvenLengthGroups(list3)) // 1 -> 0 -> 1 -> 5 -> 6

    list11 := makeListNode([]int{5,2,6,3,9,1,7,3,8,4})
    printListNode(list11) // 5 -> 2 -> 6 -> 3 -> 9 -> 1 -> 7 -> 3 -> 8 -> 4
    printListNode(reverseEvenLengthGroups(list11)) // 5 -> 6 -> 2 -> 3 -> 9 -> 1 -> 4 -> 8 -> 3 -> 7
    list12 := makeListNode([]int{1,1,0,6})
    printListNode(list12) // 1 -> 1 -> 0 -> 6
    printListNode(reverseEvenLengthGroups(list12)) // 1 -> 0 -> 1 -> 6
    list13 := makeListNode([]int{1,1,0,6,5})
    printListNode(list13) // 1 -> 1 -> 0 -> 6 -> 5
    printListNode(reverseEvenLengthGroups(list13)) // 1 -> 0 -> 1 -> 5 -> 6
}