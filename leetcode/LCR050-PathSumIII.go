package main

// LCR 050. 路径总和 III
// 给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
// 路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg" />
// 输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// 输出：3
// 解释：和等于 8 的路径有 3 条，如图所示。

// 示例 2：
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：3
 
// 提示:
//     二叉树的节点个数的范围是 [0,1000]
//     -10^9 <= Node.val <= 10^9 
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
// dfs 
func pathSum(root *TreeNode, targetSum int) int {
    res := 0
    var dfs func(tree *TreeNode, targetSum int, pathStack []*TreeNode)
    dfs = func (tree *TreeNode, targetSum int, pathStack []*TreeNode) {
        if tree == nil {
            return
        }
        sum := 0
        pathStack = append(pathStack, tree)
        for i := len(pathStack) - 1; i >= 0; i-- {
            sum += pathStack[i].Val
            if sum == targetSum {
                res++
            }
        }
        dfs(tree.Left, targetSum, pathStack)
        dfs(tree.Right, targetSum, pathStack)
    }
    dfs(root, targetSum, []*TreeNode{})
    return res
}

func pathSum1(root *TreeNode, targetSum int) int {
    res, m := 0, map[int]int{0:1}
    var dfs func(*TreeNode, int)
    dfs = func(root *TreeNode, cur int) {
        if root == nil {
            return
        }
        cur += root.Val
        res += m[cur - targetSum]
        m[cur]++
        defer func() {m[cur]--}()
        dfs(root.Left, cur)
        dfs(root.Right, cur)
    }
    dfs(root, 0)
    return res
}

func main() {
    tree1 := &TreeNode {
        10,
        &TreeNode { 
            5, 
            &TreeNode {
                3,
                &TreeNode{3, nil, nil},
                &TreeNode{-2, nil, nil},
            },
            &TreeNode {
                2,
                nil,
                &TreeNode{1, nil, nil},
            },
        },
        &TreeNode {
            -3,
            nil,
            &TreeNode{11, nil, nil},
        },
    }
    tree2 := &TreeNode {
        5,
        &TreeNode { 
            4, 
            &TreeNode {
                11,
                &TreeNode{7, nil, nil},
                &TreeNode{2, nil, nil},
            },
            nil,
        },
        &TreeNode {
            8,
            &TreeNode{13, nil, nil},
            &TreeNode{
                4,
                &TreeNode{5, nil, nil},
                &TreeNode{1, nil, nil},
            },
        },
    }
    fmt.Println(pathSum(tree1,8)) // 3
    fmt.Println(pathSum(tree2,22)) // 3

    fmt.Println(pathSum1(tree1,8)) // 3
    fmt.Println(pathSum1(tree2,22)) // 3
}