package main

// 156. Binary Tree Upside Down
// Given the root of a binary tree, turn the tree upside down and return the new root.
// You can turn a binary tree upside down with the following steps:
//     The original left child becomes the new root.
//     The original root becomes the new right child.
//     The original right child becomes the new left child.
    
// <img src="https://assets.leetcode.com/uploads/2020/08/29/main.jpg" />
// The mentioned steps are done level by level. 
// It is guaranteed that every right node has a sibling (a left node with the same parent) and has no children.

// Example 1:
//         1            1           4
//        /  \         /          /   \
//       2    3  =>   2 -- 3 =>  5     2
//      /  \         /                /  \
//     4    5       4 -- 5           3    1
// <img src="https://assets.leetcode.com/uploads/2020/08/29/updown.jpg" />
// Input: root = [1,2,3,4,5]
// Output: [4,5,2,null,null,3,1]

// Example 2:
// Input: root = []
// Output: []

// Example 3:
// Input: root = [1]
// Output: [1]

// Constraints:
//     The number of nodes in the tree will be in the range [0, 10].
//     1 <= Node.val <= 10
//     Every right node in the tree has a sibling (a left node that shares the same parent).
//     Every right node in the tree has no children.

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
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
    if root == nil || root.Left == nil {
        return root
    }
    var parent, parent_right *TreeNode // 当前节点的父节点，当前节点的右节点
    for root != nil {
        root_left := root.Left // 保存左节点，颠倒为跟节点
        root.Left = parent_right // 当前节点的左节点设为右节点
        parent_right = root.Right // 更新下次迭代的右节点
        root.Right = parent // 原来父节点更新为右节点
        parent = root // 更新下次迭代的 left 的父节点
        root = root_left // 更新左节点为父节点
    }
    return parent
}

func main() {
    // Example 1:
    //         1            1           4
    //        /  \         /          /   \
    //       2    3  =>   2 -- 3 =>  5     2
    //      /  \         /                /  \
    //     4    5       4 -- 5           3    1
    // Input: root = [1,2,3,4,5]
    // Output: [4,5,2,null,null,3,1]
    tree1 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}, },
        &TreeNode{3, nil, nil},
    }
    fmt.Println(upsideDownBinaryTree(tree1).Val) // 4
    // Example 2:
    // Input: root = []
    // Output: []
    fmt.Println(upsideDownBinaryTree(nil)) // nil
    // Example 3:
    // Input: root = [1]
    // Output: [1]
    tree3 := &TreeNode{1, nil, nil}
    fmt.Println(upsideDownBinaryTree(tree3).Val) // 1
}