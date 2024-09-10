package main

// 2773. Height of Special Binary Tree
// You are given a root, which is the root of a special binary tree with n nodes. T
// he nodes of the special binary tree are numbered from 1 to n. 
// Suppose the tree has k leaves in the following order: b1 < b2 < ... < bk.

// The leaves of this tree have a special property! That is, for every leaf bi, the following conditions hold:
//     The right child of bi is bi + 1 if i < k, and b1 otherwise.
//     The left child of bi is bi - 1 if i > 1, and bk otherwise.

// Return the height of the given tree.

// Note: The height of a binary tree is the length of the longest path from the root to any other node.

// Example 1:
// Input: root = [1,2,3,null,null,4,5]
// Output: 2
// Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 2.
// <img src="https://assets.leetcode.com/uploads/2023/07/12/1.png" />

// Example 2:
// Input: root = [1,2]
// Output: 1
// Explanation: The given tree is shown in the following picture. There is only one leaf, so it doesn't have any left or right child. We can see that the graph has a height of 1.
// <img src="https://assets.leetcode.com/uploads/2023/07/12/2.png" />

// Example 3:
// Input: root = [1,2,3,null,null,4,null,5,6]
// Output: 3
// Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 3.
// <img src="https://assets.leetcode.com/uploads/2023/07/12/3.png" />

// Constraints:
//     n == number of nodes in the tree
//     2 <= n <= 10^4
//     1 <= node.val <= n
//     The input is generated such that each node.val is unique.

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
func heightOfTree(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode, depht int)
    dfs = func(node *TreeNode, depht int) {
        if depht > res {
            res = depht
        }
        if node.Left != nil && node.Left.Right != node {
            dfs(node.Left, depht+1)
        }
        if node.Right != nil && node.Right.Left != node {
            dfs(node.Right, depht+1)
        }
    } 
    dfs(root, 0)
    return res
}

func main() {
    // Example 1:
    // Input: root = [1,2,3,null,null,4,5]
    // Output: 2
    // Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 2.
    // <img src="https://assets.leetcode.com/uploads/2023/07/12/1.png" />
    tree1 := &TreeNode {
        1,
        &TreeNode { 2 , nil, nil, },
        &TreeNode { 3 , &TreeNode { 4, nil , nil, },  &TreeNode { 5 , nil , nil, }, },
    }
    fmt.Println(heightOfTree(tree1)) // 2
    // Example 2:
    // Input: root = [1,2]
    // Output: 1
    // Explanation: The given tree is shown in the following picture. There is only one leaf, so it doesn't have any left or right child. We can see that the graph has a height of 1.
    // <img src="https://assets.leetcode.com/uploads/2023/07/12/2.png" />
    tree2 := &TreeNode {
        1,
        &TreeNode { 2 , nil, nil, },
        nil,
    }
    fmt.Println(heightOfTree(tree2)) // 1
    // Example 3:
    // Input: root = [1,2,3,null,null,4,null,5,6]
    // Output: 3
    // Explanation: The given tree is shown in the following picture. Each leaf's left child is the leaf to its left (shown with the blue edges). Each leaf's right child is the leaf to its right (shown with the red edges). We can see that the graph has a height of 3.
    // <img src="https://assets.leetcode.com/uploads/2023/07/12/3.png" />
    tree3 := &TreeNode {
        1,
        &TreeNode { 2 , nil, nil, },
        &TreeNode { 3 , &TreeNode { 4 , &TreeNode { 5 , nil, nil, }, &TreeNode { 6 , nil, nil, }, }, nil, },
    }
    fmt.Println(heightOfTree(tree3)) // 3
}