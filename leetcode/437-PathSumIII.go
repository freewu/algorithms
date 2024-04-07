package main

// 437. Path Sum III
// Given the root of a binary tree and an integer targetSum, 
// return the number of paths where the sum of the values along the path equals targetSum.

// The path does not need to start or end at the root or a leaf,
// but it must go downwards (i.e., traveling only from parent nodes to child nodes).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/04/09/pathsum3-1-tree.jpg" / >
// Input: root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
// Output: 3
// Explanation: The paths that sum to 8 are shown.

// Example 2:
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: 3
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 1000].
//     -10^9 <= Node.val <= 10^9
//     -1000 <= targetSum <= 1000

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
func pathSum(root *TreeNode, targetSum int) int {
    res := 0
    var dfs func(tree *TreeNode, targetSum int, pathStack []*TreeNode)
    dfs = func (tree *TreeNode, targetSum int, pathStack []*TreeNode) {
        if tree == nil {
            return
        }
        sum := 0
        pathStack = append(pathStack, tree)
        for i := len(pathStack) - 1; i >= 0; i-- {
            sum += pathStack[i].Val
            if sum == targetSum {
                res++
            }
        }
        dfs(tree.Left, targetSum, pathStack)
        dfs(tree.Right, targetSum, pathStack)
    }
    dfs(root, targetSum, []*TreeNode{})
    return res
}

func main() {
    tree1 := &TreeNode {
        10,
        &TreeNode { 
            5, 
            &TreeNode {
                3,
                &TreeNode{3, nil, nil},
                &TreeNode{-2, nil, nil},
            },
            &TreeNode {
                2,
                nil,
                &TreeNode{1, nil, nil},
            },
        },
        &TreeNode {
            -3,
            nil,
            &TreeNode{11, nil, nil},
        },
    }
    tree2 := &TreeNode {
        5,
        &TreeNode { 
            4, 
            &TreeNode {
                11,
                &TreeNode{7, nil, nil},
                &TreeNode{2, nil, nil},
            },
            nil,
        },
        &TreeNode {
            8,
            &TreeNode{13, nil, nil},
            &TreeNode{
                4,
                &TreeNode{5, nil, nil},
                &TreeNode{1, nil, nil},
            },
        },
    }
    fmt.Println(pathSum(tree1,8)) // 3
    fmt.Println(pathSum(tree2,22)) // 3
}