package main

// 83. Remove Duplicates from Sorted List
// Given the head of a sorted linked list, delete all duplicates such that each element appears only once.
// Return the linked list sorted as well.

// Example 1:
// 1 -> 2 -> 2 -> 3 -> 9 -> 9 -> 10 -> 11
//            |
// 1 -> 2 -> 3 -> 9 -> 10 -> 11
// <img src="https://assets.leetcode.com/uploads/2021/01/04/list1.jpg" />
// Input: head = [1,1,2]
// Output: [1,2]

// Example 2:
// 1 -> 1 -> 2 -> 3 -> 3
//      |
// 1 -> 2 -> 3
// <img src="https://assets.leetcode.com/uploads/2021/01/04/list2.jpg" />
// Input: head = [1,1,2,3,3]
// Output: [1,2,3]
 
// Constraints:
//     The number of nodes in the list is in the range [0, 300].
//     -100 <= Node.val <= 100
//     The list is guaranteed to be sorted in ascending order.
    
import "fmt"

type ListNode struct {
    Val  int
    Next *ListNode
}

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

func makeNodeList(nums []int) *ListNode {
    var n = &ListNode{-1, nil}
    var b = &ListNode{-1, n}
    for i := 0; i < len(nums); i++ {
        n.Next = &ListNode{nums[i], nil}
        n = n.Next
    }
    return b.Next.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if nil == head {
        return nil
    }
    if nil == head.Next {
        return head
    }
    var t = head.Val
    var n = &ListNode{t, nil}
    var n1 = &ListNode{-1, n}

    for nil != head {
        // 如果值大于上一个,就链上
        if head.Val > t {
            t = head.Val
            n.Next = &ListNode{head.Val, nil}
            n = n.Next
        }
        head = head.Next
    }
    return n1.Next
}

func main() {
    l1 := makeNodeList([]int{1, 2, 2, 3, 9, 9, 10, 11})
    fmt.Println("before: ")
    printListNode(l1) // 1 -> 2 -> 2 -> 3 -> 9 -> 9 -> 10 -> 11
    fmt.Println("after: ")
    printListNode(deleteDuplicates(l1)) // 1 -> 2 -> 3 -> 9 -> 10 -> 11

    l2 := makeNodeList([]int{1,1,2,3,3})
    fmt.Println("before: ")
    printListNode(l2) // 1 -> 1 -> 2 -> 3 -> 3
    fmt.Println("after: ")
    printListNode(deleteDuplicates(l2)) // 1 -> 2 -> 3
 
    l3 := makeNodeList([]int{1, 2, 2, 3, 9, 9, 10, 11,11,22,23,23})
    fmt.Println("before: ")
    printListNode(l3) // 1 -> 2 -> 2 -> 3 -> 9 -> 9 -> 10 -> 11 -> 11 -> 22 -> 23 -> 23
    fmt.Println("after: ")
    printListNode(deleteDuplicates(l3)) // 1 -> 2 -> 3 -> 9 -> 10 -> 11 -> 22 -> 23
}
