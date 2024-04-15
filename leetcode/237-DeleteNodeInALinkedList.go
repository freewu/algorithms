package main

// 237. Delete Node in a Linked List
// There is a singly-linked list head and we want to delete a node node in it.
// You are given the node to be deleted node. You will not be given access to the first node of head.
// All the values of the linked list are unique, and it is guaranteed that the given node node is not the last node in the linked list.
// Delete the given node. Note that by deleting the node, we do not mean removing it from memory. We mean:
//     The value of the given node should not exist in the linked list.
//     The number of nodes in the linked list should decrease by one.
//     All the values before node should be in the same order.
//     All the values after node should be in the same order.

// Custom testing: 
//     For the input, you should provide the entire linked list head and the node to be given node. node should not be the last node of the list and should be an actual node in the list.
//     We will build the linked list and pass the node to your function.
//     The output will be the entire list after calling your function.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/01/node1.jpg" />
// Input: head = [4,5,1,9], node = 5
// Output: [4,1,9]
// Explanation: You are given the second node with value 5, the linked list should become 4 -> 1 -> 9 after calling your function.
// Example 2:

// Input: head = [4,5,1,9], node = 1
// <img src="https://assets.leetcode.com/uploads/2020/09/01/node2.jpg" />
// Output: [4,5,9]
// Explanation: You are given the third node with value 1, the linked list should become 4 -> 5 -> 9 after calling your function.
 
// Constraints:
//     The number of the nodes in the given list is in the range [2, 1000].
//     -1000 <= Node.val <= 1000
//     The value of each node in the list is unique.
//     The node to be deleted is in the list and is not a tail node.

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
func deleteNode(node *ListNode) {
    node.Val = node.Next.Val
    node.Next = node.Next.Next
}

func main() {
    l1 := makeListNode([]int{4,5,1,9})
    lh1 := l1
    printListNode(l1) // 4 -> 5 -> 1 -> 9
    for nil != l1.Next {
        if l1.Val == 5 {
            deleteNode(l1)
        }
        l1 = l1.Next
    }
    printListNode(lh1) // 4 -> 1 -> 9


    l2 := makeListNode([]int{4,5,1,9})
    lh2 := l2
    printListNode(l2) // 4 -> 5 -> 1 -> 9
    // for nil != l2.Next{
    //     if l2.Val == 1 {
    //         deleteNode(l2)
    //     }
    //     l2 = l2.Next
    // }
    printListNode(lh2) // 4 -> 5 -> 9
}