package main

// LCR 145. 判断对称二叉树
// 请设计一个函数判断一棵二叉树是否 轴对称 。

// 示例 1：
// <img src="https://pic.leetcode.cn/1694689008-JaaRdV-%E8%BD%B4%E5%AF%B9%E7%A7%B0%E4%BA%8C%E5%8F%89%E6%A0%911.png" />
// 输入：root = [6,7,7,8,9,9,8]
// 输出：true
// 解释：从图中可看出树是轴对称的。

// 示例 2：
// <img src="https://pic.leetcode.cn/1694689054-vENzHe-%E8%BD%B4%E5%AF%B9%E7%A7%B0%E4%BA%8C%E5%8F%89%E6%A0%912.png" />
// 输入：root = [1,2,2,null,3,null,3]
// 输出：false
// 解释：从图中可看出最后一层的节点不对称。

// 提示：
//     0 <= 节点个数 <= 1000

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
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
func checkSymmetricTree(root *TreeNode) bool {
    var helper func(left *TreeNode, right *TreeNode) bool 
    helper = func(left *TreeNode, right *TreeNode) bool {
        if left == nil && right == nil { return true }
        if left == nil || right == nil { return false }
        // 左右需要交替才是对称
        return left.Val == right.Val && helper(left.Right, right.Left) && helper(left.Left, right.Right)
    }
    return helper(root, root)
}

func main() {
	fmt.Println(checkSymmetricTree(
        &TreeNode {
            1,
            &TreeNode{ 2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}, },
            &TreeNode{ 2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}, },
        },
    )) // true
    fmt.Println(checkSymmetricTree(
        &TreeNode {
            1,
            &TreeNode{ 2, nil, &TreeNode{4, nil, nil}, },
            &TreeNode{ 2, nil, &TreeNode{3, nil, nil}, },
        },
    )) // true
}