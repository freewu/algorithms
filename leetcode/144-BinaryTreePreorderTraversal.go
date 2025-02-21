package main

// 144. Binary Tree Preorder Traversal
// Given the root of a binary tree, return the preorder traversal of its nodes' values.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg"/>
// Input: root = [1,null,2,3]
// Output: [1,2,3]

// Example 2:
// Input: root = []
// Output: []

// Example 3:
// Input: root = [1]
// Output: [1]

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
func preorderTraversal(root *TreeNode) []int {
    var res []int
    preorder(root, &res)
    return res
}

// 递归
func preorder (root *TreeNode, output *[]int) {
    if root != nil {
        // 前序遍历
        *output = append(*output, root.Val)
        preorder(root.Left, output)
        preorder(root.Right, output)
    }
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg"/>
    // Input: root = [1,null,2,3]
    // Output: [1,2,3]
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode {
            2,
            &TreeNode{3, nil, nil},
            nil,
        },
    }
    fmt.Println(preorderTraversal(tree1)) // [1,2,3]
    // Example 2:
    // Input: root = []
    // Output: []
    fmt.Println(preorderTraversal(nil)) // []
    // Example 3:
    // Input: root = [1]
    // Output: [1]
    tree3 := &TreeNode {
        1,
        nil,
        nil,
    }
    fmt.Println(preorderTraversal(tree3)) // [1]
}