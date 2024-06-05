package main

// 430. Flatten a Multilevel Doubly Linked List
// You are given a doubly linked list, which contains nodes that have a next pointer, a previous pointer,
// and an additional child pointer. This child pointer may or may not point to a separate doubly linked list, 
// also containing these special nodes. These child lists may have one or more children of their own, 
// and so on, to produce a multilevel data structure as shown in the example below.

// Given the head of the first level of the list, flatten the list so that all the nodes appear in a single-level, doubly linked list. 
// Let curr be a node with a child list. 
// The nodes in the child list should appear after curr and before curr.next in the flattened list.

// Return the head of the flattened list. The nodes in the list must have all of their child pointers set to null.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten11.jpg" />
// Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
// Output: [1,2,3,7,8,11,12,9,10,4,5,6]
// Explanation: The multilevel linked list in the input is shown.
// After flattening the multilevel linked list it becomes:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten12.jpg" />

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten2.1jpg" />
// Input: head = [1,2,null,3]
// Output: [1,3,2]
// Explanation: The multilevel linked list in the input is shown.
// After flattening the multilevel linked list it becomes:
// <img src="https://assets.leetcode.com/uploads/2021/11/24/list.jpg" />

// Example 3:
// Input: head = []
// Output: []
// Explanation: There could be empty list in the input.
 
// Constraints:
//     The number of Nodes will not exceed 1000.
//     1 <= Node.val <= 10^5
 
// How the multilevel linked list is represented in test cases:
// We use the multilevel linked list from Example 1 above:
//  1---2---3---4---5---6--NULL
//          |
//          7---8---9---10--NULL
//              |
//              11--12--NULL

// The serialization of each level is as follows:
//     [1,2,3,4,5,6,null]
//     [7,8,9,10,null]
//     [11,12,null]

// To serialize all levels together, we will add nulls in each level to signify no node connects to the upper node of the previous level. 
// The serialization becomes:
//     [1,    2,    3, 4, 5, 6, null]
//                 |
//     [null, null, 7,    8, 9, 10, null]
//                     |
//     [            null, 11, 12, null]

// Merging the serialization of each level and removing trailing nulls we obtain:
//     [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]

import "fmt"

type Node struct {
    Val int
    Prev *Node
    Next *Node
    Child *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Prev *Node
 *     Next *Node
 *     Child *Node
 * }
 */
// stack
func flatten(root *Node) *Node {
    stack, head := []*Node{}, root
    for root != nil || len(stack) != 0 {
        if root != nil && root.Child != nil {
            if root.Next != nil {
                stack = append(stack, root.Next)
            }
            child := root.Child
            root.Child = nil
            root.Next = child
            child.Prev = root
            root = child
            continue
        }
        if root != nil && root.Next != nil {
            root = root.Next
            continue
        }
        if len(stack) != 0 {
            node := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            root.Next = node
            node.Prev = root
            root = node
            continue
        }
        break
    }
    return head
}

// dfs
func flatten1(root *Node) *Node {
    var dfs func(*Node) *Node
    dfs = func(node *Node) (last *Node) {
        cur := node
        for cur != nil {
            next := cur.Next
            if cur.Child != nil {
                childLast := dfs(cur.Child)
                next = cur.Next
                cur.Next = cur.Child
                cur.Child.Prev = cur
                if next != nil {
                    childLast.Next = next
                    next.Prev = childLast
                }
                cur.Child = nil
                last = childLast
            } else {
                last = cur
            }
            cur = next
        }
        return
    }
    dfs(root)
    return root
}

func main() {
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten11.jpg" />
// Input: head = [1,2,3,4,5,6,null,null,null,7,8,9,10,null,null,11,12]
// Output: [1,2,3,7,8,11,12,9,10,4,5,6]
// Explanation: The multilevel linked list in the input is shown.
// After flattening the multilevel linked list it becomes:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten12.jpg" />

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/11/09/flatten2.1jpg" />
// Input: head = [1,2,null,3]
// Output: [1,3,2]
// Explanation: The multilevel linked list in the input is shown.
// After flattening the multilevel linked list it becomes:
// <img src="https://assets.leetcode.com/uploads/2021/11/24/list.jpg" />

// Example 3:
// Input: head = []
// Output: []
// Explanation: There could be empty list in the input.
}