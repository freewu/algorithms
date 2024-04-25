package main

// 112. Path Sum
// Given the root of a binary tree and an integer targetSum, 
// return true if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.s
// A leaf is a node with no children.

// Example 1:
//            [5]
//           /   \
//         [4]    8
//         /     / \
//       [11]   13   4
//       / \          \
//      7   [2]        1
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg" />
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
// Output: true
// Explanation: The root-to-leaf path with the target sum is shown.

// Example 2:
//         1
//        /  \
//       2    3
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
// Input: root = [1,2,3], targetSum = 5
// Output: false
// Explanation: There two root-to-leaf paths in the tree:
// (1 --> 2): The sum is 3.
// (1 --> 3): The sum is 4.
// There is no root-to-leaf path with sum = 5.

// Example 3:
// Input: root = [], targetSum = 0
// Output: false
// Explanation: Since the tree is empty, there are no root-to-leaf paths.
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 5000].
//     -1000 <= Node.val <= 1000
//     -1000 <= targetSum <= 1000

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
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
         return false
     }
     if root.Left == nil && root.Right == nil {
         return targetSum == root.Val
     }
     return hasPathSum(root.Left, targetSum - root.Val) || hasPathSum(root.Right, targetSum - root.Val)
}

// dfs
func hasPathSum1(root *TreeNode, targetSum int) bool {
    res := false
    if root == nil {
        return res
    }
    var dfs func(root *TreeNode, targetSum, currentSum int) 
    dfs = func(root *TreeNode, targetSum, currentSum int)  {
        if root.Left == nil && root.Right == nil {
            if targetSum == currentSum+root.Val {
                res = true
            }
            return
        }
        currentSum += root.Val
        if root.Left != nil {
            dfs(root.Left, targetSum, currentSum)
        }
        if root.Right != nil {
            dfs(root.Right, targetSum, currentSum)
        }
    }
    dfs(root, targetSum, 0)
    return res
}

func main() {
    // Example 1:
    //            [5]
    //           /   \
    //         [4]    8
    //         /     / \
    //       [11]   13   4
    //       / \          \
    //      7   [2]        1
    // <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum1.jpg" />
    // Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
    // Output: true
    // Explanation: The root-to-leaf path with the target sum is shown.
    tree1 := &TreeNode {
        5,
        &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}, }, nil},
        &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}, }, },
    }
    fmt.Println(hasPathSum(tree1,22)) // true
    // Example 2:
    //         1
    //        /  \
    //       2    3
    // <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
    // Input: root = [1,2,3], targetSum = 5
    // Output: false
    // Explanation: There two root-to-leaf paths in the tree:
    // (1 --> 2): The sum is 3.
    // (1 --> 3): The sum is 4.
    // There is no root-to-leaf path with sum = 5.
    tree2 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(hasPathSum(tree2,6)) // false
    // Example 3:
    // Input: root = [], targetSum = 0
    // Output: false
    // Explanation: Since the tree is empty, there are no root-to-leaf paths.
    fmt.Println(hasPathSum(nil,0)) // false


    fmt.Println(hasPathSum1(tree1,22)) // true
    fmt.Println(hasPathSum1(tree2,6)) // false
    fmt.Println(hasPathSum1(nil,0)) // false
}