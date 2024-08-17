package main

// LCR 193. 二叉搜索树的最近公共祖先
// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。
// 百度百科中最近公共祖先的定义为：“对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”
// 例如，给定如下二叉搜索树:  root = [6,2,8,0,4,7,9,null,null,3,5]
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/12/14/binarysearchtree_improved.png" />

// 示例 1:
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
// 输出: 6 
// 解释: 节点 2 和节点 8 的最近公共祖先是 6。

// 示例 2:
// 输入: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
// 输出: 2
// 解释: 节点 2 和节点 4 的最近公共祖先是 2, 因为根据定义最近公共祖先节点可以为节点本身。

// 说明:
//     所有节点的值都是唯一的。
//     p、q 为不同节点且均存在于给定的二叉搜索树中。

// 公共祖先
// 对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

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
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if p == nil || q == nil || root == nil {
        return nil
    }
    // 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）
    if p.Val < root.Val && q.Val < root.Val {
        return lowestCommonAncestor(root.Left, p, q)
    }
    if p.Val > root.Val && q.Val > root.Val {
        return lowestCommonAncestor(root.Right, p, q)
    }
    return root
}

// 迭代
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    var curr *TreeNode = root
    for curr != nil {
        if p.Val < curr.Val && q.Val < curr.Val {
            curr = curr.Left
            continue
        }
        if p.Val > curr.Val && q.Val > curr.Val {
            curr = curr.Right
            continue
        }
        return curr
    }
    return curr
}

// best solution
func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	for root != nil {
        if root.Val < p.Val && root.Val < q.Val {
            root = root.Right
        } else if root.Val > p.Val && root.Val > q.Val {
            root = root.Left
        } else {
            return root
        }
    }
    return nil
}

func main() {
    tree1 := &TreeNode {
        6,
        &TreeNode { 2, &TreeNode{0, nil, nil}, &TreeNode{ 4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}, }, },
        &TreeNode{ 8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}, },
    }
    tree3 := &TreeNode {
        2,
        &TreeNode { 1, nil, nil, },
        nil,
    }
    fmt.Println(lowestCommonAncestor(tree1, &TreeNode{2, nil, nil}, &TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor(tree1, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor(tree3, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil})) // 2

    fmt.Println(lowestCommonAncestor1(tree1, &TreeNode{2, nil, nil}, &TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor1(tree1, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor1(tree3, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil})) // 2

    fmt.Println(lowestCommonAncestor2(tree1, &TreeNode{2, nil, nil}, &TreeNode{8, nil, nil})) // 6
    fmt.Println(lowestCommonAncestor2(tree1, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil})) // 2
    fmt.Println(lowestCommonAncestor2(tree3, &TreeNode{2, nil, nil}, &TreeNode{1, nil, nil})) // 2
}