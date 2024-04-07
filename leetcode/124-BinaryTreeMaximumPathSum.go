package main

// 124. Binary Tree Maximum Path Sum
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them. A node can only appear in the sequence at most once. 
// Note that the path does not need to pass through the root.
// The path sum of a path is the sum of the node's values in the path.
// Given the root of a binary tree, return the maximum path sum of any non-empty path.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg" / >
//     1
//    /  \ 
//   2    3
// Input: root = [1,2,3]
// Output: 6
// Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg" / >
//         -10
//         /  \
//        9    20
//            /  \
//           15   7
// Input: root = [-10,9,20,null,null,15,7]
// Output: 42
// Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.

// Constraints:
//     The number of nodes in the tree is in the range [1, 3 * 10^4].
//     -1000 <= Node.val <= 1000

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
func maxPathSum(root *TreeNode) int {
    res := -1 << 63
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(root *TreeNode, res *int) int
    dfs = func(root *TreeNode, res *int) int {
        if root == nil {
            return 0
        }
        leftSum := dfs(root.Left, res)
        rightSum := dfs(root.Right, res)
        *res = max(*res, leftSum + rightSum + root.Val)
        return max(max(leftSum, rightSum) + root.Val, 0)
    }
    dfs(root, &res)
    return res
}

func maxPathSum1(root *TreeNode) int {
    res := -1 << 63
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(r *TreeNode) int
    dfs = func(r *TreeNode) int {
        if r == nil {
            return 0
        }
        left := dfs(r.Left)
        right := dfs(r.Right)
        sum := left + right + r.Val
        res = max(res, sum)
        return max(max(left, right) + r.Val, 0)
    }
    dfs(root)
    return res
}

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode { 2, nil, nil },
        &TreeNode { 3, nil, nil},
    }
    tree2 := &TreeNode {
        -10,
        &TreeNode { 9, nil, nil },
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    // Explanation: The optimal path is 2 -> 1 -> 3 with a path sum of 2 + 1 + 3 = 6.
    fmt.Println(maxPathSum(tree1)) // 6
    // Explanation: The optimal path is 15 -> 20 -> 7 with a path sum of 15 + 20 + 7 = 42.
    fmt.Println(maxPathSum(tree2)) // 42

    fmt.Println(maxPathSum1(tree1)) // 6
    fmt.Println(maxPathSum1(tree2)) // 42
}