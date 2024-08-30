package main

// 951. Flip Equivalent Binary Trees
// For a binary tree T, we can define a flip operation as follows: 
//     choose any node, and swap the left and right child subtrees.

// A binary tree X is flip equivalent to a binary tree Y if and only if we can make X equal to Y after some number of flip operations.

// Given the roots of two binary trees root1 and root2, return true if the two trees are flip equivalent or false otherwise.

// Example 1:
//               1                1
//           /       \         /     \ 
//          2         3       3       2
//        /   \       /        \     /   \
//       4     5     6          6   4     5
//            /  \                       /   \
//           7    8                     8     7
// <img src="https://assets.leetcode.com/uploads/2018/11/29/tree_ex.png" />
// Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
// Output: true
// Explanation: We flipped at nodes with values 1, 3, and 5.

// Example 2:
// Input: root1 = [], root2 = []
// Output: true

// Example 3:
// Input: root1 = [], root2 = [1]
// Output: false

// Constraints:
//     The number of nodes in each tree is in the range [0, 100].
//     Each tree will have unique node values in the range [0, 99].

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
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
    if root1 == nil && root2 == nil { return true }
    if root1 == nil || root2 == nil || root1.Val != root2.Val  { return false }
    l1, r1, l2, r2 := root1.Left, root1.Right, root2.Left, root2.Right
    // 不翻转、翻转两种情况满足一种即可算是匹配
    if flipEquiv(l1, l2) && flipEquiv(r1, r2) { return true }
    if flipEquiv(l1, r2) && flipEquiv(r1, l2) { return true }
    return false 
}

func main() {
    // Example 1:
    //               1                1
    //           /       \         /     \ 
    //          2         3       3       2
    //        /   \       /        \     /   \
    //       4     5     6          6   4     5
    //            /  \                       /   \
    //           7    8                     8     7
    // <img src="https://assets.leetcode.com/uploads/2018/11/29/tree_ex.png" />
    // Input: root1 = [1,2,3,4,5,6,null,null,null,7,8], root2 = [1,3,2,null,6,4,5,null,null,null,null,8,7]
    // Output: true
    // Explanation: We flipped at nodes with values 1, 3, and 5.
    tree11 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode{ 4, nil, nil}, &TreeNode{ 5, &TreeNode{ 7, nil, nil}, &TreeNode{ 8, nil, nil}, }, },
        &TreeNode { 3, &TreeNode{ 6, nil, nil}, nil,                    },
    }
    tree12 := &TreeNode {
        1,
        &TreeNode { 3, nil,                     &TreeNode{ 6, nil, nil},},
        &TreeNode { 2, &TreeNode{ 4, nil, nil}, &TreeNode{ 5, &TreeNode{ 8, nil, nil}, &TreeNode{ 7, nil, nil}, }, },
        
    }
    fmt.Println(flipEquiv(tree11, tree12)) // true
    // Example 2:
    // Input: root1 = [], root2 = []
    // Output: true
    fmt.Println(flipEquiv(nil, nil)) // true
    // Example 3:
    // Input: root1 = [], root2 = [1]
    // Output: false
    fmt.Println(flipEquiv(nil, &TreeNode{ 1, nil, nil})) // false
}