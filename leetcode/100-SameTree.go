package main

import "fmt"

/*
Given a sorted linked list, delete all duplicates such that each element appear only once.

For example,
Given 1->1->2, return 1->2.
Given 1->1->2->3->3, return 1->2->3.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTreeNode(t *TreeNode) {
	if nil == t {
		return
	}

	fmt.Println()
}

func makeNodeList(nums []int) *TreeNode {
	var n = &TreeNode{-1, nil, nil}
	// var b = &ListNode{-1, n}
	// for i := 0; i < len(nums); i++ {
	// 	n.Next = &ListNode{nums[i], nil}
	// 	n = n.Next
	// }
	// return b.Next.Next
	return n
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return false
}

func main() {
	fmt.Println
}
