package main

// 面试题 02.02. Kth Node From End of List LCCI
// Implement an algorithm to find the kth to last element of a singly linked list. 
// Return the value of the element.

// Note: This problem is slightly different from the original one in the book.

// Example:
// Input:  1->2->3->4->5 和 k = 2
// Output:  4

// Note:
//     k is always valid.

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
// 快慢指针
func kthToLast(head *ListNode, k int) int {
    slow, fast := head, head
    for ; k > 0; k-- {
        fast = fast.Next
    }
    for fast != nil {
        fast, slow = fast.Next, slow.Next
    }
    return slow.Val
}

func main() {
    // Example:
    // Input:  1->2->3->4->5 和 k = 2
    // Output:  4
    list1 := makeListNode([]int{1,2,3,4,5})
    printListNode(list1) // 1 -> 2 -> 3 -> 4 -> 5
    fmt.Println(kthToLast(list1, 2)) // 4

    list2 := makeListNode([]int{1,2,3,4,5,6,7,8,9})
    printListNode(list2) // 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7 -> 8 -> 9
    fmt.Println(kthToLast(list2, 2)) // 8

    list3 := makeListNode([]int{9,8,7,6,5,4,3,2,1})
    printListNode(list3) // 9 -> 8 -> 7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1
    fmt.Println(kthToLast(list3, 2)) // 2
}