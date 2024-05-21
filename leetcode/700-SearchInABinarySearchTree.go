package main

// 700. Search in a Binary Search Tree
// You are given the root of a binary search tree (BST) and an integer val.
// Find the node in the BST that the node's value equals val and return the subtree rooted with that node. 
// If such a node does not exist, return null.

// Example 1:
//        4
//      /   \
//    (2)     7
//   /   \
// (1)   (3)
// <img src="https://assets.leetcode.com/uploads/2021/01/12/tree1.jpg" />
// Input: root = [4,2,7,1,3], val = 2
// Output: [2,1,3]

// Example 2:
//        4
//      /   \
//     2     7
//   /   \
//  1     3
// <img src="https://assets.leetcode.com/uploads/2021/01/12/tree2.jpg" />
// Input: root = [4,2,7,1,3], val = 5
// Output: []

// Constraints:
//     The number of nodes in the tree is in the range [1, 5000].
//     1 <= Node.val <= 10^7
//     root is a binary search tree.
//     1 <= val <= 10^7

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
func searchBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    }
    if root.Val > val { // 小的在左边了, 大了就去左边找
        return searchBST(root.Left, val)
    }
    return searchBST(root.Right, val) // 小了就去右边找
}

func main() {
    // Example 1:
    //        4
    //      /   \
    //    (2)     7
    //   /   \
    // (1)   (3)
    // <img src="https://assets.leetcode.com/uploads/2021/01/12/tree1.jpg" />
    // Input: root = [4,2,7,1,3], val = 2
    // Output: [2,1,3]
    tree1 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}, },
        &TreeNode{7, nil, nil},
    }
    t1 := searchBST(tree1,2)
    fmt.Println("t1: ", t1) // t1:  &{2 0xc000008078 0xc000008090}
    // Example 2:
    //        4
    //      /   \
    //     2     7
    //   /   \
    //  1     3
    // <img src="https://assets.leetcode.com/uploads/2021/01/12/tree2.jpg" />
    // Input: root = [4,2,7,1,3], val = 5
    // Output: []
    t2 := searchBST(tree1,5)
    fmt.Println("t2: ", t2) // t2:  <nil>
}