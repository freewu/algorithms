package main

// 1372. Longest ZigZag Path in a Binary Tree
// You are given the root of a binary tree.
// A ZigZag path for a binary tree is defined as follow:
//     Choose any node in the binary tree and a direction (right or left).
//     If the current direction is right, move to the right child of the current node; otherwise, move to the left child.
//     Change the direction from right to left or from left to right.
//     Repeat the second and third steps until you can't move in the tree.

// Zigzag length is defined as the number of nodes visited - 1. (A single node has a length of 0).
// Return the longest ZigZag path contained in that tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_1_1702.png" />
// Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
// Output: 3
// Explanation: Longest ZigZag path in blue nodes (right -> left -> right).

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_2_1702.png" />
// Input: root = [1,1,1,null,1,null,null,1,1,null,1]
// Output: 4
// Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).

// Example 3:
// Input: root = [1]
// Output: 0
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 5 * 10^4].
//     1 <= Node.val <= 100


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
func longestZigZag(root *TreeNode) int {
    res := 0
    var dfs func(n *TreeNode, max *int, count int, parentRight bool)
    dfs = func(n *TreeNode, max *int, count int, parentRight bool) {
        if n == nil { return; }
        if count > *max {
            *max = count
        }
        if parentRight {
            dfs(n.Left, max, count+1, false)
            dfs(n.Right, max, 1, true)
        } else {
            dfs(n.Left, max, 1, false)
            dfs(n.Right, max, count+1, true)
        }
    }
    dfs(root, &res, 0 , false)
    return res
}

func longestZigZag1(root *TreeNode) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node *TreeNode, lastIsLeft bool, len int)
    dfs = func(node *TreeNode, lastIsLeft bool, len int) {
        if node == nil { return; }
        res = max(res, len)
        if lastIsLeft {
            dfs(node.Left, true, 1) // 方向不变，从头开始计算
            dfs(node.Right, false, len+1) // 方向改变，长度加1
            return
        }
        dfs(node.Left, true, len+1) // 方向改变，长度加1
        dfs(node.Right, false, 1) // 方向不变，从头开始计算
    }
    dfs(root, false, 0)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_1_1702.png" />
    // Input: root = [1,null,1,1,1,null,null,1,1,null,1,null,null,null,1]
    // Output: 3
    // Explanation: Longest ZigZag path in blue nodes (right -> left -> right).
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode { 1, &TreeNode { 1, nil, nil }, &TreeNode { 1, &TreeNode { 1, nil, &TreeNode { 1, nil, &TreeNode { 1, nil, nil }, }, }, &TreeNode { 1, nil, nil }, } },
    }
    fmt.Println(longestZigZag(tree1)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/01/22/sample_2_1702.png" />
    // Input: root = [1,1,1,null,1,null,null,1,1,null,1]
    // Output: 4
    // Explanation: Longest ZigZag path in blue nodes (left -> right -> left -> right).
    tree2 := &TreeNode { 
        1, 
        &TreeNode { 1, nil,  &TreeNode { 1, &TreeNode { 1, nil, &TreeNode { 1, nil, nil, }, }, &TreeNode { 1, nil, nil, } }, }, 
        &TreeNode { 1, nil, nil, },
    }
    fmt.Println(longestZigZag(tree2)) // 4
    // Example 3:
    // Input: root = [1]
    // Output: 0
    tree3 := &TreeNode { 1, nil, nil }
    fmt.Println(longestZigZag(tree3)) // 0

    fmt.Println(longestZigZag1(tree1)) // 3
    fmt.Println(longestZigZag1(tree2)) // 4
    fmt.Println(longestZigZag1(tree3)) // 0
}