package main

// 21. Merge Two Sorted Lists
// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

// Example 1:
// before:
// [1] -> [2] -> [4]
// (1) -> (3) -> (4)
// after:
// [1] -> (1) -> [2] -> (3) -> [4] -> (4)
// <img src="https://assets.leetcode.com/uploads/2020/10/03/merge_ex1.jpg" />
// Input: list1 = [1,2,4], list2 = [1,3,4]
// Output: [1,1,2,3,4,4]

// Example 2:
// Input: list1 = [], list2 = []
// Output: []

// Example 3:
// Input: list1 = [], list2 = [0]
// Output: [0]
 
// Constraints:
//     The number of nodes in both lists is in the range [0, 50].
//     -100 <= Node.val <= 100
//     Both list1 and list2 are sorted in non-decreasing order.


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
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if nil == l1 && nil == l2 {
        return nil
    }

    var t = &ListNode{-1,nil}
    var l3 = &ListNode{-1,t}

    for {
        if nil == l1 || nil == l2 {
            break
        }
        if l1.Val < l2.Val {
            t.Next = l1
            l1 = l1.Next
        } else {
            t.Next = l2
            l2 = l2.Next
        }
        t = t.Next
    }

    if nil == l1 {
        t.Next = l2
    } else {
        t.Next = l1
    }
    return l3.Next.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
    // 头节点
    var head = &ListNode{}
    var last = head
    for list1 != nil && list2 != nil{
        // list1 & list2 谁小合并谁
        if list1.Val <= list2.Val {
            last.Next = list1
            list1 = list1.Next
        } else {
            last.Next = list2
            list2 = list2.Next
        }
        last = last.Next
    }
    // list1 还有,则合并余下的 list1
    for list1 != nil {
        last.Next = list1
        list1 = list1.Next
        last = last.Next
    }
    // list2 还有,则合并余下的 list2
    for list2 != nil {
        last.Next = list2
        list2 = list2.Next
        last = last.Next
    }
    return head.Next
}


func main() {
	// Input: 1->2->4, 1->3->4
	//var l1 *ListNode;
	// var l11 = &ListNode{4,nil}
	// var l12 = &ListNode{2,l11}
	// var l13 = &ListNode{1,l12}

	// var l21 = &ListNode{4,nil}
	// var l22 = &ListNode{3,l21}
	// var l23 = &ListNode{1,l22}

	// printListNode(l11)
	// printListNode(l12)
	// printListNode(l13)
	// printListNode(l23)
    // printListNode(mergeTwoLists(l13,l23))

    l11 := makeListNode([]int{1,2,4})
    l12 := makeListNode([]int{1,3,4})
    fmt.Println("l11: ")
    printListNode(l11)
    fmt.Println("l12: ")
    printListNode(l12)
    fmt.Println("merged: ")
	printListNode(mergeTwoLists(l11,l12))

    l21 := makeListNode([]int{1,2,4})
    l22 := makeListNode([]int{1,3,4})
    fmt.Println("l21: ")
    printListNode(l21)
    fmt.Println("l22: ")
    printListNode(l22)
    fmt.Println("merged: ")
	printListNode(mergeTwoLists1(l21,l22))
}