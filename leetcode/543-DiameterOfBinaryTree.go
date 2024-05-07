package main

// 543. Diameter of Binary Tree
// Given the root of a binary tree, return the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. 
// This path may or may not pass through the root.
// The length of a path between two nodes is represented by the number of edges between them.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/03/06/diamtree.jpg" />
// Input: root = [1,2,3,4,5]
// Output: 3
// Explanation: 3 is the length of the path [4,2,1,3] or [5,2,1,3].

// Example 2:
// Input: root = [1,2]
// Output: 1
 
// Constraints:
//         The number of nodes in the tree is in the range [1, 10^4].
//         -100 <= Node.val <= 100

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
func diameterOfBinaryTree(root *TreeNode) int {
	result := 0
    var checkDiameter func(root *TreeNode, result *int) int 
    checkDiameter = func(root *TreeNode, result *int) int {
        if root == nil {
            return 0
        }
        max := func (x, y int) int { if x > y { return x; }; return y; }
        left := checkDiameter(root.Left, result)
        right := checkDiameter(root.Right, result)
        *result = max(*result, left+right)
        return max(left, right) + 1
    }
	checkDiameter(root, &result)
	return result
}

func diameterOfBinaryTree1(root *TreeNode) (ans int) {
    res := 0
    var depth func(*TreeNode) int
    depth = func (root *TreeNode) int {
        if root == nil {
            return 0
        }
        left := depth(root.Left)
        right := depth(root.Right)
        max := func (x, y int) int { if x > y { return x; }; return y; }
        res = max(res, left + right)
        return max(left, right) + 1
    }
    depth(root)
    return res
}

func main() {
    fmt.Println(diameterOfBinaryTree(
        &TreeNode {
            1,
            &TreeNode{
                2, 
                &TreeNode{4, nil, nil},
                &TreeNode{5, nil, nil},
            },
            &TreeNode{3, nil, nil},
        },
    )) // 3
    fmt.Println(diameterOfBinaryTree(
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            nil,
        },
    )) // 1

    fmt.Println(diameterOfBinaryTree1(
        &TreeNode {
            1,
            &TreeNode{
                2, 
                &TreeNode{4, nil, nil},
                &TreeNode{5, nil, nil},
            },
            &TreeNode{3, nil, nil},
        },
    )) // 3
    fmt.Println(diameterOfBinaryTree1(
        &TreeNode {
            1,
            &TreeNode{2, nil, nil},
            nil,
        },
    )) // 1
 }