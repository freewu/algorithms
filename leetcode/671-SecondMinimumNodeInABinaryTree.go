package main

// 671. Second Minimum Node In a Binary Tree
// Given a non-empty special binary tree consisting of nodes with the non-negative value, 
// where each node in this tree has exactly two or zero sub-node. 
// If the node has two sub-nodes, then this node's value is the smaller value among its two sub-nodes. 
// More formally, the property root.val = min(root.left.val, root.right.val) always holds.

// Given such a binary tree, you need to output the second minimum value in the set made of all the nodes' value in the whole tree.
// If no such second minimum value exists, output -1 instead.

// Example 1:
//         2
//       /   \
//      2     5
//          /   \
//         5     7
// <img src="https://assets.leetcode.com/uploads/2020/10/15/smbt1.jpg" />
// Input: root = [2,2,5,null,null,5,7]
// Output: 5
// Explanation: The smallest value is 2, the second smallest value is 5.

// Example 2:
//       2
//     /   \
//    2     2
// <img src="https://assets.leetcode.com/uploads/2020/10/15/smbt2.jpg" />
// Input: root = [2,2,2]
// Output: -1
// Explanation: The smallest value is 2, but there isn't any second smallest value.

// Constraints:
//     The number of nodes in the tree is in the range [1, 25].
//     1 <= Node.val <= 2^31 - 1
//     root.val == min(root.left.val, root.right.val) for each internal node of the tree.

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
func findSecondMinimumValue(root *TreeNode) int {
    v, inf := root.Val, 1 << 32 -1
    res := inf
    var dfs func(root *TreeNode)
    dfs = func(root *TreeNode) {
        if root == nil {  return }
        if root.Val < v { v = root.Val }
        if root.Val > v && root.Val < res { res = root.Val }
        dfs(root.Left)
        dfs(root.Right)
    }
    dfs(root)
    if res == inf {
        return -1
    }
    return res
}

func main() {
    // Example 1:
    //         2
    //       /   \
    //      2     5
    //          /   \
    //         5     7
    // <img src="https://assets.leetcode.com/uploads/2020/10/15/smbt1.jpg" />
    // Input: root = [2,2,5,null,null,5,7]
    // Output: 5
    // Explanation: The smallest value is 2, the second smallest value is 5.
    tree1 := &TreeNode {
        2,
        &TreeNode { 2, nil,                    nil,                    },
        &TreeNode { 5, &TreeNode{5, nil, nil}, &TreeNode{7, nil, nil}, },
    }
    fmt.Println(findSecondMinimumValue(tree1)) // 5
    // Example 2:
    //       2
    //     /   \
    //    2     2
    // <img src="https://assets.leetcode.com/uploads/2020/10/15/smbt2.jpg" />
    // Input: root = [2,2,2]
    // Output: -1
    // Explanation: The smallest value is 2, but there isn't any second smallest value.
    tree2 := &TreeNode {
        2,
        &TreeNode{2, nil, nil},
        &TreeNode{2, nil, nil},
    }
    fmt.Println(findSecondMinimumValue(tree2)) // - 1
}