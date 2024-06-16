package main

// 549. Binary Tree Longest Consecutive Sequence II
// Given the root of a binary tree, return the length of the longest consecutive path in the tree.

// A consecutive path is a path where the values of the consecutive nodes in the path differ by one. 
// This path can be either increasing or decreasing.
//     For example, [1,2,3,4] and [4,3,2,1] are both considered valid, but the path [1,2,4,3] is not valid.

// On the other hand, the path can be in the child-Parent-child order, where not necessarily be parent-child order.

// Example 1:
//         1
//        /  \
//       2    3
// <img src="https://assets.leetcode.com/uploads/2021/03/14/consec2-1-tree.jpg" />
// Input: root = [1,2,3]
// Output: 2
// Explanation: The longest consecutive path is [1, 2] or [2, 1].

// Example 2:
//         2
//        /  \
//       1    3
// <img src="https://assets.leetcode.com/uploads/2021/03/14/consec2-2-tree.jpg" />
// Input: root = [2,1,3]
// Output: 3
// Explanation: The longest consecutive path is [1, 2, 3] or [3, 2, 1].
 
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

func longestConsecutive(root *TreeNode) int {
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var cal func(n *TreeNode, d int) int
    cal = func (n *TreeNode, d int) int {
        if nil == n {
            return 0
        }
        l, r := 0, 0
        if nil != n.Left && n.Val-n.Left.Val == d {
            l = cal(n.Left, d) + 1
        }
        if nil != n.Right && n.Val-n.Right.Val == d {
            r = cal(n.Right, d) + 1
        }
        return max(l, r)
    }
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if nil == root {
            return 0
        }
        return max(cal(root, 1) + cal(root, -1) + 1, max(dfs(root.Left), dfs(root.Right)))
    }
    return dfs(root)
}

func longestConsecutive1(root *TreeNode) int {
    res := 0
    if root == nil {
        return res
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node *TreeNode, val int) (int, int)
    dfs = func(node *TreeNode, val int) (int, int) {
        if node == nil {
            return 0, 0
        }
        lAsc, lDesc := dfs(node.Left, node.Val)
        rAsc, rDesc := dfs(node.Right, node.Val)
        res = max(res, 1 + lAsc + rDesc)
        res = max(res, 1 + lDesc + rAsc)
        if node.Val == val + 1 {
            return 1 + max(lAsc, rAsc), 0
        } else if node.Val == val - 1 {
            return 0, 1 + max(lDesc, rDesc)
        }
        return 0, 0
    }
    dfs(root, 1 << 32 - 1)
    return res
}

func main() {
    // Example 1:
    //         1
    //        /  \
    //       2    3
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/consec2-1-tree.jpg" />
    // Input: root = [1,2,3]
    // Output: 2
    // Explanation: The longest consecutive path is [1, 2] or [2, 1].
    tree1 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(longestConsecutive(tree1)) // 2
    // Example 2:
    //         2
    //        /  \
    //       1    3
    // <img src="https://assets.leetcode.com/uploads/2021/03/14/consec2-2-tree.jpg" />
    // Input: root = [2,1,3]
    // Output: 3
    // Explanation: The longest consecutive path is [1, 2, 3] or [3, 2, 1].
    tree2 := &TreeNode{
        2, 
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(longestConsecutive(tree2)) // 3

    fmt.Println(longestConsecutive1(tree1)) // 2
    fmt.Println(longestConsecutive1(tree2)) // 3
}