package main

// 104. Maximum Depth of Binary Tree
// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/11/26/tmp-tree.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: 3

// Example 2:
// Input: root = [1,null,2]
// Output: 2
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
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
// dfs
func maxDepth3(root *TreeNode) int {
    res := 0
    var dfs func (node *TreeNode, count int)
    dfs = func (node *TreeNode, count int) {
        if node != nil {
            dfs(node.Left, count+1)
            dfs(node.Right, count+1)
        }
        if count > res {
            res = count
        }
    }
    dfs(root, res)
    return res
}

// bfs
func maxDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    nodeList := []*TreeNode{root}
    res := 0
    for len(nodeList) > 0 {
        for _, node := range nodeList {
            nodeList = nodeList[1:]
            if node.Left != nil {
                nodeList = append(nodeList, node.Left)
            }
            if node.Right != nil {
                nodeList = append(nodeList, node.Right)
            }
        }
        res++
    }
    return res
}

// 递归
func maxDepth2(root *TreeNode) int {
    if root == nil {
        return 0
    }
    l := maxDepth2(root.Left) + 1
    r := maxDepth2(root.Right) + 1
    max := func (a int, b int) int { if a > b { return a; }; return b; }
    return max(l,r)
}

// 递归
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    leftDepth := maxDepth(root.Left)
    rightDepth := maxDepth(root.Right)
    if leftDepth < rightDepth {
        return rightDepth + 1
    } else {
        return leftDepth + 1
    }
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode { 9, nil, nil },
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(maxDepth(tree1)) // 3

    tree2 := &TreeNode {
        1,
        nil,
        &TreeNode{2, nil, nil},
    }
    fmt.Println(maxDepth(tree2)) // 2

    fmt.Println(maxDepth1(tree1)) // 3
    fmt.Println(maxDepth1(tree2)) // 2

    fmt.Println(maxDepth2(tree1)) // 3
    fmt.Println(maxDepth2(tree2)) // 2

    fmt.Println(maxDepth3(tree1)) // 3
    fmt.Println(maxDepth3(tree2)) // 2
}