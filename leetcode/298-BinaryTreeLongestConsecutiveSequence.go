package main

// 298. Binary Tree Longest Consecutive Sequence
// Given the root of a binary tree, return the length of the longest consecutive sequence path.
// A consecutive sequence path is a path where the values increase by one along the path.
// Note that the path can start at any node in the tree, and you cannot go from a node to its parent in the path.

// Example 1:
//          1
//           \
//            (3)
//           /   \
//          2     (4)
//                  \
//                  (5)
// <img src="https://assets.leetcode.com/uploads/2021/03/14/consec1-1-tree.jpg" />
// Input: root = [1,null,3,2,4,null,null,null,5]
// Output: 3
// Explanation: Longest consecutive sequence path is 3-4-5, so return 3.

// Example 2:
//     (2)
//       \
//        (3)
//        /
//       2
//     / 
//    1
// <img src="https://assets.leetcode.com/uploads/2021/03/14/consec1-2-tree.jpg" />
// Input: root = [2,null,3,2,null,1]
// Output: 2
// Explanation: Longest consecutive sequence path is 2-3, not 3-2-1, so return 2.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 3 * 10^4].
//     -3 * 10^4 <= Node.val <= 3 * 10^4

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
// 前序+自顶向下
func longestConsecutive(root *TreeNode) int {
    var max = func(x, y int) int { if x > y { return x } else { return y } }
    var dfs func(root, parent *TreeNode, l int) int
    dfs = func(root, parent *TreeNode, l int) int {
        if root == nil {
            return 0
        }
        if parent == nil || parent.Val + 1 != root.Val{
            l = 1
        } else {
            l++
        }
        return max(l, max(dfs(root.Left, root, l), dfs(root.Right, root, l)))
    }
    return dfs(root, nil, 0)
}

// 后序+自底向上
func longestConsecutive1(root *TreeNode) int {
    res := 0
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        mx, l, r := 1, dfs(root.Left) + 1, dfs(root.Right) + 1
        if root.Left != nil && root.Val + 1 == root.Left.Val {
            mx = l
        }
        if root.Right != nil && root.Val + 1 == root.Right.Val && mx < r {
            mx = r
        }
        if mx > res { 
            res = mx 
        }
        return mx
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    //          1
    //           \
    //            (3)
    //           /   \
    //          2     (4)
    //                  \
    //                  (5)
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/consec1-1-tree.jpg" />
    // Input: root = [1,null,3,2,4,null,null,null,5]
    // Output: 3
    // Explanation: Longest consecutive sequence path is 3-4-5, so return 3.
    tree1 := &TreeNode {
        1,
        nil,
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, &TreeNode{5, nil, nil}, }, },
    }
    fmt.Println(longestConsecutive(tree1)) // 3
    // Example 2:
    //     (2)
    //       \
    //        (3)
    //        /
    //       2
    //     / 
    //    1
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/consec1-2-tree.jpg" />
    // Input: root = [2,null,3,2,null,1]
    // Output: 2
    // Explanation: Longest consecutive sequence path is 2-3, not 3-2-1, so return 2.
    tree2 := &TreeNode {
        2,
        nil,
        &TreeNode{3, &TreeNode{ 2, &TreeNode{ 1, nil, nil}, nil}, nil },
    }
    fmt.Println(longestConsecutive(tree2)) // 2

    fmt.Println(longestConsecutive1(tree1)) // 3
    fmt.Println(longestConsecutive1(tree2)) // 2
}