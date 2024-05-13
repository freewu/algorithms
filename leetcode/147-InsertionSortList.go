package main

// 147. Insertion Sort List
// Given the head of a singly linked list, sort the list using insertion sort, and return the sorted list's head.
// The steps of the insertion sort algorithm:
//     Insertion sort iterates, consuming one input element each repetition and growing a sorted output list.
//     At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list and inserts it there.
//     It repeats until no input elements remain.

// The following is a graphical example of the insertion sort algorithm. 
// The partially sorted list (black) initially contains only the first element in the list. 
// One element (red) is removed from the input data and inserted in-place into the sorted list with each iteration.
// <img src="https://upload.wikimedia.org/wikipedia/commons/0/0f/Insertion-sort-example-300px.gif" />

// Example 1:
// 4 -> 2 -> 1 -> 3   =>   1 -> 2 -> 3 -> 4
// <img src="https://assets.leetcode.com/uploads/2021/03/04/sort1linked-list.jpg" />
// Input: head = [4,2,1,3]
// Output: [1,2,3,4]

// Example 2:
// -1 -> 5 -> 3 -> 4 -> 0   =>  -1 -> 0 -> 3 -> 4 -> 4 -> 5
// <img src="https://assets.leetcode.com/uploads/2021/03/04/sort2linked-list.jpg" />
// Input: head = [-1,5,3,4,0]
// Output: [-1,0,3,4,5]
 
// Constraints:
//     The number of nodes in the list is in the range [1, 5000].
//     -5000 <= Node.val <= 5000

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
func insertionSortList(head *ListNode) *ListNode {
    dummy := new(ListNode)
    for head != nil {
        cur := dummy
        for ; cur.Next != nil && cur.Next.Val < head.Val; cur = cur.Next {
        }
        cur.Next, head.Next, head = head, cur.Next, head.Next
    }
    return dummy.Next
}

func insertionSortList1(head *ListNode) *ListNode {
    if head == nil { return nil; }
    dummyHead:=&ListNode{ Next:head }
    lastSorted, cur := head, head.Next
    for cur != nil {
        if lastSorted.Val <= cur.Val {
            lastSorted = lastSorted.Next
        } else {
            pre := dummyHead
            for pre.Next.Val <= cur.Val {
                pre = pre.Next
            }
            lastSorted.Next = cur.Next
            cur.Next = pre.Next
            pre.Next = cur
        }
        cur = lastSorted.Next
    }
    return dummyHead.Next
}

func main() {
    // Example 1:
    // 4 -> 2 -> 1 -> 3   =>   1 -> 2 -> 3 -> 4
    // <img src="https://assets.leetcode.com/uploads/2021/03/04/sort1linked-list.jpg" />
    // Input: head = [4,2,1,3]
    // Output: [1,2,3,4]
    l1 := makeListNode([]int{4,2,1,3})
    printListNode(l1)
    printListNode(insertionSortList(l1))
    // Example 2:
    // -1 -> 5 -> 3 -> 4 -> 0   =>  -1 -> 0 -> 3 -> 4 -> 4 -> 5
    // <img src="https://assets.leetcode.com/uploads/2021/03/04/sort2linked-list.jpg" />
    // Input: head = [-1,5,3,4,0]
    // Output: [-1,0,3,4,5]
    l2 := makeListNode([]int{-1,5,3,4,0})
    printListNode(l2)
    printListNode(insertionSortList(l2))

    l11 := makeListNode([]int{4,2,1,3})
    printListNode(l11)
    printListNode(insertionSortList(l11))
    l12 := makeListNode([]int{-1,5,3,4,0})
    printListNode(l12)
    printListNode(insertionSortList(l12))
}