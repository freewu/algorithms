package main

// 776. Split BST
// Given the root of a binary search tree (BST) and an integer target, 
// split the tree into two subtrees where the first subtree has nodes that are all smaller or equal to the target value, 
// while the second subtree has all nodes that are greater than the target value. 
// It is not necessarily the case that the tree contains a node with the value target.

// Additionally, most of the structure of the original tree should remain. 
// Formally, for any child c with parent p in the original tree, 
// if they are both in the same subtree after the split, then node c should still have the parent p.

// Return an array of the two roots of the two subtrees in order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/06/13/split-tree.jpg" />
// Input: root = [4,2,6,1,3,5,7], target = 2
// Output: [[2,1],[4,3,6,null,null,5,7]]

// Example 2:
// Input: root = [1], target = 1
// Output: [[1],[]]

// Constraints:
//     The number of nodes in the tree is in the range [1, 50].
//     0 <= Node.val, target <= 1000

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
// 递归
func splitBST(root *TreeNode, target int) []*TreeNode {
    res := make([]*TreeNode, 2)
    if root == nil {
        return res
    }
    if root.Val <= target {
        res = splitBST(root.Right, target)
        root.Right = res[0]
        res[0] = root
    } else {
        res = splitBST(root.Left, target)
        root.Left = res[1]
        res[1] = root
    }
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/06/13/split-tree.jpg" />
    // Input: root = [4,2,6,1,3,5,7], target = 2
    // Output: [[2,1],[4,3,6,null,null,5,7]]
    tree1 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(splitBST(tree1, 2))
    // Example 2:
    // Input: root = [1], target = 1
    // Output: [[1],[]]
    tree2 := &TreeNode{1, nil, nil}
    fmt.Println(splitBST(tree2, 1))
}