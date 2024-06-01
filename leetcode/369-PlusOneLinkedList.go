package main

// 369. Plus One Linked List
// Given a non-negative integer represented as a linked list of digits, plus one to the integer.
// The digits are stored such that the most significant digit is at the head of the list.

// Example 1:
// Input: head = [1,2,3]
// Output: [1,2,4]

// Example 2:
// Input: head = [0]
// Output: [1]

// Constraints:
//     The number of nodes in the linked list is in the range [1, 100].
//     0 <= Node.val <= 9
//     The number represented by the linked list does not contain leading zeros except for the zero itself. 

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
// dfs
func plusOne(head *ListNode) *ListNode {
    var dfs func(node *ListNode) int
    dfs = func(node *ListNode) int {
        if node == nil {
            return 1
        }
        num := dfs(node.Next)
        val := node.Val
        node.Val = (val + num) % 10
        return int((val + num) / 10)
    }
    num := dfs(head)
    if num == 0 {
        return head
    }
    pre := &ListNode{Val:1, Next:head} //  处理  999...9 这种需要进位的
    return pre
}

// 快慢指针
func plusOne1(head *ListNode) *ListNode {
    fast, slow := head, &ListNode{ Next: head }
    for fast != nil {
        if fast.Val != 9 {
            slow = fast
        }
        fast = fast.Next
    }
    slow.Val += 1
    cur := slow.Next
    for cur != nil {
        cur.Val = 0
        cur = cur.Next
    }
    if slow.Next == head {
        return slow
    }
    return head
}

func main() {
    // Example 1:
    // Input: head = [1,2,3]
    // Output: [1,2,4]
    list1 := makeListNode([]int{1,2,4})
    printListNode(list1) // 1 -> 2 -> 3
    printListNode(plusOne(list1)) // 1 -> 2 -> 4
    // Example 2:
    // Input: head = [0]
    // Output: [1]
    list2 := makeListNode([]int{0})
    printListNode(list2) // 0
    printListNode(plusOne(list2)) // 1

    list3 := makeListNode([]int{9,9,9})
    printListNode(list3) // 9 -> 9 -> 9
    printListNode(plusOne(list3)) // 1 -> 0 -> 0 -> 0

    list11 := makeListNode([]int{1,2,4})
    printListNode(list11) // 1 -> 2 -> 3
    printListNode(plusOne1(list11)) // 1 -> 2 -> 4
    list12 := makeListNode([]int{0})
    printListNode(list12) // 0
    printListNode(plusOne1(list12)) // 1
    list13 := makeListNode([]int{9,9,9})
    printListNode(list13) // 9 -> 9 -> 9
    printListNode(plusOne1(list13)) // 1 -> 0 -> 0 -> 0
}