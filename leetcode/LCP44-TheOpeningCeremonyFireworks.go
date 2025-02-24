package main

// LCP 44. 开幕式焰火
// 「力扣挑战赛」开幕式开始了，空中绽放了一颗二叉树形的巨型焰火。 
// 给定一棵二叉树root代表焰火，节点值表示巨型焰火这一位置的颜色种类。
// 请帮小扣计算巨型焰火有多少种不同的颜色。

// 示例 1：
// 输入：root = [1,3,2,1,null,2]
// 输出：3
// 解释：焰火中有 3 个不同的颜色，值分别为 1、2、3

// 示例 2：
// 输入：root = [3,3,3]
// 输出：1
// 解释：焰火中仅出现 1 个颜色，值为 3

// 提示：
//     1 <= 节点个数 <= 1000
//     1 <= Node.val <= 1000

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
func numColor(root *TreeNode) int {
    mp := make(map[int]bool)
    var dfs func(node *TreeNode) 
    dfs = func(node *TreeNode) {
        if node == nil { return }
        mp[node.Val] = true
        if node.Left != nil  { dfs(node.Left) }
        if node.Right != nil { dfs(node.Right) }
    }
    dfs(root)
    return len(mp)
}

func main() {
    // 示例 1：
    // 输入：root = [1,3,2,1,null,2]
    // 输出：3
    // 解释：焰火中有 3 个不同的颜色，值分别为 1、2、3
    tree1 := &TreeNode{
        1, 
        &TreeNode { 3, &TreeNode { 1, nil, nil, }, nil, }, 
        &TreeNode { 2, &TreeNode { 2, nil, nil, }, nil, },
    }
    fmt.Println(numColor(tree1)) // 3
    // 示例 2：
    // 输入：root = [3,3,3]
    // 输出：1
    // 解释：焰火中仅出现 1 个颜色，值为 3
    tree2 := &TreeNode{
        3, 
        &TreeNode { 3, nil, nil, }, 
        &TreeNode { 3, nil, nil, },
    }
    fmt.Println(numColor(tree2)) // 1
}