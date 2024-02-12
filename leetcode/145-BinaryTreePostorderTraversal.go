package main

// 145. Binary Tree Postorder Traversal
// Given the root of a binary tree, return the postorder traversal of its nodes' values.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/08/28/pre1.jpg" />
// Input: root = [1,null,2,3]
// Output: [3,2,1]

// Example 2:
// Input: root = []
// Output: []

// Example 3:
// Input: root = [1]
// Output: [1]
 
// Constraints:
// 		The number of the nodes in the tree is in the range [0, 100].
// 		-100 <= Node.val <= 100
		
// Follow up: Recursive solution is trivial, could you do it iteratively?

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
// 递归处理
func postorderTraversal(root *TreeNode) []int {
	var res []int
	postorder(root, &res)
	return res
}

// 递归
func postorder (root *TreeNode, output *[]int) {
	if root != nil {
		postorder(root.Left, output)
		postorder(root.Right, output)
		// 后序遍历
		*output = append(*output, root.Val)
	}
}

// 迭代处理
func postorderTraversal1(root *TreeNode) (res []int) {
	// 递归的时候隐式地维护了一个栈，而在迭代的时候需要显式地将这个栈模拟出来
    stack := []*TreeNode{}
    var prev *TreeNode
    for root != nil || len(stack) > 0 {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        root = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if root.Right == nil || root.Right == prev {
            res = append(res, root.Val)
            prev = root
            root = nil
        } else {
            stack = append(stack, root)
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
	fmt.Println(postorderTraversal(tree1)) // [3,2,1]
	fmt.Println(postorderTraversal(nil)) // []
	fmt.Println(postorderTraversal(tree3)) // [1]

	fmt.Println(postorderTraversal1(tree1)) // [3,2,1]
	fmt.Println(postorderTraversal1(nil)) // []
	fmt.Println(postorderTraversal1(tree3)) // [1]
}