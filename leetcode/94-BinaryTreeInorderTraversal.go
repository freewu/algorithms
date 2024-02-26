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
// 递归
func inorderTraversal(root *TreeNode) []int {
	var res []int
	inorder(root, &res)
	return res
}

func inorder (root *TreeNode, output *[]int) {
	if root != nil {
		inorder(root.Left, output)
		// 中序遍历
		*output = append(*output, root.Val)
		inorder(root.Right, output)
	}
}

// 迭代
func inorderTraversal1(root *TreeNode) (res []int) {
	stack := []*TreeNode{}
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, root.Val)
		root = root.Right
	}
	return
}

// Morris 中序遍历
func inorderTraversal2(root *TreeNode) (res []int) {
	for root != nil {
		if root.Left != nil {
			// predecessor 节点表示当前 root 节点向左走一步，然后一直向右走至无法走为止的节点
			predecessor := root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				// 有右子树且没有设置过指向 root，则继续向右走
				predecessor = predecessor.Right
			}
			if predecessor.Right == nil {
				// 将 predecessor 的右指针指向 root，这样后面遍历完左子树 root.Left 后，就能通过这个指向回到 root
				predecessor.Right = root
				// 遍历左子树
				root = root.Left
			} else { // predecessor 的右指针已经指向了 root，则表示左子树 root.Left 已经访问完了
				res = append(res, root.Val)
				// 恢复原样
				predecessor.Right = nil
				// 遍历右子树
				root = root.Right
			}
		} else { // 没有左子树
			res = append(res, root.Val)
			// 若有右子树，则遍历右子树
			// 若没有右子树，则整颗左子树已遍历完，root 会通过之前设置的指向回到这颗子树的父节点
			root = root.Right
		}
	}
	return
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

	fmt.Println(inorderTraversal1(tree1)) // [1,3,2]
	fmt.Println(inorderTraversal1(nil)) // []
	fmt.Println(inorderTraversal1(tree3)) // [1]

	fmt.Println(inorderTraversal2(tree1)) // [1,3,2]
	fmt.Println(inorderTraversal2(nil)) // []
	fmt.Println(inorderTraversal2(tree3)) // [1]
}