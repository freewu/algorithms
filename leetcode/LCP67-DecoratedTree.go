package main

// LCP 67. 装饰树
// 力扣嘉年华上的 DIY 手工展位准备了一棵缩小版的二叉装饰树root和灯饰，你需要将灯饰逐一插入装饰树中，要求如下：
//     1. 完成装饰的二叉树根结点与root的根结点值相同
//     2. 若一个节点拥有父节点，则在该节点和他的父节点之间插入一个灯饰（即插入一个值为-1的节点）。具体地：
//         2.1 在一个 父节点 x 与其左子节点 y 之间添加 -1 节点， 节点 -1、节点 y 为各自父节点的左子节点，
//         2.2 在一个 父节点 x 与其右子节点 y 之间添加 -1 节点， 节点 -1、节点 y 为各自父节点的右子节点，

// 现给定二叉树的根节点root，请返回完成装饰后的树的根节点。

// 示例 1：
// 输入：root = [7,5,6]
// 输出：[7,-1,-1,5,null,null,6]
// 解释：如下图所示，
// <img src="https://pic.leetcode-cn.com/1663575757-yRLGaq-image.png"/>

// 示例 2：
// 输入：root = [3,1,7,3,8,null,4]
// 输出：[3,-1,-1,1,null,null,7,-1,-1,null,-1,3,null,null,8,null,4]
// 解释：如下图所示
// <img src="https://pic.leetcode-cn.com/1663577920-sjrAYH-image.png"/>

// 提示：
//     0 <= root.Val <= 1000 root节点数量范围为[1, 10^5]

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
func expandBinaryTree1(root *TreeNode) *TreeNode {
    var dfs func(root, parent *TreeNode, flag bool) 
    dfs = func(root, parent *TreeNode, flag bool) {
        if root == nil { return }
        if parent != nil {
            if flag {
                parent.Left  = &TreeNode{ Val: -1, Left:  root }
            } else {
                parent.Right = &TreeNode{ Val: -1, Right: root }
            }
        }
        dfs(root.Left, root, true)
        dfs(root.Right, root, false)
    }
    dfs(root, nil, true)
    return root
}

func expandBinaryTree(node *TreeNode) *TreeNode {
    if node == nil { return nil }
    l := expandBinaryTree(node.Left)
    if l != nil {
        node.Left = &TreeNode{ Val: -1, Left: l }
    }
    r := expandBinaryTree(node.Right)
    if r != nil {
        node.Right = &TreeNode{ Val:-1, Right: r }
    }
    return node
}

func main() {
    // 示例 1：
    // 输入：root = [7,5,6]
    // 输出：[7,-1,-1,5,null,null,6]
    // 解释：如下图所示，
    // <img src="https://pic.leetcode-cn.com/1663575757-yRLGaq-image.png"/>
    tree1 := &TreeNode {
        7,
        &TreeNode{5, nil, nil, },
        &TreeNode{6, nil, nil, },
    }
    fmt.Println(expandBinaryTree(tree1)) // 
    // 示例 2：
    // 输入：root = [3,1,7,3,8,null,4]
    // 输出：[3,-1,-1,1,null,null,7,-1,-1,null,-1,3,null,null,8,null,4]
    // 解释：如下图所示
    // <img src="https://pic.leetcode-cn.com/1663577920-sjrAYH-image.png"/>
    tree2 := &TreeNode {
        3,
        &TreeNode{1, &TreeNode{3, nil, nil, }, &TreeNode{8, nil, nil, }, },
        &TreeNode{7, nil,                      &TreeNode{4, nil, nil, }, },
    }
    fmt.Println(expandBinaryTree(tree2)) // 

    tree11 := &TreeNode {
        7,
        &TreeNode{5, nil, nil, },
        &TreeNode{6, nil, nil, },
    }
    fmt.Println(expandBinaryTree1(tree11)) // 
    tree12 := &TreeNode {
        3,
        &TreeNode{1, &TreeNode{3, nil, nil, }, &TreeNode{8, nil, nil, }, },
        &TreeNode{7, nil,                      &TreeNode{4, nil, nil, }, },
    }
    fmt.Println(expandBinaryTree1(tree12)) // 
}