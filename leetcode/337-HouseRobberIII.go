package main

// 337. House Robber III
// The thief has found himself a new place for his thievery again. 
// There is only one entrance to this area, called root.

// Besides the root, each house has one and only one parent house. 
// After a tour, the smart thief realized that all houses in this place form a binary tree. 
// It will automatically contact the police if two directly-linked houses were broken into on the same night.

// Given the root of the binary tree, return the maximum amount of money the thief can rob without alerting the police.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/10/rob1-tree.jpg" />
// Input: root = [3,2,3,null,3,null,1]
// Output: 7
// Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/03/10/rob2-tree.jpg" />
// Input: root = [3,4,5,1,3,null,1]
// Output: 9
// Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     0 <= Node.val <= 10^4

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
func rob(root *TreeNode) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(*TreeNode) (int, int)
    dfs = func(root *TreeNode) (int, int) { // 返回值 偷当前节点的收益，不偷当前节点的收益
        if root == nil {
            return 0,0
        }
        withRootLeft, withoutRootLeft := dfs(root.Left)
        withRootRight, withoutRootRight := dfs(root.Right)
        withRoot := root.Val + withoutRootLeft + withoutRootRight // 不能偷相邻
        withoutRoot := max(withRootLeft, withoutRootLeft) + max(withRootRight, withoutRootRight) // 取最大为优
        return withRoot, withoutRoot
    }
    withRoot, withoutRoot := dfs(root)
    return max(withRoot, withoutRoot)
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode{2, nil, &TreeNode{3, nil, nil}, },
        &TreeNode{3, nil, &TreeNode{1, nil, nil}, },
    }
    tree2 := &TreeNode {
        3,
        &TreeNode{4,  &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{5, nil, &TreeNode{1, nil, nil}, },
    }
    fmt.Println(rob(tree1)) // 7
    fmt.Println(rob(tree2)) // 9
}