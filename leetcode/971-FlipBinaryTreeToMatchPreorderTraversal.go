package main

// 971. Flip Binary Tree To Match Preorder Traversal
// You are given the root of a binary tree with n nodes, where each node is uniquely assigned a value from 1 to n. 
// You are also given a sequence of n values voyage, which is the desired pre-order traversal of the binary tree.

// Any node in the binary tree can be flipped by swapping its left and right subtrees. 
// For example, flipping node 1 will have the following effect:
// <img src="https://assets.leetcode.com/uploads/2021/02/15/fliptree.jpg" />
//         1                   1
//       /   \                /  \
//      2     3       =>     3    2
//           /  \          /   \
//          4    5        4     5

// Flip the smallest number of nodes so that the pre-order traversal of the tree matches voyage.

// Return a list of the values of all flipped nodes. You may return the answer in any order. 
// If it is impossible to flip the nodes in the tree to make the pre-order traversal match voyage, return the list [-1].

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-01.png" />
//      1
//     /
//    2
// Input: root = [1,2], voyage = [2,1]
// Output: [-1]
// Explanation: It is impossible to flip the nodes such that the pre-order traversal matches voyage.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
//      1
//     /  \
//    2    3
// Input: root = [1,2,3], voyage = [1,3,2]
// Output: [1]
// Explanation: Flipping node 1 swaps nodes 2 and 3, so the pre-order traversal matches voyage.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
//      1
//     /  \
//    2    3
// Input: root = [1,2,3], voyage = [1,2,3]
// Output: []
// Explanation: The tree's pre-order traversal already matches voyage, so no nodes need to be flipped.

// Constraints:
//     The number of nodes in the tree is n.
//     n == voyage.length
//     1 <= n <= 100
//     1 <= Node.val, voyage[i] <= n
//     All the values in the tree are unique.
//     All the values in voyage are unique.

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
func flipMatchVoyage(root *TreeNode, voyage []int) []int {
    stack := []*TreeNode{ root }
    res, i := []int{}, 0
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if voyage[i] != node.Val { return []int{-1} }
        i++
        if node.Right != nil && node.Left != nil && node.Left.Val != voyage[i] {
            node.Left, node.Right = node.Right, node.Left
            res = append(res, node.Val)
        }
        if node.Right != nil { stack = append(stack, node.Right) }
        if node.Left != nil { stack = append(stack, node.Left) }
    }
    return res
}

// dfs
func flipMatchVoyage1(root *TreeNode, voyage []int) []int {
    res, i := []int{}, 0
    var dfs func(*TreeNode) bool
    dfs = func(root *TreeNode) bool {
        if root == nil {
            return true
        } else if root.Val != voyage[i] {
            return false
        }
        i++
        if root.Left != nil && root.Left.Val != voyage[i] {
            res = append(res, root.Val)
            return dfs(root.Right) && dfs(root.Left)
        }
        return dfs(root.Left) && dfs(root.Right)
    }
    if dfs(root) {
        return res
    }
    return []int{-1}
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-01.png" />
    //      1
    //     /
    //    2
    // Input: root = [1,2], voyage = [2,1]
    // Output: [-1]
    // Explanation: It is impossible to flip the nodes such that the pre-order traversal matches voyage.
    tree1 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        nil,
    }
    fmt.Println(flipMatchVoyage(tree1,[]int{2,1})) // [-1]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
    //      1
    //     /  \
    //    2    3
    // Input: root = [1,2,3], voyage = [1,3,2]
    // Output: [1]
    // Explanation: Flipping node 1 swaps nodes 2 and 3, so the pre-order traversal matches voyage.
    tree2 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(flipMatchVoyage(tree2,[]int{1,3,2})) // [1]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/01/02/1219-02.png" />
    //      1
    //     /  \
    //    2    3
    // Input: root = [1,2,3], voyage = [1,2,3]
    // Output: []
    // Explanation: The tree's pre-order traversal already matches voyage, so no nodes need to be flipped.
    tree3 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(flipMatchVoyage(tree3,[]int{1,2,3})) // []

    tree11 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        nil,
    }
    fmt.Println(flipMatchVoyage1(tree11,[]int{2,1})) // [-1]
    tree12 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(flipMatchVoyage1(tree12,[]int{1,3,2})) // [1]
    tree13 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(flipMatchVoyage1(tree13,[]int{1,2,3})) // []
}