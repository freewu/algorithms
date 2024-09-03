package main

// LCR 051. 二叉树中的最大路径和
// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
// 同一个节点在一条路径序列中 至多出现一次 。
// 该路径 至少包含一个 节点，且不一定经过根节点。

// 路径和 是路径中各节点值的总和。

// 给定一个二叉树的根节点 root ，返回其 最大路径和，即所有路径上节点值之和的最大值。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg" />
// 输入：root = [1,2,3]
// 输出：6
// 解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg" />
// 输入：root = [-10,9,20,null,null,15,7]
// 输出：42
// 解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42

// 提示：
//     树中节点数目范围是 [1, 3 * 10^4]
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

func maxPathSum2(root *TreeNode) int {
    if root == nil { return 0 }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node *TreeNode) (int, int) // @return 最大路径和 + 最大单条路径和
    dfs = func(node *TreeNode) (int, int) {
        if node == nil { return -1 << 31, 0 }
        lres, lPathSum := dfs(node.Left)
        rres, rPathSum := dfs(node.Right)
        return max(max(lres, rres), lPathSum + rPathSum+node.Val),  max(0, max(lPathSum, rPathSum) + node.Val)
    }
    res, _ := dfs(root)
    return res
}

func maxPathSum3(root *TreeNode) int {
    res := -1 << 31
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0 }
        l, r := max(dfs(node.Left), 0), max(dfs(node.Right), 0)
        sum := node.Val + l + r
        res = max(res, sum)
        return node.Val + max(l, r)
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

    fmt.Println(maxPathSum2(tree1)) // 6
    fmt.Println(maxPathSum2(tree2)) // 42

    fmt.Println(maxPathSum3(tree1)) // 6
    fmt.Println(maxPathSum3(tree2)) // 42
}