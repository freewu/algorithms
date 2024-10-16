package main

// 1740. Find Distance in a Binary Tree
// Given the root of a binary tree and two integers p and q, 
// return the distance between the nodes of value p and value q in the tree.

// The distance between two nodes is the number of edges on the path from one to the other.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"/>
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 0
// Output: 3
// Explanation: There are 3 edges between 5 and 0: 5-3-1-0.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"/>
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 7
// Output: 2
// Explanation: There are 2 edges between 5 and 7: 5-2-7.

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"/>
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 5
// Output: 0
// Explanation: The distance between a node and itself is 0.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     0 <= Node.val <= 10^9
//     All Node.val are unique.
//     p and q are values in the tree.

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
func findDistance(root *TreeNode, p int, q int) int {
    var dfs func(root *TreeNode,level int) int
    dfs = func(root *TreeNode,level int) int {
        if root == nil { return -1 }
        if root.Val == p && root.Val == q {  return 0 }
        l, r := dfs(root.Left,level + 1), dfs(root.Right,level + 1)
        // 上下的话，下减上
        if root.Val == p || root.Val == q {
            if l != -1 { return l - level }
            if r != -1 { return r - level }
            return level
        }
        if l == -1 { return r }
        if r == -1 { return l }
        // 同一个祖先分别在左右，左+右
        return (l - level) + (r - level)
    }
    return dfs(root, 0)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"/>
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 0
    // Output: 3
    // Explanation: There are 3 edges between 5 and 0: 5-3-1-0.
    tree1 := &TreeNode {
        3,
        &TreeNode { 5, &TreeNode{6, nil, nil}, &TreeNode{ 2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}, }, },
        &TreeNode { 1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}, },
    }
    fmt.Println(findDistance(tree1, 5, 0)) // 3
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"/>
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 7
    // Output: 2
    // Explanation: There are 2 edges between 5 and 7: 5-2-7.
    fmt.Println(findDistance(tree1, 5, 7)) // 2
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png"/>
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 5
    // Output: 0
    // Explanation: The distance between a node and itself is 0.
    fmt.Println(findDistance(tree1, 5, 5)) // 0
}