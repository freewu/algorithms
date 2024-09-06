package main

// 3217. Delete Nodes From Linked List Present in Array
// You are given an array of integers nums and the head of a linked list. 
// Return the head of the modified linked list after removing all nodes from the linked list that have a value that exists in nums.

// Example 1:
// Input: nums = [1,2,3], head = [1,2,3,4,5]
// Output: [4,5]
// Explanation:
// Remove the nodes with values 1, 2, and 3.

// Example 2:
// Input: nums = [1], head = [1,2,1,2,1,2]
// Output: [2,2,2]
// Explanation:
// Remove the nodes with value 1.

// Example 3:
// Input: nums = [5], head = [1,2,3,4]
// Output: [1,2,3,4]
// Explanation:
// No node has value 5.

// Constraints:
//     1 <= nums.length <= 10^5
//     1 <= nums[i] <= 10^5
//     All elements in nums are unique.
//     The number of nodes in the given list is in the range [1, 10^5].
//     1 <= Node.val <= 10^5
//     The input is generated such that there is at least one node in the linked list that has a value not present in nums.

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
func modifiedList(nums []int, head *ListNode) *ListNode {
    mp := make(map[int]bool)
    for i := range nums {
        mp[nums[i]] = true
    }
    dummy := &ListNode{ Val: -1, Next: head, }
    for prev, cur:= dummy, dummy.Next; cur != nil; cur = cur.Next {
        if mp[cur.Val] != true {
            prev = cur
            continue
        }
        prev.Next = cur.Next // 删除节点
    }
    return dummy.Next
}

func main() {
    // Example 1:
    // Input: nums = [1,2,3], head = [1,2,3,4,5]
    // Output: [4,5]
    // Explanation:
    // Remove the nodes with values 1, 2, and 3.
    list1 := makeListNode([]int{1,2,3,4,5})
    printListNode(list1) // 1 -> 2 -> 3 -> 4 -> 5
    printListNode(modifiedList([]int{1,2,3}, list1)) // 4 -> 5
    // Example 2:
    // Input: nums = [1], head = [1,2,1,2,1,2]
    // Output: [2,2,2]
    // Explanation:
    // Remove the nodes with value 1.
    list2 := makeListNode([]int{1,2,1,2,1,2})
    printListNode(list2) // 1 -> 2 -> 1 -> 2 -> 1 -> 2
    printListNode(modifiedList([]int{1}, list2)) // 2 -> 2 -> 2
    // Example 3:
    // Input: nums = [5], head = [1,2,3,4]
    // Output: [1,2,3,4]
    // Explanation:
    // No node has value 5.
    list3 := makeListNode([]int{1,2,3,4})
    printListNode(list3) // 1 -> 2 -> 3 -> 4
    printListNode(modifiedList([]int{5}, list3)) // 1 -> 2 -> 3 -> 4
}