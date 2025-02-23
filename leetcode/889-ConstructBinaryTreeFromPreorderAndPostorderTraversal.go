package main

// 889. Construct Binary Tree from Preorder and Postorder Traversal 
// Given two integer arrays, preorder and postorder where preorder is the preorder traversal of a binary tree of distinct values and postorder is the postorder traversal of the same tree, reconstruct and return the binary tree.
// If there exist multiple answers, you can return any of them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/07/24/lc-prepost.jpg" />
// Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
// Output: [1,2,3,4,5,6,7]

// Example 2:
// Input: preorder = [1], postorder = [1]
// Output: [1]
 
// Constraints:
//     1 <= preorder.length <= 30
//     1 <= preorder[i] <= preorder.length
//     All the values of preorder are unique.
//     postorder.length == preorder.length
//     1 <= postorder[i] <= postorder.length
//     All the values of postorder are unique.
//     It is guaranteed that preorder and postorder are the preorder traversal and postorder traversal of the same binary tree.

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
func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
    n := len(postorder)
    if n == 0 { return nil }
    var build func(preorder []int,  pres int,  pree int, postorder []int,  posts int,  poste int) *TreeNode
    build = func(preorder []int,  pres int,  pree int, postorder []int,  posts int,  poste int) *TreeNode {
        if pres > pree { return nil }
        if pres == pree {
            return  &TreeNode{ Val: preorder[pres], Left : nil, Right : nil }
        }
        rootval, leftval, index := preorder[pres], preorder[pres + 1], -1
        for i := posts; i <= poste; i = i + 1 {
            if postorder[i] == leftval {
                index = i
                break
            }
        }
        size := index - posts + 1
        node := &TreeNode {
            Val: rootval,
            Left : build(preorder, pres + 1, pres + size, postorder, posts, index),
            Right : build(preorder, pres + size + 1, pree,postorder, index + 1, poste - 1),
        }
        return node
    }
    // 由于这无法确定唯一的一颗二叉树，所以我们假设排除根节点后，左边第一个节点是左子树的根节点，然后进行递归操作
    return build(preorder, 0, n - 1, postorder, 0, n - 1)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/07/24/lc-prepost.jpg" />
    // Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
    // Output: [1,2,3,4,5,6,7]
    fmt.Printf("%v\n",constructFromPrePost([]int{1,2,4,5,3,6,7},[]int{4,5,2,6,7,3,1}))
    // Example 2:
    // Input: preorder = [1], postorder = [1]
    // Output: [1]
    fmt.Printf("%v\n",constructFromPrePost([]int{1},[]int{1}))
}