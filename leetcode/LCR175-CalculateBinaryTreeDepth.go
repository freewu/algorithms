package main

// LCR 175. 计算二叉树的深度
// 某公司架构以二叉树形式记录，请返回该公司的层级数。

// 示例 1：
// <img src="https://pic.leetcode.cn/1695101942-FSrxqu-image.png" />
// 输入：root = [1, 2, 2, 3, null, null, 5, 4, null, null, 4]
// 输出: 4
// 解释: 上面示例中的二叉树的最大深度是 4，沿着路径 1 -> 2 -> 3 -> 4 或 1 -> 2 -> 5 -> 4 到达叶节点的最长路径上有 4 个节点。

// 提示：
//     节点总数 <= 10000

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
// dfs
func calculateDepth(root *TreeNode) int {
    res := 0
    var dfs func (node *TreeNode, count int)
    dfs = func (node *TreeNode, count int) {
        if node != nil {
            dfs(node.Left, count+1)
            dfs(node.Right, count+1)
        }
        if count > res {
            res = count
        }
    }
    dfs(root, res)
    return res
}

// bfs
func calculateDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    nodeList := []*TreeNode{root}
    res := 0
    for len(nodeList) > 0 {
        for _, node := range nodeList {
            nodeList = nodeList[1:]
            if node.Left != nil {
                nodeList = append(nodeList, node.Left)
            }
            if node.Right != nil {
                nodeList = append(nodeList, node.Right)
            }
        }
        res++
    }
    return res
}

// 递归
func calculateDepth2(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l, r := calculateDepth2(root.Left) + 1, calculateDepth2(root.Right) + 1
    max := func (a int, b int) int { if a > b { return a; }; return b; }
    return max(l,r)
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode { 9, nil, nil },
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(calculateDepth(tree1)) // 3

    tree2 := &TreeNode {
        1,
        nil,
        &TreeNode{2, nil, nil},
    }
    fmt.Println(calculateDepth(tree2)) // 2

    fmt.Println(calculateDepth1(tree1)) // 3
    fmt.Println(calculateDepth1(tree2)) // 2

    fmt.Println(calculateDepth2(tree1)) // 3
    fmt.Println(calculateDepth2(tree2)) // 2
}