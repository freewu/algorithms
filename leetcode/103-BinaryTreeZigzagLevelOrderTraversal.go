package main

// 103. Binary Tree Zigzag Level Order Traversal
// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. 
// (i.e., from left to right, then right to left for the next level and alternate between).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" / >
// Input: root = [3,9,20,null,null,15,7]
// Output: [[3],[20,9],[15,7]]

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Example 3:
// Input: root = []
// Output: []

// Constraints:
// 		The number of nodes in the tree is in the range [0, 2000].
// 		-100 <= Node.val <= 100

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
func zigzagLevelOrder(root *TreeNode) [][]int {
	var res [][]int
	// 按层序从上到下遍历一颗树
	search(root, 0, &res)
	return res
}

func search(root *TreeNode, depth int, res *[][]int) {
	if root == nil {
		return
	}
	for len(*res) < depth+1 {
		*res = append(*res, []int{})
	}
	// 每一层的顺序是相互反转的，即上一层是从左往右，下一层就是从右往左
	if depth%2 == 0 {
		(*res)[depth] = append((*res)[depth], root.Val)
	} else {
		(*res)[depth] = append([]int{root.Val}, (*res)[depth]...)
	}
	search(root.Left, depth+1, res)
	search(root.Right, depth+1, res)
}

// BFS
func zigzagLevelOrder1(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	q := []*TreeNode{root}
	size, i, j, lay, tmp, flag := 0, 0, 0, []int{}, []*TreeNode{}, false
	for len(q) > 0 {
		size = len(q)
		tmp = []*TreeNode{}
		lay = make([]int, size)
		j = size - 1
		for i = 0; i < size; i++ {
			root = q[0]
			q = q[1:]
			if !flag {
				lay[i] = root.Val
			} else {
				lay[j] = root.Val
				j--
			}
			if root.Left != nil {
				tmp = append(tmp, root.Left)
			}
			if root.Right != nil {
				tmp = append(tmp, root.Right)
			}

		}
		res = append(res, lay)
		flag = !flag
		q = tmp
	}
	return res
}


func main() {
	tree1 := &TreeNode {
		3,
		&TreeNode{9, nil, nil},
		&TreeNode {
			20,
			&TreeNode{15, nil, nil},
			&TreeNode{7, nil, nil},
		},
	}
	tree3 := &TreeNode {
		1,
		nil,
		nil,
	}
	fmt.Println(zigzagLevelOrder(tree1)) // [[3],[20,9],[15,7]]
	fmt.Println(zigzagLevelOrder(nil)) // []
	fmt.Println(zigzagLevelOrder(tree3)) // [1]

	fmt.Println(zigzagLevelOrder1(tree1)) // [[3],[20,9],[15,7]]
	fmt.Println(zigzagLevelOrder1(nil)) // []
	fmt.Println(zigzagLevelOrder1(tree3)) // [1]
}