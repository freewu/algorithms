package main

// LCR 144. 翻转二叉树
// 给定一棵二叉树的根节点 root，请左右翻转这棵二叉树，并返回其根节点。

// 示例 1：
//          5                   5
//        /   \               /   \
//       7     9     =>      9     7
//      /  \  /  \         /  \   /  \
//     8   3  2   4       4    2 3    9
// <img src="https://pic.leetcode.cn/1694686821-qlvjod-%E7%BF%BB%E8%BD%AC%E4%BA%8C%E5%8F%89%E6%A0%91.png" />
// 输入：root = [5,7,9,8,3,2,4]
// 输出：[5,9,7,4,2,3,8]

// 提示：
//     树中节点数目范围在 [0, 100] 内
//     -100 <= Node.val <= 100

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
func mirrorTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    // General cases:
    // invert child node of current root
    root.Left, root.Right = root.Right, root.Left   
    // invert subtree with DFS
    mirrorTree(root.Left)
    mirrorTree(root.Right)
    return root
}

func mirrorTree1(root *TreeNode) *TreeNode {
    if root == nil{
        return nil
    }
    root.Left, root.Right = mirrorTree1(root.Right), mirrorTree1(root.Left)
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/invert1-tree.jpg" />
    // Input: root = [4,2,7,1,3,6,9]
    // Output: [4,7,2,9,6,3,1]
    tree1 := &TreeNode {
        4,
        &TreeNode { 2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode { 7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}, },
    }
    fmt.Println(mirrorTree(tree1))
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/invert2-tree.jpg" />
    // Input: root = [2,1,3]
    // Output: [2,3,1]
    tree2 := &TreeNode {
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(mirrorTree(tree2))

    tree11 := &TreeNode{
        4,
        &TreeNode { 2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode { 7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}, },
    }
    fmt.Println(mirrorTree1(tree11))
    tree12 := &TreeNode{
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(mirrorTree1(tree12))
}