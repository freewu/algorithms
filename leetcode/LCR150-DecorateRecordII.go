package main

// LCR 150. 彩灯装饰记录 II
// 一棵圣诞树记作根节点为 root 的二叉树，节点值为该位置装饰彩灯的颜色编号。
// 请按照从左到右的顺序返回每一层彩灯编号，每一层的结果记录于一行。

// 示例 1：
// <img src="https://pic.leetcode.cn/1694758674-XYrUiV-%E5%89%91%E6%8C%87%20Offer%2032%20-%20I_%E7%A4%BA%E4%BE%8B1.png" />
// 输入：root = [8,17,21,18,null,null,6]
// 输出：[[8],[17,21],[18,6]]

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
// BFS
func decorateRecord(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        return res
    }
     
    queue := []*TreeNode{ root } // 准备一个队列 queue 先把 root 放里面
    // 有多个层就会循环多少次
    for len(queue) > 0 {
        l := len(queue)
        tmp := make([]int, 0, l)
        for i := 0; i < l; i++ { // i < l 这里是关键
            // 左枝不为空 放到队列 queue 中
            if queue[i].Left != nil {
                queue = append(queue, queue[i].Left)
            }
            // 右枝不为空 放到队列 queue 中
            if queue[i].Right != nil {
                queue = append(queue, queue[i].Right)
            }
            // 把值加入到 tmp 中
            tmp = append(tmp, queue[i].Val)
        }
        // 取出队
        queue = queue[l:] // 这里是关键
        res = append(res, tmp)
    }
    return res
}

// DFS
func decorateRecord1(root *TreeNode) [][]int {
    res := [][]int{}
    var dfs func(node *TreeNode, level int)
    dfs = func(node *TreeNode, level int) {
        if node == nil {
            return
        }
        // 第一次进入需创建新的一行
        if len(res) == level {
            res = append(res, []int{node.Val})
        } else {
            res[level] = append(res[level], node.Val)
        }
        dfs(node.Left, level+1)
        dfs(node.Right, level+1)
    }
    dfs(root, 0)
    return res
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode{ 9, nil, nil},
        &TreeNode{ 20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}, },
    }
    tree3 := &TreeNode { 1, nil, nil, }
    fmt.Println(levelOrder(tree1)) // [[3],[9,20],[15,7]]
    fmt.Println(levelOrder(nil)) // []
    fmt.Println(levelOrder(tree3)) // [1]

    fmt.Println(levelOrder1(tree1)) // [[3],[9,20],[15,7]]
    fmt.Println(levelOrder1(nil)) // []
    fmt.Println(levelOrder1(tree3)) // [1]
}