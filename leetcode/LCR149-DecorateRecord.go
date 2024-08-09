package main

// LCR 149. 彩灯装饰记录 I
// 一棵圣诞树记作根节点为 root 的二叉树，节点值为该位置装饰彩灯的颜色编号。请按照从 左 到 右 的顺序返回每一层彩灯编号。

// 示例 1：
// <img src="https://pic.leetcode.cn/1694758674-XYrUiV-%E5%89%91%E6%8C%87%20Offer%2032%20-%20I_%E7%A4%BA%E4%BE%8B1.png" />
// 输入：root = [8,17,21,18,null,null,6]
// 输出：[8,17,21,18,6]

// 提示：
//     节点总数 <= 1000

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
// // dfs
// func decorateRecord(root *TreeNode) []int {
//     res := []int{}
//     var traversal func(node *TreeNode)
//     traversal = func(node *TreeNode) {
//         if node == nil { return }
//         res = append(res, node.Val)
//         if node.Left != nil { traversal(node.Left) }
//         if node.Right != nil { traversal(node.Right) }
//     }
//     traversal(root)
//     return res
// }

// bfs
func decorateRecord(root *TreeNode) []int {
    res := []int{}
    if root == nil {
        return res
    }
    queue := []*TreeNode{ root }
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]
        res = append(res, node.Val)
        if node.Left != nil  { queue = append(queue, node.Left) }
        if node.Right != nil { queue = append(queue, node.Right)  } 
    }
    return res
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode.cn/1694758674-XYrUiV-%E5%89%91%E6%8C%87%20Offer%2032%20-%20I_%E7%A4%BA%E4%BE%8B1.png" />
    // 输入：root = [8,17,21,18,null,null,6]
    // 输出：[8,17,21,18,6]
    tree1 := &TreeNode {
        8,
        &TreeNode{17, &TreeNode{18, nil, nil, }, nil, },
        &TreeNode{21, nil, &TreeNode{6, nil, nil, }, },
    }
    fmt.Println(decorateRecord(tree1)) // [8,17,21,18,6]

    fmt.Println(decorateRecord(nil)) // []
}
