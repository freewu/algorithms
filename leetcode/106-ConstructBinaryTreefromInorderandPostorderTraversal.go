package main

// 106. Construct Binary Tree from Inorder and Postorder Traversal
// Given two integer arrays inorder and postorder where inorder is the inorder traversal of a binary tree and postorder is the postorder traversal of the same tree, construct and return the binary tree.

// Constraints:
// 		1 <= inorder.length <= 3000
// 		postorder.length == inorder.length
// 		-3000 <= inorder[i], postorder[i] <= 3000
// 		inorder and postorder consist of unique values.
// 		Each value of postorder also appears in inorder.
// 		inorder is guaranteed to be the inorder traversal of the tree.
// 		postorder is guaranteed to be the postorder traversal of the tree.

// Example 1:
//     3
//    / \
//   9  20
//     /  \
//    15   7
// 	Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
// 	Output: [3,9,20,null,null,15,7]

// Example 2:
// 	Input: inorder = [-1], postorder = [-1]
// 	Output: [-1]

// 解题思路:
// 	根据一棵树的中序遍历与后序遍历构造二叉树

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 直接传入需要的 slice 范围作为输入, 可以避免申请对应 inorder 索引的内存,
func buildTree(inorder []int, postorder []int) *TreeNode {
	postorderLen := len(postorder)
	if len(inorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: postorder[postorderLen-1]}
	postorder = postorder[:postorderLen-1]
	for pos, node := range inorder {
		if node == root.Val {
			root.Left = buildTree(inorder[:pos], postorder[:len(inorder[:pos])])
			root.Right = buildTree(inorder[pos+1:], postorder[len(inorder[:pos]):])
		}
	}
	return root
}

// 
func buildTree1(inorder []int, postorder []int) *TreeNode {
	inPos := make(map[int]int)
	for i := 0; i < len(inorder); i++ {
		inPos[inorder[i]] = i
	}
	return buildInPos2TreeDFS(postorder, 0, len(postorder)-1, 0, inPos)
}

func buildInPos2TreeDFS(post []int, postStart int, postEnd int, inStart int, inPos map[int]int) *TreeNode {
	if postStart > postEnd {
		return nil
	}
	root := &TreeNode{Val: post[postEnd]}
	rootIdx := inPos[post[postEnd]]
	leftLen := rootIdx - inStart
	root.Left = buildInPos2TreeDFS(post, postStart, postStart+leftLen-1, inStart, inPos)
	root.Right = buildInPos2TreeDFS(post, postStart+leftLen, postEnd-1, rootIdx+1, inPos)
	return root
}

// best solution
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(postorder) == 0 {
		return nil
	}

	root := &TreeNode{postorder[len(postorder)-1], nil, nil}

	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			break
		}
	}

	root.Left = buildTree(inorder[0:i], postorder[0:i])
	root.Right = buildTree(inorder[i+1:], postorder[i:len(postorder)-1])
	return root
}

func main() {
	fmt.Printf("buildTree([]int{9,3,15,20,7},[]int{9,15,7,20,3}) = %v\n",buildTree([]int{9,3,15,20,7},[]int{9,15,7,20,3})) // [3,9,20,null,null,15,7]
	fmt.Printf("buildTree1([]int{9,3,15,20,7},[]int{9,15,7,20,3}) = %v\n",buildTree1([]int{9,3,15,20,7},[]int{9,15,7,20,3})) // [3,9,20,null,null,15,7]

	fmt.Printf("buildTree([]int{-1},[]int{-1}) = %v\n",buildTree([]int{-1},[]int{-1})) // [-1]
	fmt.Printf("buildTree1([]int{-1},[]int{-1}) = %v\n",buildTree1([]int{-1},[]int{-1})) // [-1]
}
