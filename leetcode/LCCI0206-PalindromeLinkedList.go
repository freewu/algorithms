package main

// 面试题 02.06. Palindrome Linked List LCCI
// Implement a function to check if a linked list is a palindrome.

// Example 1:
// Input: 1 -> 2
// Output: false

// Example 2:
// Input: 1->2->2->1
// Output: true 

// Follow up:
//     Could you do it in O(n) time and O(1) space?

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
func isPalindrome(head *ListNode) bool {
    arr := []int{}
    for ; head != nil; head = head.Next { // 遍历链表
        arr = append(arr, head.Val)
    }
    n := len(arr)
    for i, v := range arr[:n/2] {
        if v != arr[n-i-1] {
            return false
        }
    }
    return true
}

// 递归
func isPalindrome1(head *ListNode) bool {
    frontPointer := head
    var check func(*ListNode) bool
    check = func(node *ListNode) bool {
        if node != nil {
            if !check(node.Next) { return false }
            if node.Val != frontPointer.Val { return false }
            frontPointer = frontPointer.Next
        }
        return true
    }
    return check(head)
}

// 快慢指针
func isPalindrome2(head *ListNode) bool {
    slow, fast := head, head
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 原本计划用fast.Next == nil 代表是奇数，但是可能存在fast是nil导致空指针的情况，所以得用fast != nil来表示奇数长度，奇数slow得往下走一步
    if fast != nil {
        slow = slow.Next
    }
    reverse := func(head *ListNode) *ListNode {
        cur := head
        var prev *ListNode 
        for cur != nil {
            next := cur.Next
            cur.Next = prev
            prev = cur
            cur = next
        }
        // prev会一直往后移动的。所以应该返回prev,你返回head的，head的next是空的
        return prev
    }
    left, right := head, reverse(slow)
    for right != nil {
        if left.Val != right.Val {
            return false
        }
        left, right = left.Next, right.Next
    }
    return true
}

func main() {
    // Example 1:
    // Input: 1 -> 2
    // Output: false
    list1 := makeListNode([]int{1,2})
    printListNode(list1) // 1 -> 2
    fmt.Println(isPalindrome(list1)) // false
    // Example 2:
    // Input: 1->2->2->1
    // Output: true 
    list2 := makeListNode([]int{1,2,2,1})
    printListNode(list2) // 1 -> 2 -> 2 -> 1
    fmt.Println(isPalindrome(list2)) // true

    list11 := makeListNode([]int{1,2})
    printListNode(list11) // 1 -> 2
    fmt.Println(isPalindrome1(list11)) // false
    list12 := makeListNode([]int{1,2,2,1})
    printListNode(list12) // 1 -> 2 -> 2 -> 1
    fmt.Println(isPalindrome1(list12)) // true

    list21 := makeListNode([]int{1,2})
    printListNode(list21) // 1 -> 2
    fmt.Println(isPalindrome2(list21)) // false
    list22 := makeListNode([]int{1,2,2,1})
    printListNode(list22) // 1 -> 2 -> 2 -> 1
    fmt.Println(isPalindrome2(list22)) // true
}