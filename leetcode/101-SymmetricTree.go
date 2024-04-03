package main

// 101. Symmetric Tree
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/symtree1.jpg" />
// Input: root = [1,2,2,3,4,4,3]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/19/symtree2.jpg" />
// Input: root = [1,2,2,null,3,null,3]
// Output: false
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     -100 <= Node.val <= 100
    
// Follow up: Could you solve it both recursively and iteratively?

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
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
func isSymmetric(root *TreeNode) bool {
    var helper func(left *TreeNode, right *TreeNode) bool 
    helper = func(left *TreeNode, right *TreeNode) bool {
        if left == nil && right == nil {
            return true
        }
        if left == nil || right == nil {
            return false
        }
        // 左右需要交替才是对称
        return left.Val == right.Val && helper(left.Right, right.Left) && helper(left.Left, right.Right)
    }
    return helper(root, root)
}

func main() {
	fmt.Println(isSymmetric(
        &TreeNode {
            1,
            &TreeNode{
                2, 
                &TreeNode{3, nil, nil},
                &TreeNode{4, nil, nil},
            },
            &TreeNode{
                2, 
                &TreeNode{4, nil, nil},
                &TreeNode{3, nil, nil},
            },
        },
    )) // true
    fmt.Println(isSymmetric(
        &TreeNode {
            1,
            &TreeNode{
                2, 
                nil,
                &TreeNode{4, nil, nil},
            },
            &TreeNode{
                2, 
                nil,
                &TreeNode{3, nil, nil},
            },
        },
    )) // true
}