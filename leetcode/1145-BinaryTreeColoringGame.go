package main

// 1145. Binary Tree Coloring Game
// Two players play a turn based game on a binary tree. 
// We are given the root of this binary tree, and the number of nodes n in the tree. 
// n is odd, and each node has a distinct value from 1 to n.

// Initially, the first player names a value x with 1 <= x <= n, 
// and the second player names a value y with 1 <= y <= n and y != x. 
// The first player colors the node with value x red, and the second player colors the node with value y blue.

// Then, the players take turns starting with the first player. 
// In each turn, that player chooses a node of their color (red if player 1, blue if player 2) 
// and colors an uncolored neighbor of the chosen node (either the left child, right child, or parent of the chosen node.)

// If (and only if) a player cannot choose such a node in this way, they must pass their turn. 
// If both players pass their turn, the game ends, and the winner is the player that colored more nodes.

// You are the second player. 
// If it is possible to choose such a y to ensure you win the game, return true. 
// If it is not possible, return false.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/08/01/1480-binary-tree-coloring-game.png" />
// Input: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
// Output: true
// Explanation: The second player can choose the node with value 2.

// Example 2:
// Input: root = [1,2,3], n = 3, x = 1
// Output: false

// Constraints:
//     The number of nodes in the tree is n.
//     1 <= x <= n <= 100
//     n is odd.
//     1 <= Node.val <= n
//     All the values of the tree are unique.

import "fmt"

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
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
    left, right := 0, 0
    var dfs func(node *TreeNode, isLeft, isRight bool, x int)
    dfs = func(node *TreeNode, isLeft, isRight bool, x int) {
        if node == nil { return }
        if node.Val == x {
            dfs(node.Left,  true,  false, x)
            dfs(node.Right, false, true, x)
            return
        }
        if isLeft { left++ }
        if isRight { right++ }
        dfs(node.Left,  isLeft, isRight, x)
        dfs(node.Right, isLeft, isRight, x)
    }
    dfs(root, false, false, x)
    if left > n / 2 || right > n / 2 { return true } // if any subtree nodes more than half, WIN
    if right + left + 1 <= n / 2 { return true } // if x root tree nodes less than or equal to n / 2, WIN
    return false
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/08/01/1480-binary-tree-coloring-game.png" />
    // Input: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
    // Output: true
    // Explanation: The second player can choose the node with value 2.
    tree1 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{4, &TreeNode{8, nil, nil, }, &TreeNode{9, nil, nil, }, }, &TreeNode{5, &TreeNode{10, nil, nil, }, &TreeNode{11, nil, nil, }, }, },
        &TreeNode{3, &TreeNode{6, nil, nil, }, &TreeNode{7, nil, nil, }, },
    }
    fmt.Println(btreeGameWinningMove(tree1, 11, 3)) // true
    // Example 2:
    // Input: root = [1,2,3], n = 3, x = 1
    // Output: false
    tree2 := &TreeNode {
        1,
        &TreeNode{2, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(btreeGameWinningMove(tree2, 3, 1)) // false
}