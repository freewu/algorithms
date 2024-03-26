package main

// 226. Invert Binary Tree  
// Given the root of a binary tree, invert the tree, and return its root.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg" />
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg" />
// Input: root = [2,1,3]
// Output: [2,3,1]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [0, 100].
//     -100 <= Node.val <= 100

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// dfs
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    // General cases:
    // invert child node of current root
    root.Left, root.Right = root.Right, root.Left   
    // invert subtree with DFS
    invertTree(root.Left)
    invertTree(root.Right)
    return root
}

func invertTree1(root *TreeNode) *TreeNode {
    if root == nil{
        return nil
    }
    root.Left, root.Right = invertTree1(root.Right), invertTree1(root.Left)
    return root
}

func main() {
    // [4,2,7,1,3,6,9]
    tree1 := &TreeNode {
        4,
        &TreeNode {
            2,
            &TreeNode{1, nil, nil},
            &TreeNode{3, nil, nil},
        },
        &TreeNode {
            7,
            &TreeNode{6, nil, nil},
            &TreeNode{9, nil, nil},
        },
    }
    fmt.Println(invertTree(tree1))

    tree2 := &TreeNode {
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(invertTree(tree2))
}