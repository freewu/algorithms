package main

import "fmt"

// 94. Binary Tree Inorder Traversal
// Given the root of a binary tree, return the inorder traversal of its nodes' values.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/09/15/inorder_1.jpg" />
// Input: root = [1,null,2,3]
// Output: [1,3,2]

// Example 2:
// Input: root = []
// Output: []

// Example 3:
// Input: root = [1]
// Output: [1]

// Constraints:
// 		The number of nodes in the tree is in the range [0, 100].
// 		-100 <= Node.val <= 100

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
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}

// 递归
func inorder (root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		// 中序遍历
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

func main() {
	tree1 := &TreeNode {
		1,
		nil,
		&TreeNode {
			2,
			&TreeNode{3, nil, nil},
			nil,
		},
	}
	tree3 := &TreeNode {
		1,
		nil,
		nil,
	}
	fmt.Println(inorderTraversal(tree1)) // [1,3,2]
	fmt.Println(inorderTraversal(nil)) // []
	fmt.Println(inorderTraversal(tree3)) // [1]
}