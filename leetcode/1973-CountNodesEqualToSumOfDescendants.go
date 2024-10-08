package main

// 1973. Count Nodes Equal to Sum of Descendants
// Given the root of a binary tree, 
// return the number of nodes where the value of the node is equal to the sum of the values of its descendants.

// A descendant of a node x is any node that is on the path from node x to some leaf node. 
// The sum is considered to be 0 if the node has no descendants.

// Example 1:
//           (10)
//           /   \
//         (3)     4
//        /   \
//       2     1
// <img src="https://assets.leetcode.com/uploads/2021/08/17/screenshot-2021-08-17-at-17-16-50-diagram-drawio-diagrams-net.png" />
// Input: root = [10,3,4,2,1]
// Output: 2
// Explanation:
// For the node with value 10: The sum of its descendants is 3+4+2+1 = 10.
// For the node with value 3: The sum of its descendants is 2+1 = 3.

// Example 2:
//         2
//       /
//      3
//     /
//    2
// <img src="https://assets.leetcode.com/uploads/2021/08/17/screenshot-2021-08-17-at-17-25-21-diagram-drawio-diagrams-net.png" />
// Input: root = [2,3,null,2,null]
// Output: 0
// Explanation:
// No node has a value that is equal to the sum of its descendants.

// Example 3:
//     (0)
// <img src="https://assets.leetcode.com/uploads/2021/08/17/screenshot-2021-08-17-at-17-23-53-diagram-drawio-diagrams-net.png" />
// Input: root = [0]
// Output: 1
// For the node with value 0: The sum of its descendants is 0 since it has no descendants.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     0 <= Node.val <= 10^5

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
func equalToDescendants(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0  }
        if node.Left == nil && node.Right == nil {
            if node.Val == 0 { res++ }
            return node.Val
        }
        left, right := dfs(node.Left), dfs(node.Right)
        if node.Val == left + right { // 值等于左右子树之和
            res++
        }
        return left + right + node.Val
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    //           (10)
    //           /   \
    //         (3)     4
    //        /   \
    //       2     1
    // <img src="https://assets.leetcode.com/uploads/2021/08/17/screenshot-2021-08-17-at-17-16-50-diagram-drawio-diagrams-net.png" />
    // Input: root = [10,3,4,2,1]
    // Output: 2
    // Explanation:
    // For the node with value 10: The sum of its descendants is 3+4+2+1 = 10.
    // For the node with value 3: The sum of its descendants is 2+1 = 3.
    tree1 := &TreeNode {
        10,
        &TreeNode { 3, &TreeNode { 2, nil, nil, }, &TreeNode { 1, nil, nil, }, },
        &TreeNode { 4, nil, nil, },
    }
    fmt.Println(equalToDescendants(tree1)) // 2
    // Example 2:
    //         2
    //       /
    //      3
    //     /
    //    2
    // <img src="https://assets.leetcode.com/uploads/2021/08/17/screenshot-2021-08-17-at-17-25-21-diagram-drawio-diagrams-net.png" />
    // Input: root = [2,3,null,2,null]
    // Output: 0
    // Explanation:
    // No node has a value that is equal to the sum of its descendants.
    tree2 := &TreeNode {
        2,
        &TreeNode { 3, &TreeNode { 2, nil, nil, }, nil, },
        nil,
    }
    fmt.Println(equalToDescendants(tree2)) // 0
    // Example 3:
    //     (0)
    // <img src="https://assets.leetcode.com/uploads/2021/08/17/screenshot-2021-08-17-at-17-23-53-diagram-drawio-diagrams-net.png" />
    // Input: root = [0]
    // Output: 1
    // For the node with value 0: The sum of its descendants is 0 since it has no descendants.
    tree3 := &TreeNode{ 0, nil, nil, }
    fmt.Println(equalToDescendants(tree3)) // 1
}