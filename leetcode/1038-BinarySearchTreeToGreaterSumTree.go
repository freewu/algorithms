package main

// 1038. Binary Search Tree to Greater Sum Tree
// Given the root of a Binary Search Tree (BST), 
// convert it to a Greater Tree such that every key of the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.
// As a reminder, a binary search tree is a tree that satisfies these constraints:
//     The left subtree of a node contains only nodes with keys less than the node's key.
//     The right subtree of a node contains only nodes with keys greater than the node's key.
//     Both the left and right subtrees must also be binary search trees.
 
// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/05/02/tree.png" />
// Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

// Example 2:
// Input: root = [0,null,1]
// Output: [1,null,1]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     0 <= Node.val <= 100
//     All the values in the tree are unique.

import "fmt"

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
func bstToGst(root *TreeNode) *TreeNode {
    var traverse func(root *TreeNode, sum int) int 
    traverse = func (root *TreeNode, sum int) int {
        if root == nil {
            return sum
        }
        sum = traverse(root.Right, sum) // right
        root.Val += sum
        return traverse(root.Left, root.Val) // left
    }
    traverse(root, 0)
    return root
}

func main() {
    // [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
    tree1 := &TreeNode {
        4,
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, &TreeNode{3, nil, nil}, }, },
        &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, &TreeNode{8, nil, nil}, }, },
    }
    fmt.Println(tree1.Val) // 4
    fmt.Println(bstToGst(tree1)) // [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
    fmt.Println(tree1.Val) // 30

    // [0,null,1] 
    tree2 := &TreeNode {
        0,
        nil,
        &TreeNode{1, nil, nil},
    }
    fmt.Println(tree2.Val) // 0
    fmt.Println(bstToGst(tree2)) // [1,null,1]
    fmt.Println(tree2.Val) // 1
}