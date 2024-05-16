package main

// 257. Binary Tree Paths
// Given the root of a binary tree, return all root-to-leaf paths in any order.
// A leaf is a node with no children.

// Example 1:
//         1
//        /  \
//       2    3
//        \
//         5 
// <img src="https://assets.leetcode.com/uploads/2021/03/12/paths-tree.jpg">
// Input: root = [1,2,3,null,5]
// Output: ["1->2->5","1->3"]

// Example 2:
// Input: root = [1]
// Output: ["1"]

// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
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
func binaryTreePaths(root *TreeNode) []string {
    var binaryTreePathsWithString func (node *TreeNode, str string) []string
    binaryTreePathsWithString = func (node *TreeNode, str string) []string {
        if node == nil {
            return []string{}
        }
        if node.Left == nil && node.Right == nil {
            str += fmt.Sprint(node.Val)
            return []string{str}
        }
        str += fmt.Sprintf("%v->", node.Val)
        return append(
            binaryTreePathsWithString(node.Left, str),
            binaryTreePathsWithString(node.Right, str)...,
        )
    }
    return binaryTreePathsWithString(root, "")
}

func main() {
    // Example 1:
    //         1
    //        /  \
    //       2    3
    //        \
    //         5 
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/paths-tree.jpg">
    // Input: root = [1,2,3,null,5]
    // Output: ["1->2->5","1->3"]
    tree1 := &TreeNode {
        1,
        &TreeNode {
            2,
            nil,
            &TreeNode{5, nil, nil},
        },
        &TreeNode{3, nil, nil},
    }
    fmt.Println(binaryTreePaths(tree1)) // ["1->2->5","1->3"]
    // Example 2:
    // Input: root = [1]
    // Output: ["1"]
    tree2 := &TreeNode{1, nil, nil}
    fmt.Println(binaryTreePaths(tree2)) // ["1"]
}