package main

// 面试题 04.02. Minimum Height Tree LCCI
// Given a sorted (increasing order) array with unique integer elements, write an algo­rithm to create a binary search tree with minimal height.

// Example:
// Given sorted array: [-10,-3,0,5,9],
// One possible answer is: [0,-3,9,-10,null,5]，which represents the following tree: 
//           0 
//          / \ 
//        -3   9 
//        /   / 
//      -10  5 

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
func sortedArrayToBST(nums []int) *TreeNode {
    var build func(nums[]int ,left,right int) *TreeNode
    build = func(nums[]int ,left,right int) *TreeNode {
        if left > right { return nil }
        mid := (left + right) >> 1 // 总是选择中间位置右边的数字作为根节点
        tree := &TreeNode{ Val: nums[mid] }
        tree.Left, tree.Right = build(nums, left, mid - 1), build(nums, mid + 1, right)
        return tree
    }
    return build(nums, 0, len(nums) - 1)
}

func main() {
    // Example:
    // Given sorted array: [-10,-3,0,5,9],
    // One possible answer is: [0,-3,9,-10,null,5]，which represents the following tree: 
    //           0 
    //          / \ 
    //        -3   9 
    //        /   / 
    //      -10  5 
    fmt.Println(sortedArrayToBST([]int{-10,-3,0,5,9})) // &{0 0xc000008060 0xc000008090}

    fmt.Println(sortedArrayToBST([]int{1,2,3,4,5,6,7,8,9})) // &{5 0xc0000940d8 0xc000094138}
    fmt.Println(sortedArrayToBST([]int{9,8,7,6,5,4,3,2,1})) // &{5 0xc0000940d8 0xc000094138}
}