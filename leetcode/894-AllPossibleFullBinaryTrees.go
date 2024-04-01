package main

// 894. All Possible Full Binary Trees
// Given an integer n, return a list of all possible full binary trees with n nodes. 
// Each node of each tree in the answer must have Node.val == 0.

// Each element of the answer is the root node of one possible tree. 
// You may return the final list of trees in any order.

// A full binary tree is a binary tree where each node has exactly 0 or 2 children.

// Example 1:
// <img src="https://s3-lc-upload.s3.amazonaws.com/uploads/2018/08/22/fivetrees.png" />
// Input: n = 7 
// Output: [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]

// Example 2:
// Input: n = 3
// Output: [[0,0,0]]
 
// Constraints:
//     1 <= n <= 20

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
func allPossibleFBT(n int) []*TreeNode {
    var cache map[int][]*TreeNode = make(map[int][]*TreeNode)
    if n % 2 == 0 {
        return make([]*TreeNode, 0)
    }
    var getAllFBTs func (n int, cache map[int][]*TreeNode) []*TreeNode 
    getAllFBTs = func (n int, cache map[int][]*TreeNode) []*TreeNode {
        if n == 1 {
            return []*TreeNode{&TreeNode{}}
        }
        if val, ok := cache[n]; ok {
            return val
        }
        result := make([]*TreeNode, 0)
        for left := 1 ; left <= n - 2; left += 2 {
            leftNodes, rightNodes := getAllFBTs(left, cache), getAllFBTs(n - 1 - left, cache)
            for _, leftNode := range leftNodes {
                for _, rightNode := range rightNodes {
                    root := &TreeNode{Left: leftNode, Right: rightNode}
                    result = append(result, root)
                }
            }
        }
        cache[n] = result
        return result
    }
    return getAllFBTs(n, cache)
}

func main() {
    fmt.Println(allPossibleFBT(7))
    fmt.Println(allPossibleFBT(3))
}