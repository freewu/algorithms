package main

// LCR 049. 求根节点到叶节点数字之和
// 给定一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。

// 每条从根节点到叶节点的路径都代表一个数字：
//     例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。

// 计算从根节点到叶节点生成的 所有数字之和 。
// 叶节点 是指没有子节点的节点。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2021/02/19/num1tree.jpg" />
// 输入：root = [1,2,3]
// 输出：25
// 解释：
// 从根到叶子节点路径 1->2 代表数字 12
// 从根到叶子节点路径 1->3 代表数字 13
// 因此，数字总和 = 12 + 13 = 25

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2021/02/19/num2tree.jpg" />
// 输入：root = [4,9,0,5,1]
// 输出：1026
// 解释：
// 从根到叶子节点路径 4->9->5 代表数字 495
// 从根到叶子节点路径 4->9->1 代表数字 491
// 从根到叶子节点路径 4->0 代表数字 40
// 因此，数字总和 = 495 + 491 + 40 = 1026

// 提示：
//     树中节点的数目在范围 [1, 1000] 内
//     0 <= Node.val <= 9
//     树的深度不超过 10

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func printTreeNode(t *TreeNode) {
    if nil == t {
        return
    }
    fmt.Println()
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    var dfs func(root *TreeNode, sum int) int 
    dfs = func(root *TreeNode, sum int) int {
        if root == nil {
            return 0
        }
        sum = sum * 10 + root.Val // 每递归一层 * 10
        if root.Left == nil && root.Right == nil {
            return sum
        } else {
            return dfs(root.Left, sum) + dfs(root.Right, sum)
        }
    }
    return dfs(root, 0)
}

func sumNumbers1(root *TreeNode) int {
    var dfs func(root *TreeNode, num int) int
    dfs = func(root *TreeNode, num int) int {
        if root.Left == nil && root.Right == nil {
            return num * 10 + root.Val
        }
        sum := 0
        if root.Left != nil  {
            sum += dfs(root.Left, num * 10 + root.Val)
        }
        if root.Right != nil {
            sum += dfs(root.Right, num * 10 + root.Val)
        }
        return sum
    }
    return dfs(root, 0)
}

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    tree2 := &TreeNode {
        4,
        &TreeNode{
            9, 
            &TreeNode{5, nil, nil}, 
            &TreeNode{1, nil, nil},
        },
        &TreeNode{0, nil, nil},
    }
    // The root-to-leaf path 1->2 represents the number 12.
    // The root-to-leaf path 1->3 represents the number 13.
    // Therefore, sum = 12 + 13 = 25.
    fmt.Println(sumNumbers(tree1)) // 25
    // The root-to-leaf path 4->9->5 represents the number 495.
    // The root-to-leaf path 4->9->1 represents the number 491.
    // The root-to-leaf path 4->0 represents the number 40.
    // Therefore, sum = 495 + 491 + 40 = 1026.
    fmt.Println(sumNumbers(tree2)) // 1026

    fmt.Println(sumNumbers1(tree1)) // 25
    fmt.Println(sumNumbers1(tree2)) // 1026
}