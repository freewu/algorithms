package main 

// 95. Unique Binary Search Trees II
// Given an integer n, return all the structurally unique BST's (binary search trees), 
// which has exactly n nodes of unique values from 1 to n. Return the answer in any order.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/18/uniquebstn3.jpg">
// Input: n = 3
// Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]

// Example 2:
// Input: n = 1
// Output: [[1]]

// Constraints:
//     1 <= n <= 8

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
func generateTrees(n int) []*TreeNode {
    if n == 0 {
        return []*TreeNode{}
    }
    var generate func(start, end int) []*TreeNode 
    generate = func(start, end int) []*TreeNode {
        if start > end {
            return []*TreeNode{nil}
        }
        res := []*TreeNode{}
        for i := start; i <= end; i++ {
            leftTrees := generate(start, i - 1)
            rightTrees := generate(i + 1, end)
            for _, l := range leftTrees {
                for _, r := range rightTrees {
                    currentTree := &TreeNode{Val: i, Left: l, Right: r}
                    res = append(res, currentTree)
                }
            }
        }
        return res
    }
    return generate(1, n)
}

func main() {
    fmt.Println(generateTrees(1))
    fmt.Println(generateTrees(2))
    fmt.Println(generateTrees(3))
    fmt.Println(generateTrees(4))
    fmt.Println(generateTrees(5))
}