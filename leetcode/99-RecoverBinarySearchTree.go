package main

// 99. Recover Binary Search Tree
// You are given the root of a binary search tree (BST), 
// where the values of exactly two nodes of the tree were swapped by mistake. 
// Recover the tree without changing its structure.

// Example 1:
//      (1)        3
//      /         /
//    (3)    =>  1
//      \         \
//       2          2
// <img src="https://assets.leetcode.com/uploads/2020/10/28/recover1.jpg"/>
// Input: root = [1,3,null,null,2]
// Output: [3,1,null,null,2]
// Explanation: 3 cannot be a left child of 1 because 3 > 1. Swapping 1 and 3 makes the BST valid.

// Example 2:
//     (3)           2
//     /  \         /  \
//    1    4   =>  1    4
//        /            /
///     (2)           3
// <img src="https://assets.leetcode.com/uploads/2020/10/28/recover2.jpg"/>
// Input: root = [3,1,4,null,null,2]
// Output: [2,1,4,null,null,3]
// Explanation: 2 cannot be in the right subtree of 3 because 2 < 3. Swapping 2 and 3 makes the BST valid.
 
// Constraints:
//     The number of nodes in the tree is in the range [2, 1000].
//     -2^31 <= Node.val <= 2^31 - 1
 
// Follow up: A solution using O(n) space is pretty straight-forward. Could you devise a constant O(1) space solution?

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
func recoverTree(root *TreeNode)  {
    var prev, first, second *TreeNode
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil { return }
        dfs(node.Left)
        if prev != nil {
            if first == nil && prev.Val >= node.Val { first = prev }
            if first != nil && prev.Val >= node.Val { second = node }
        }
        prev = node
        dfs(node.Right)
    }
    dfs(root)
    first.Val, second.Val = second.Val, first.Val
}

func main() {
    tree1 := &TreeNode {
        1,
        &TreeNode{3, nil, &TreeNode{2, nil, nil}},
        nil,
    }
    fmt.Println("before: ", tree1)
    recoverTree(tree1)
    fmt.Println("after: ", tree1)

    tree2 := &TreeNode {
        3,
        &TreeNode{1, nil, nil},
        &TreeNode{4, &TreeNode{2, nil, nil}, nil},
    }
    fmt.Println("before: ", tree2)
    recoverTree(tree2)
    fmt.Println("after: ", tree2)
}