package main

// 1836. Remove Duplicates From an Unsorted Linked List
// Given the head of a linked list, 
// find all the values that appear more than once in the list and delete the nodes that have any of those values.

// Return the linked list after the deletions.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/21/tmp-linked-list.jpg" />
// Input: head = [1,2,3,2]
// Output: [1,3]
// Explanation: 2 appears twice in the linked list, so all 2's should be deleted. After deleting all 2's, we are left with [1,3].

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/04/21/tmp-linked-list-1.jpg" />
// Input: head = [2,1,1,2]
// Output: []
// Explanation: 2 and 1 both appear twice. All the elements should be deleted.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2021/04/21/tmp-linked-list-2.jpg" />
// Input: head = [3,2,2,1,3,2,4]
// Output: [1,4]
// Explanation: 3 appears twice and 2 appears three times. After deleting all 3's and 2's, we are left with [1,4].

// Constraints:
//     The number of nodes in the list is in the range [1, 10^5]
//     1 <= Node.val <= 10^5


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
func deleteDuplicatesUnsorted(head *ListNode) *ListNode {
    mp := make(map[int]int)
    for n := head; n != nil; n = n.Next { // 统计出现数量
        mp[n.Val]++
    }
    dummy := &ListNode{0,head}
    pre := dummy
    for n := head; n != nil; n = n.Next {
        if mp[n.Val] > 1 { // 重复的剔除
            pre.Next = n.Next
        } else {
            pre = n
        }
    }
    return dummy.Next
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/04/21/tmp-linked-list.jpg" />
    // Input: head = [1,2,3,2]
    // Output: [1,3]
    // Explanation: 2 appears twice in the linked list, so all 2's should be deleted. After deleting all 2's, we are left with [1,3].
    list1 := makeListNode([]int{1,2,3,2})
    printListNode(list1) // 1 -> 2 -> 3 -> 2
    printListNode(deleteDuplicatesUnsorted(list1)) // 1 -> 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/04/21/tmp-linked-list-1.jpg" />
    // Input: head = [2,1,1,2]
    // Output: []
    // Explanation: 2 and 1 both appear twice. All the elements should be deleted.
    list2 := makeListNode([]int{2,1,1,2})
    printListNode(list2) // 2 -> 1 -> 1 -> 2
    printListNode(deleteDuplicatesUnsorted(list2)) // []
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2021/04/21/tmp-linked-list-2.jpg" />
    // Input: head = [3,2,2,1,3,2,4]
    // Output: [1,4]
    // Explanation: 3 appears twice and 2 appears three times. After deleting all 3's and 2's, we are left with [1,4].
    list3 := makeListNode([]int{3,2,2,1,3,2,4})
    printListNode(list3) // 3 -> 2 -> 2 -> 1 -> 3 -> 2 -> 4
    printListNode(deleteDuplicatesUnsorted(list3)) // 1 -> 4
}