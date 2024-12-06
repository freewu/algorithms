package main

// 3157. Find the Level of Tree with Minimum Sum
// Given the root of a binary tree root where each node has a value, 
// return the level of the tree that has the minimum sum of values among all the levels (in case of a tie, return the lowest level).

// Note that the root of the tree is at level 1 and the level of any other node is its distance from the root + 1.

// Example 1:
// Input: root = [50,6,2,30,80,7]
// Output: 2
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/17/image_2024-05-17_16-15-46.png" />

// Example 2:
// Input: root = [36,17,10,null,null,24]
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/17/image_2024-05-17_16-14-18.png" />

// Example 3:
// Input: root = [5,null,5,null,5]
// Output: 1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/05/19/image_2024-05-19_19-07-20.png" />

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^9

import "fmt"
//import "math"

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
func minimumLevel(root *TreeNode) int {
    queue := []*TreeNode{ root }
    res, sum := 0, 1 << 31
    for level := 1; len(queue) > 0; level++ {
        mn := 0
        for i := len(queue); i > 0; i-- {
            node := queue[0]
            queue = queue[1:]
            mn += node.Val
            if node.Left != nil  { queue = append(queue, node.Left)  }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        if sum > mn {
            sum, res = mn, level
        }
    }
    return res
}

func main() {
    // Example 1:
    // Input: root = [50,6,2,30,80,7]
    // Output: 2
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/17/image_2024-05-17_16-15-46.png" />
    tree1 := &TreeNode{
        50, 
        &TreeNode{ 6, &TreeNode{ 30, nil, nil, }, &TreeNode{ 80, nil, nil, }, },
        &TreeNode{ 2, &TreeNode{  7, nil, nil, }, nil, },
    } 
    fmt.Println(minimumLevel(tree1)) // 2
    // Example 2:
    // Input: root = [36,17,10,null,null,24]
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/17/image_2024-05-17_16-14-18.png" />
    tree2 := &TreeNode{
        36, 
        &TreeNode{ 17, nil, nil, },
        &TreeNode{ 10, &TreeNode{  24, nil, nil, }, nil, },
    } 
    fmt.Println(minimumLevel(tree2)) // 3
    // Example 3:
    // Input: root = [5,null,5,null,5]
    // Output: 1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/05/19/image_2024-05-19_19-07-20.png" />
    tree3 := &TreeNode{
        5, 
        nil,
        &TreeNode{ 5, nil, &TreeNode{  5, nil, nil, },},
    } 
    fmt.Println(minimumLevel(tree3)) // 1
}