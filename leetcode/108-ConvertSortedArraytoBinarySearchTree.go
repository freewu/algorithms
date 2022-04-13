package main

import "fmt"

/**
108. Convert Sorted Array to Binary Search Tree
Given an integer array nums where the elements are sorted in ascending order, convert it to a height-balanced binary search tree.
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.

Constraints:

	1 <= nums.length <= 10^4
	-10^4 <= nums[i] <= 10^4
	nums is sorted in a strictly increasing order.

Example 1:

	Input: nums = [-10,-3,0,5,9]
	Output: [0,-3,9,-10,null,5]
	Explanation: [0,-10,5,null,-3,null,9] is also accepted:

		  0             0
		 / \		   /  \
	   -3   9   =>  -10    5
	   /   /          \      \
	 -10  5            -3     9

Example 2:

	Input: nums = [1,3]
	Output: [3,1]
	Explanation: [1,null,3] and [3,1] are both height-balanced BSTs.

	 3     1
	/	=>  \
   1         3

解题思路:
	把一个有序数组转换成高度平衡的二叉搜索数
 */

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
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val: nums[len(nums)/2], // 取中间值
		Left: sortedArrayToBST(nums[:len(nums)/2]), // 取左半部分 递归
		Right: sortedArrayToBST(nums[len(nums)/2+1:]), // 取右半部分 递归
	}
}

// best solution
func sortedArrayToBSTBest(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 { // 多了这部分 如果只有一个值时直接返回无叶子的节点，减少递归调用
		return &TreeNode{Val: nums[0]}
	}
	mid := len(nums)/2
	root := &TreeNode{
		Val: nums[mid],
		Left: sortedArrayToBSTBest(nums[:mid]),
		Right: sortedArrayToBSTBest(nums[mid+1:]),
	}
	return root
}

func main() {
	fmt.Printf("sortedArrayToBST([]int{-10,-3,0,5,9}) = %v\n",sortedArrayToBST([]int{-10,-3,0,5,9}))
	fmt.Printf("sortedArrayToBST([]int{13})= %v\n",sortedArrayToBST([]int{13}))

	fmt.Printf("sortedArrayToBSTBest([]int{-10,-3,0,5,9}) = %v\n",sortedArrayToBSTBest([]int{-10,-3,0,5,9}))
	fmt.Printf("sortedArrayToBSTBest([]int{13})= %v\n",sortedArrayToBSTBest([]int{13}))
}
