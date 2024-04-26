package main

// 117. Populating Next Right Pointers in Each Node II
// Given a binary tree
//     struct Node {
//         int val;
//         Node *left;
//         Node *right;
//         Node *next;
//     }

// Populate each next pointer to point to its next right node. 
// If there is no next right node, the next pointer should be set to NULL.
// Initially, all next pointers are set to NULL.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/02/15/117_sample.png" />
// Input: root = [1,2,3,4,5,null,7]
// Output: [1,#,2,3,#,4,5,7,#]
// Explanation: Given the above binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

// Example 2:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 6000].
//     -100 <= Node.val <= 100

// Follow-up:
//     You may only use constant extra space.
//     The recursive approach is fine. You may assume implicit stack space does not count as extra space for this problem.

import "fmt"

type Node struct {
    Val int
    Left *Node
    Right *Node
    Next *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
type queue []*Node

func (q *queue) push(n *Node) {
    (*q) = append((*q), n)
}

func (q *queue) pop() *Node {
    t := (*q)[0]
    (*q) = (*q)[1:]
    return t
}

func (q queue) isEmpty() bool {
    return len(q) == 0
}

func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    q := new(queue)
    q.push(root)
    for !q.isEmpty() {
        var prev *Node
        q1 := new(queue)
        for !q.isEmpty() {
            t := q.pop()
            t.Next = prev
            prev = t
            if t.Right != nil {
                q1.push(t.Right)
            }
            if t.Left != nil {
                q1.push(t.Left)
            }
        }
        for !q1.isEmpty() {
            q.push(q1.pop())
        }
    }
    return root
}

func connect1(root *Node) *Node {
    if root == nil {
        return nil
    }
    nodeRecord := []*Node{root}
    for len(nodeRecord) > 0 {
        curSize := len(nodeRecord)
        for i := 0; i < curSize; i++ {
            curNode := nodeRecord[0]
            nodeRecord = nodeRecord[1:]
            if i < curSize-1 {
                curNode.Next = nodeRecord[0]
            } else {
                curNode.Next = nil
            }
            if curNode.Left != nil {
                nodeRecord = append(nodeRecord, curNode.Left)
            }
            if curNode.Right != nil {
                nodeRecord = append(nodeRecord, curNode.Right)
            }
        }
    }
    return root
}

func main() {
    fmt.Println()
}