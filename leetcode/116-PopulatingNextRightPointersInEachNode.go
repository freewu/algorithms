package main

// 116. Populating Next Right Pointers in Each Node
// You are given a perfect binary tree where all leaves are on the same level, 
// and every parent has two children. The binary tree has the following definition:
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
// <img src="https://assets.leetcode.com/uploads/2019/02/14/116_sample.png" />
// Input: root = [1,2,3,4,5,6,7]
// Output: [1,#,2,3,#,4,5,6,7,#]
// Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.

// Example 2:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 2^12 - 1].
//     -1000 <= Node.val <= 1000
    
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
// Recursive
func connect(root *Node) *Node {
    if root == nil {
        return nil
    }
    if root.Left != nil {
        root.Left.Next = root.Right
        if root.Next != nil {
            root.Right.Next = root.Next.Left
        }
    }
    connect(root.Left)
    connect(root.Right)
    return root
}

// Iterative (BFS with Queue):
func connect1(root *Node) *Node {
    if root == nil {
        return nil
    }
    
    type entry struct{
        level int
        node *Node
    }
    
    queue, prev, currLevel := []entry{}, &Node{}, -1
    queue = append(queue, entry{level: 0, node: root})
    for len(queue) != 0 {
        curr := queue[0]
        queue = queue[1:]
        if currLevel != curr.level {
            if prev != nil {
                prev.Next = nil
            }
            currLevel = curr.level
        } else {
            prev.Next = curr.node
        }
        prev = curr.node
        if curr.node.Left != nil {
            queue = append(queue, entry{level: curr.level+1, node: curr.node.Left})
        }
        if curr.node.Right != nil {
            queue = append(queue, entry{level: curr.level+1, node: curr.node.Right})
        }
    }
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/02/14/116_sample.png" />
    // Input: root = [1,2,3,4,5,6,7]
    // Output: [1,#,2,3,#,4,5,6,7,#]
    // Explanation: Given the above perfect binary tree (Figure A), your function should populate each next pointer to point to its next right node, just like in Figure B. The serialized output is in level order as connected by the next pointers, with '#' signifying the end of each level.
    tree1 := &Node {
        1,
        &Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil},
        &Node{3, &Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil},
        nil,
    }
    fmt.Println(connect(tree1)) // [1,#,2,3,#,4,5,6,7,#]
    // Example 2:
    // Input: root = []
    // Output: []
    fmt.Println(connect(nil)) // nil

    tree11 := &Node {
        1,
        &Node{2, &Node{4, nil, nil, nil}, &Node{5, nil, nil, nil}, nil},
        &Node{3, &Node{6, nil, nil, nil}, &Node{7, nil, nil, nil}, nil},
        nil,
    }
    fmt.Println(connect1(tree11)) // [1,#,2,3,#,4,5,6,7,#]
    fmt.Println(connect1(nil)) // nil
}