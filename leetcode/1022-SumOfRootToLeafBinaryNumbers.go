package main

// 1022. Sum of Root To Leaf Binary Numbers
// You are given the root of a binary tree where each node has a value 0 or 1. 
// Each root-to-leaf path represents a binary number starting with the most significant bit.
//     For example, if the path is 0 -> 1 -> 1 -> 0 -> 1, then this could represent 01101 in binary, which is 13.

// For all leaves in the tree, consider the numbers represented by the path from the root to that leaf. 
// Return the sum of these numbers.

// The test cases are generated so that the answer fits in a 32-bits integer.

// Example 1:
//             1
//           /   \
//          0     1
//        /   \  /  \
//       0    1 0    1
// <img src="https://assets.leetcode.com/uploads/2019/04/04/sum-of-root-to-leaf-binary-numbers.png" />
// Input: root = [1,0,1,0,1,0,1]
// Output: 22
// Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22

// Example 2:
// Input: root = [0]
// Output: 0

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     Node.val is 0 or 1.

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
func sumRootToLeaf(root *TreeNode) int {
    var dfs func(root *TreeNode, curr int) int 
    dfs = func(root *TreeNode, curr int) int {
        if root == nil { return 0 }
        if root.Left == nil && root.Right == nil { // leaf
            return curr + root.Val 
        }
        curr = (root.Val + curr) << 1
        return dfs(root.Left, curr) + dfs(root.Right, curr)
    }
    return dfs(root, 0)
}

func sumRootToLeaf1(root *TreeNode) int {
    var dfs func(root *TreeNode, res int) int
    dfs = func(root *TreeNode, res int) int {
        if root == nil { return 0 }
        res = res * 2 + root.Val
        if root.Left == nil && root.Right == nil { return res }
        return dfs(root.Left, res) + dfs(root.Right, res)
    }
    return dfs(root, 0)
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //          0     1
    //        /   \  /  \
    //       0    1 0    1
    // <img src="https://assets.leetcode.com/uploads/2019/04/04/sum-of-root-to-leaf-binary-numbers.png" />
    // Input: root = [1,0,1,0,1,0,1]
    // Output: 22
    // Explanation: (100) + (101) + (110) + (111) = 4 + 5 + 6 + 7 = 22
    tree1 := &TreeNode{
        1, 
        &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}, },
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(sumRootToLeaf(tree1)) // 22
    // Example 2:
    // Input: root = [0]
    // Output: 0
    tree2 := &TreeNode{ 0, nil, nil, }
    fmt.Println(sumRootToLeaf(tree2)) // 0

    fmt.Println(sumRootToLeaf1(tree1)) // 22
    fmt.Println(sumRootToLeaf1(tree2)) // 0
}