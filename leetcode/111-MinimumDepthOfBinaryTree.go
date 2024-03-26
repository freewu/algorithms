package main

// 111. Minimum Depth of Binary Tree
// Given a binary tree, find its minimum depth.
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
// Note: A leaf is a node with no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/10/12/ex_depth.jpg"/>
// Input: root = [3,9,20,null,null,15,7]
// Output: 2

// Example 2:
// Input: root = [2,null,3,null,4,null,5,null,6]
// Output: 5
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^5].
//     -1000 <= Node.val <= 1000

import "fmt"
import "math"

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
// 递归
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // 叶子节点返回 1
    if root.Left == nil && root.Right == nil {
        return 1
    }
    if root.Left == nil {
        return minDepth(root.Right) + 1
    }
    if root.Right == nil {
        return minDepth(root.Left) + 1
    }
    min := func (a int, b int) int { if a < b { return a; }; return b; }
    return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

// dfs
func minDepth1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    min := func (a int, b int) int { if a < b { return a; }; return b; }
    var dfs  func(node *TreeNode, counter, res int) int
    dfs = func(node *TreeNode, counter, res int) int {
        counter += 1
        // is leaf a leaf is a node with no children.
        if node.Left == nil && node.Right == nil {
            return min(res, counter)
        }
        if l := math.MaxInt; node.Left != nil {
            l = dfs(node.Left, counter, res)
            res = min(res, l)
        }
        if r := math.MaxInt; node.Right != nil {
            r = dfs(node.Right, counter, res)
            res = min(res, r)
        }
        return res
    }
    return dfs(root, 0, math.MaxInt)
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
    fmt.Println(minDepth(tree1)) // 2

    // [2,null,3,null,4,null,5,null,6]
    tree2 := &TreeNode {
        2,
        nil,
        &TreeNode{
            3, 
            nil, 
            &TreeNode{
                4, 
                nil, 
                &TreeNode{
                    5, 
                    nil, 
                    &TreeNode{
                        6, 
                        nil, 
                        nil,
                    },
                },
            },
        },
    }
    fmt.Println(minDepth(tree2)) // 5

    fmt.Println(minDepth1(tree1)) // 2
    fmt.Println(minDepth1(tree2)) // 5
}