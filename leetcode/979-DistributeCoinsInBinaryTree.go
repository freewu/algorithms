package main

// 979. Distribute Coins in Binary Tree
// You are given the root of a binary tree with n nodes where each node in the tree has node.val coins. 
// There are n coins in total throughout the whole tree.

// In one move, we may choose two adjacent nodes and move one coin from one node to another. 
// A move may be from parent to child, or from child to parent.

// Return the minimum number of moves required to make every node have exactly one coin.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/01/18/tree1.png" />
// Input: root = [3,0,0]
// Output: 2
// Explanation: From the root of the tree, we move one coin to its left child, and one coin to its right child.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/01/18/tree2.png" />
// Input: root = [0,3,0]
// Output: 3
// Explanation: 
// From the left child of the root, we move two coins to the root [taking two moves]. 
// Then, we move one coin from the root of the tree to the right child.

// Constraints:
//     The number of nodes in the tree is n.
//     1 <= n <= 100
//     0 <= Node.val <= n
//     The sum of all Node.val is n.

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
func distributeCoins(root *TreeNode) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    var dfs func (root *TreeNode) (move int, pop int) 
    dfs = func (root *TreeNode) (move int, pop int) {
        if root == nil {
            return 0,0
        }
        move1, pop1 := dfs(root.Left)
        move2, pop2 := dfs(root.Right)
        return move1 + move2 + abs(root.Val + pop1 + pop2 - 1), root.Val + pop1 + pop2 - 1
    }
    res, _ := dfs(root)
    return res
}

func distributeCoins1(root *TreeNode) int {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    res := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left, right := 0, 0
        if node.Left != nil { left = dfs(node.Left); }
        if node.Right != nil { right = dfs(node.Right); }
        res += abs(left) + abs(right)
        return left + right + node.Val - 1
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/01/18/tree1.png" />
    // Input: root = [3,0,0]
    // Output: 2
    // Explanation: From the root of the tree, we move one coin to its left child, and one coin to its right child.
    tree1 := &TreeNode {
        3,
        &TreeNode{0, nil, nil },
        &TreeNode{0, nil, nil },
    }
    fmt.Println(distributeCoins(tree1)) // 2
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/01/18/tree2.png" />
    // Input: root = [0,3,0]
    // Output: 3
    // Explanation: 
    // From the left child of the root, we move two coins to the root [taking two moves]. 
    // Then, we move one coin from the root of the tree to the right child.
    tree2 := &TreeNode {
        0,
        &TreeNode{3, nil, nil },
        &TreeNode{0, nil, nil },
    }
    fmt.Println(distributeCoins(tree2)) // 3

    fmt.Println(distributeCoins1(tree1)) // 2
    fmt.Println(distributeCoins1(tree2)) // 3
}