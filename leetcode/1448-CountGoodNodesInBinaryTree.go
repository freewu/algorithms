package main

// 1448. Count Good Nodes in Binary Tree
// Given a binary tree root, a node X in the tree is named good 
// if in the path from root to X there are no nodes with a value greater than X.
// Return the number of good nodes in the binary tree.

// Example 1:
//         (3)
//        /   \
//       1    (4)
//      /     /  \
//    (3)    1    (5)
// <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png" />
// Input: root = [3,1,4,3,null,1,5]
// Output: 4
// Explanation: Nodes in blue are good.
// Root Node (3) is always a good node.
// Node 4 -> (3,4) is the maximum value in the path starting from the root.
// Node 5 -> (3,4,5) is the maximum value in the path
// Node 3 -> (3,1,3) is the maximum value in the path.

// Example 2:
//        (3)
//        /
//       (3)
//      /  \
//    (4)   2
// <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png" />
// Input: root = [3,3,null,4,2]
// Output: 3
// Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.

// Example 3:
// Input: root = [1]
// Output: 1
// Explanation: Root is considered as good.

// Constraints:
//     The number of nodes in the binary tree is in the range [1, 10^5].
//     Each node's value is between [-10^4, 10^4].

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
func goodNodes(root *TreeNode) int {
    count := 0 // Initialize a variable to count good nodes.
    var dfs func(node *TreeNode, max int) // Define a helper function for the recursive traversal.
    dfs = func(node *TreeNode, max int) { // The dfs function performs a depth-first search traversal of the binary tree.
        if node == nil { return; } // Base case: If the node is nil, return without processing.
        if node.Val >= max {
            max = node.Val   // Update the maximum value encountered so far.
            count++          // Increment the count to indicate a good node.
        }
        // Recursively traverse the left and right subtrees with the updated max value.
        dfs(node.Left, max)
        dfs(node.Right, max)
    }
    dfs(root, root.Val) // Start the traversal with the root node and its value.
    return count // Return the total count of good nodes in the binary tree.
}

func goodNodes1(root *TreeNode) int {
    value := root.Val
    var dfs func(node *TreeNode, maxValue int) int
    dfs = func(node *TreeNode, maxValue int) int {
        if node == nil { return 0; }
        if node.Val >= maxValue {
            return 1 + dfs(node.Left, node.Val) + dfs(node.Right, node.Val)
        } else {
            return dfs(node.Left, maxValue) + dfs(node.Right, maxValue)
        }
    }
    return 1 + dfs(root.Left, value) + dfs(root.Right, value)
}

func main() {
    // Example 1:
    //         (3)
    //        /   \
    //       1    (4)
    //      /     /  \
    //    (3)    1    (5)
    // <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_1.png" />
    // Input: root = [3,1,4,3,null,1,5]
    // Output: 4
    // Explanation: Nodes in blue are good.
    // Root Node (3) is always a good node.
    // Node 4 -> (3,4) is the maximum value in the path starting from the root.
    // Node 5 -> (3,4,5) is the maximum value in the path
    // Node 3 -> (3,1,3) is the maximum value in the path.
    tree1 := &TreeNode {
        3,
        &TreeNode { 1, &TreeNode { 3, nil, nil }, nil },
        &TreeNode { 4, &TreeNode { 1, nil, nil }, &TreeNode { 5, nil, nil } },
    }
    fmt.Println(goodNodes(tree1)) // 4
    // Example 2:
    //         (3)
    //         /
    //       (3)
    //      /  \
    //    (4)   2
    // <img src="https://assets.leetcode.com/uploads/2020/04/02/test_sample_2.png" />
    // Input: root = [3,3,null,4,2]
    // Output: 3
    // Explanation: Node 2 -> (3, 3, 2) is not good, because "3" is higher than it.
    tree2 := &TreeNode {
        3,
        &TreeNode { 3, &TreeNode { 4, nil, nil }, &TreeNode { 2, nil, nil } },
        nil,
    }
    fmt.Println(goodNodes(tree2)) // 3
    // Example 3:
    // Input: root = [1]
    // Output: 1
    // Explanation: Root is considered as good.
    tree3 := &TreeNode { 1, nil, nil }
    fmt.Println(goodNodes(tree3)) // 1

    fmt.Println(goodNodes1(tree1)) // 4
    fmt.Println(goodNodes1(tree2)) // 3
    fmt.Println(goodNodes1(tree3)) // 1
}