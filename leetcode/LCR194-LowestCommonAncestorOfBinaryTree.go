package main

// LCR 194. 二叉树的最近公共祖先
// 给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉树:  root = [3,5,1,6,2,0,8,null,null,7,4]
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/15/binarytree.png" />

// 示例 1:
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// 输出: 3
// 解释: 节点 5 和节点 1 的最近公共祖先是节点 3。

// 示例 2:
// 输入: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// 输出: 5
// 解释: 节点 5 和节点 4 的最近公共祖先是节点 5。因为根据定义最近公共祖先节点可以为节点本身。

// 说明:
//     所有节点的值都是唯一的。
//     p、q 为不同节点且均存在于给定的二叉树中。

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 一个节点也可以是它自己的祖先
    if root == nil || root == q || root == p {
        return root
    }
    // 递归查找 left / right
    left, right := lowestCommonAncestor(root.Left, p, q), lowestCommonAncestor(root.Right, p, q)
    // 如果 left 和 right 都不为 nil 返回 root
    if left != nil && right != nil {
        return root
    }
    if left == nil {
        return right
    }
    return left
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
    // Output: 3
    // Explanation: The LCA of nodes 5 and 1 is 3.
    tree1 := &TreeNode {
        3,
        &TreeNode { 5, &TreeNode{6, nil, nil}, &TreeNode{ 2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}, }, },
        &TreeNode { 1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}, },
    }
    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{5, nil, nil},&TreeNode{1, nil, nil})) // 3
    
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
    // Output: 5
    // Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.
    tree2 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        nil,
    }
    //fmt.Println(lowestCommonAncestor(tree2,&TreeNode{1, nil, nil},&TreeNode{2, nil, nil})) // 1
    fmt.Println(lowestCommonAncestor(tree2,&TreeNode{1, nil, nil},&TreeNode{2, nil, nil})) // 1

    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{5, nil, nil},&TreeNode{4, nil, nil})) // 5
}