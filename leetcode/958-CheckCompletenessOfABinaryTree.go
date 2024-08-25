package main

// 958. Check Completeness of a Binary Tree
// Given the root of a binary tree, determine if it is a complete binary tree.

// In a complete binary tree, every level, except possibly the last, is completely filled, and all nodes in the last level are as far left as possible. 
// It can have between 1 and 2h nodes inclusive at the last level h.

// Example 1:
//                 1
//               /   \
//              2     3
//            /   \   /
//           4     5 6
// <img src="https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-1.png" />
// Input: root = [1,2,3,4,5,6]
// Output: true
// Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.

// Example 2:
//             1
//           /   \
//          2     3
//         /  \    \
//        4    5    7
// <img src="https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-2.png" />
// Input: root = [1,2,3,4,5,null,7]
// Output: false
// Explanation: The node with value 7 isn't as far left as possible.

// Constraints:
//     The number of nodes in the tree is in the range [1, 100].
//     1 <= Node.val <= 1000

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
// Horizontal Traverse bfs
func isCompleteTree(root *TreeNode) bool {
    queue, flag := []*TreeNode{ root }, false
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            t := queue[0]
            queue = queue[1:]
            if flag && t != nil {
                return false
            }
            if t == nil {
                flag = true
                continue
            }
            queue = append(queue, t.Left)
            queue = append(queue, t.Right)
        }
    }
    return true
}

func main() {
    // Example 1:
    //                 1
    //               /   \
    //              2     3
    //            /   \   /
    //           4     5 6
    // <img src="https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-1.png" />
    // Input: root = [1,2,3,4,5,6]
    // Output: true
    // Explanation: Every level before the last is full (ie. levels with node-values {1} and {2, 3}), and all nodes in the last level ({4, 5, 6}) are as far left as possible.
    tree1 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode{4, nil, nil},  &TreeNode{5, nil, nil}, },
        &TreeNode { 3, &TreeNode{6, nil, nil},  nil, },
    }
    fmt.Println(isCompleteTree(tree1)) // true
    // Example 2:
    //             1
    //           /   \
    //          2     3
    //         /  \    \
    //        4    5    7
    // <img src="https://assets.leetcode.com/uploads/2018/12/15/complete-binary-tree-2.png" />
    // Input: root = [1,2,3,4,5,null,7]
    // Output: false
    // Explanation: The node with value 7 isn't as far left as possible.
    tree2 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode{4, nil, nil},  &TreeNode{5, nil, nil}, },
        &TreeNode { 3, nil,                     &TreeNode{7, nil, nil}, },
    }
    fmt.Println(isCompleteTree(tree2)) // false
}