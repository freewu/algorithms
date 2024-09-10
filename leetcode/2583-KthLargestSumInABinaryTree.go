package main

// 2583. Kth Largest Sum in a Binary Tree
// You are given the root of a binary tree and a positive integer k.
// The level sum in the tree is the sum of the values of the nodes that are on the same level.
// Return the kth largest level sum in the tree (not necessarily distinct). 
// If there are fewer than k levels in the tree, return -1.
// Note that two nodes are on the same level if they have the same distance from the root.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/12/14/binaryytreeedrawio-2.png"/>
// Input: root = [5,8,9,2,1,3,7,4,6], k = 2
// Output: 13
// Explanation: The level sums are the following:
// - Level 1: 5.
// - Level 2: 8 + 9 = 17.
// - Level 3: 2 + 1 + 3 + 7 = 13.
// - Level 4: 4 + 6 = 10.
// The 2nd largest level sum is 13.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/12/14/treedrawio-3.png"/>
// Input: root = [1,2,null,3], k = 1
// Output: 3
// Explanation: The largest level sum is 3.

// Constraints:
//         The number of nodes in the tree is n.
//         2 <= n <= 10^5
//         1 <= Node.val <= 10^6
//         1 <= k <= n

import "fmt"
import "sort"

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
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    levelSums := make([]int64, 0)
    nextLevel := []*TreeNode{root}
    for {
        levelSum := 0
        currentLevel := nextLevel
        nextLevel = make([]*TreeNode, 0)
        for _, v := range currentLevel {
            // 计算每层的和
            levelSum += v.Val
            if v.Right != nil {
                nextLevel = append(nextLevel, v.Right)
            }
            if v.Left != nil {
                nextLevel = append(nextLevel, v.Left)
            }
        }
        levelSums = append(levelSums, int64(levelSum))
        if len(nextLevel) == 0 {
            break
        }
    }
    // 取的层数不足返回 -1 
    if k > len(levelSums) {
        return -1
    }
    sort.Slice(levelSums,func(i, j int) bool {
        return levelSums[i] > levelSums[j]
    })
    // 取 top k
    return levelSums[k-1]
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/12/14/binaryytreeedrawio-2.png"/>
    // Input: root = [5,8,9,2,1,3,7,4,6], k = 2
    // Output: 13
    // Explanation: The level sums are the following:
    // - Level 1: 5.
    // - Level 2: 8 + 9 = 17.
    // - Level 3: 2 + 1 + 3 + 7 = 13.
    // - Level 4: 4 + 6 = 10.
    // The 2nd largest level sum is 13.
    tree1 := &TreeNode {
        5,
        &TreeNode { 8 , &TreeNode { 2 , &TreeNode { 4 ,nil, nil }, &TreeNode { 6 ,nil, nil }, },  &TreeNode { 1 ,nil, nil }, },
        &TreeNode { 9 , &TreeNode { 3 ,nil, nil },  &TreeNode { 7 ,nil, nil }, },
    }
    fmt.Println(kthLargestLevelSum(tree1, 2)) // 13
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/12/14/treedrawio-3.png"/>
    // Input: root = [1,2,null,3], k = 1
    // Output: 3
    // Explanation: The largest level sum is 3.
    tree2 := &TreeNode {
        1,
        &TreeNode { 2 , &TreeNode { 3 , nil , nil, }, nil, },
        nil,
    }
    fmt.Println(kthLargestLevelSum(tree2, 1)) // 3
}