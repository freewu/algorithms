package main

// 222. Count Complete Tree Nodes
// Given the root of a complete binary tree, return the number of the nodes in the tree.

// According to Wikipedia, every level, except possibly the last, 
// is completely filled in a complete binary tree, 
// and all nodes in the last level are as far left as possible. 
// It can have between 1 and 2h nodes inclusive at the last level h.

// Design an algorithm that runs in less than O(n) time complexity.

// Example 1:
//         1
//       /   \
//      2     3
//     /  \  /  
//    4    5 6   
// <img src="https://assets.leetcode.com/uploads/2021/01/14/complete.jpg" />
// Input: root = [1,2,3,4,5,6]
// Output: 6

// Example 2:
// Input: root = []
// Output: 0

// Example 3:
// Input: root = [1]
// Output: 1
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 5 * 10^4].
//     0 <= Node.val <= 5 * 10^4
//     The tree is guaranteed to be complete.

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
func countNodes(root *TreeNode) int {
    res := 1 
    if root == nil {
        return 0
    }
    var dfs func(root *TreeNode) 
    dfs = func(root *TreeNode) {
        if root == nil {
            return 
        }
        if root.Left != nil {
            res++
            dfs(root.Left)
        }
        if root.Right != nil {
            res++
            dfs(root.Right)
        }
    }
    dfs(root)
    return res 
}

func countNodes1(root *TreeNode) int {
    // Define var to track count
    res := 0
    var dfs func (root *TreeNode, currSum *int)
    dfs = func (root *TreeNode, currSum *int) {
        // Cut condition. If node is nil, return
        if root == nil {
            return
        }
        
        *currSum++ // As node is not nil, count it
        if root.Left != nil { // If there's left subtree to check, go and count nodes there
            dfs(root.Left, currSum) 
        }
        if root.Right != nil { // If there's right subtree to check, go and count nodes there
            dfs(root.Right, currSum) 
        }
    }
    // Pass it as reference
    dfs(root, &res)
    return res
}

func main() {
    // Example 1:
    //         1
    //       /   \
    //      2     3
    //     /  \  /  
    //    4    5 6   
    // <img src="https://assets.leetcode.com/uploads/2021/01/14/complete.jpg" />
    // Input: root = [1,2,3,4,5,6]
    // Output: 6
    tree1 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode { 4, nil, nil}, &TreeNode { 5, nil, nil}, },
        &TreeNode { 3, &TreeNode { 6, nil, nil}, nil},
    }
    fmt.Println(countNodes(tree1)) // 6
    // Example 2:
    // Input: root = []
    // Output: 0
    fmt.Println(countNodes(nil)) // 0
    // Example 3:
    // Input: root = [1]
    // Output: 1
    fmt.Println(countNodes(&TreeNode { 1, nil, nil })) // 1

    fmt.Println(countNodes1(tree1)) // 6
    fmt.Println(countNodes1(nil)) // 0
    fmt.Println(countNodes1(&TreeNode { 1, nil, nil })) // 1
}
