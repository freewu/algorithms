package main

// LCR 151. 彩灯装饰记录 III
// 一棵圣诞树记作根节点为 root 的二叉树，节点值为该位置装饰彩灯的颜色编号。请按照如下规则记录彩灯装饰结果：
//     第一层按照从左到右的顺序记录
//     除第一层外每一层的记录顺序均与上一层相反。即第一层为从左到右，第二层为从右到左。

// 示例 1：
// <img src="https://pic.leetcode.cn/1694758674-XYrUiV-%E5%89%91%E6%8C%87%20Offer%2032%20-%20I_%E7%A4%BA%E4%BE%8B1.png" />
// 输入：root = [8,17,21,18,null,null,6]
// 输出：[[8],[21,17],[18,6]]

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
// bfs
func decorateRecord(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
    queue, flag :=[]*TreeNode{ root }, true
    for len(queue) > 0 {
        vals, index := make([]int, len(queue)), 0
        if !flag {
            index = len(vals) - 1
        }
        for i := 0; i < len(vals); i++ {
            node := queue[i]
            vals[index] = node.Val
            if !flag {
                index --
            } else {
                index++
            }
            if node.Left != nil {   queue = append(queue, node.Left) }
            if node.Right != nil {  queue = append(queue, node.Right) }
        }
        res = append(res,vals) 
        flag = !flag 
        queue = queue[len(vals):]
    }
    return res
}

func main() {
    // 示例 1：
    // <img src="https://pic.leetcode.cn/1694758674-XYrUiV-%E5%89%91%E6%8C%87%20Offer%2032%20-%20I_%E7%A4%BA%E4%BE%8B1.png" />
    // 输入：root = [8,17,21,18,null,null,6]
    // 输出：[[8],[21,17],[18,6]]
    tree1 := &TreeNode {
        8,
        &TreeNode{17, &TreeNode{18, nil, nil, }, nil, },
        &TreeNode{21, nil, &TreeNode{6, nil, nil, }, },
    }
    fmt.Println(decorateRecord(tree1)) // [[8],[21,17],[18,6]]
    fmt.Println(decorateRecord(nil)) // []
}
