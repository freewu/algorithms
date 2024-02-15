package main

// 107. Binary Tree Level Order Traversal II
// Given the root of a binary tree, 
// return the bottom-up level order traversal of its nodes' values. (i.e., from left to right, level by level from leaf to root).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg" /> 
// Input: root = [3,9,20,null,null,15,7]
// Output: [[15,7],[9,20],[3]]

// Example 2:
// Input: root = [1]
// Output: [[1]]

// Example 3:
// Input: root = []
// Output: []
 
// Constraints:
// 		The number of nodes in the tree is in the range [0, 2000].
// 		-1000 <= Node.val <= 1000

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 反转数组
func reversalList(l [][]int) (res [][]int) {
	for i := len(l) - 1; i >= 0; i-- {
		res = append(res, l[i])
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// BFS
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	// 准备一个队列 queue 先把 root 放里面
	queue := []*TreeNode{root}
	res := make([][]int, 0)
	// 有多个层就会循环多少次
	for len(queue) > 0 {
		l := len(queue)
		tmp := make([]int, 0, l)
		for i := 0; i < l; i++ { // i < l 这里是关键
			// 左枝不为空 放到队列 queue 中
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			// 右枝不为空 放到队列 queue 中
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			// 把值加入到 tmp 中
			tmp = append(tmp, queue[i].Val)
		}
		fmt.Println("queue: ",queue)
		fmt.Println("tmp: ",tmp)
		// 取出队
		queue = queue[l:] // 这里是关键
		res = append(res, tmp)
	}
	return reversalList(res)
}

// DFS
func levelOrderBottom1(root *TreeNode) [][]int {
	var res [][]int
	var dfsLevel func(node *TreeNode, level int)
	dfsLevel = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		// 第一次进入需创建新的一行
		if len(res) == level {
			res = append(res, []int{node.Val})
		} else {
			res[level] = append(res[level], node.Val)
		}
		dfsLevel(node.Left, level+1)
		dfsLevel(node.Right, level+1)
	}
	dfsLevel(root, 0)
	return reversalList(res)
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
	fmt.Println(levelOrderBottom(tree1)) // [[15,7],[9,20],[3]]
	fmt.Println(levelOrderBottom(nil)) // []
	fmt.Println(levelOrderBottom(tree3)) // [1]

	fmt.Println(levelOrderBottom1(tree1)) // [[15,7],[9,20],[3]]
	fmt.Println(levelOrderBottom1(nil)) // []
	fmt.Println(levelOrderBottom1(tree3)) // [1]
}