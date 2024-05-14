package main

// 328. Odd Even Linked List
// Given the head of a singly linked list, group all the nodes with odd indices together followed by the nodes with even indices, and return the reordered list.
// The first node is considered odd, and the second node is even, and so on.
// Note that the relative order inside both the even and odd groups should remain as it was in the input.
// You must solve the problem in O(1) extra space complexity and O(n) time complexity.

// Example 1:
// 1 -> (2) -> 3 -> (4) -> 5  =>  1 -> 3 -> 5 -> (2） -> (4)
// <img src="https://assets.leetcode.com/uploads/2021/03/10/oddeven-linked-list.jpg" />
// Input: head = [1,2,3,4,5]
// Output: [1,3,5,2,4]

// Example 2:
// 2 -> (1) -> 3 -> (5) -> 6 -> (4) -> 7   =>  2 -> 3 -> 6 -> 7 -> (1) -> (4) -> (5)
// <img src="https://assets.leetcode.com/uploads/2021/03/10/oddeven2-linked-list.jpg" />
// Input: head = [2,1,3,5,6,4,7]
// Output: [2,3,6,7,1,5,4]

// Constraints:
//     The number of nodes in the linked list is in the range [0, 10^4].
//     -10^6 <= Node.val <= 10^6

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
func oddEvenList(head *ListNode) *ListNode {
    index, current, odd, even := 1, head, []int{}, []int{}
    for current != nil { // 取出奇数位&偶数位的数据到不同的数组中
        if index % 2 == 1 {
            odd = append(odd, current.Val)
        } else {
            even = append(even, current.Val)
        }
        current = current.Next
        index++
    }
    current = head
    for _, v := range odd { // 重放奇数位数组到链表开头
        current.Val = v
        current = current.Next
    }
    for _, v := range even {  // 重放偶数位数组至链表结束
        current.Val = v
        current = current.Next
    }
    return head
}

func oddEvenList1(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    odd,even := head, head.Next
    next := even.Next
    evenHead := even
    for {
        if even == nil || next == nil {
            break
        }
        odd.Next, even.Next  = next, next.Next
        odd, even = odd.Next, even.Next
        if even != nil {
            next = even.Next
        }
    }
    odd.Next = evenHead
    return head
}

func main() {
    // Example 1:
    // 1 -> (2) -> 3 -> (4) -> 5  =>  1 -> 3 -> 5 -> (2） -> (4)
    // <img src="https://assets.leetcode.com/uploads/2021/03/10/oddeven-linked-list.jpg" />
    // Input: head = [1,2,3,4,5]
    // Output: [1,3,5,2,4]
    l1 := makeListNode([]int{1,2,3,4,5})
    printListNode(l1) // 1 -> (2) -> 3 -> (4) -> 5
    printListNode(oddEvenList(l1)) // 1 -> 3 -> 5 -> (2） -> (4)
    // Example 2:
    // 2 -> (1) -> 3 -> (5) -> 6 -> (4) -> 7   =>  2 -> 3 -> 6 -> 7 -> (1) -> (4) -> (5)
    // <img src="https://assets.leetcode.com/uploads/2021/03/10/oddeven2-linked-list.jpg" />
    // Input: head = [2,1,3,5,6,4,7]
    // Output: [2,3,6,7,1,5,4]
    l2 := makeListNode([]int{2,1,3,5,6,4,7})
    printListNode(l2) // 2 -> (1) -> 3 -> (5) -> 6 -> (4) -> 7 
    printListNode(oddEvenList(l2)) // 2 -> 3 -> 6 -> 7 -> (1) -> (4) -> (5)

    l11 := makeListNode([]int{1,2,3,4,5})
    printListNode(l11) // 1 -> (2) -> 3 -> (4) -> 5
    printListNode(oddEvenList(l11)) // 1 -> 3 -> 5 -> (2） -> (4)
    l12 := makeListNode([]int{2,1,3,5,6,4,7})
    printListNode(l12) // 2 -> (1) -> 3 -> (5) -> 6 -> (4) -> 7 
    printListNode(oddEvenList(l12)) // 2 -> 3 -> 6 -> 7 -> (1) -> (4) -> (5)
}