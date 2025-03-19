package main

// 1339. Maximum Product of Splitted Binary Tree
// Given the root of a binary tree, split the binary tree into two subtrees by removing one edge 
// such that the product of the sums of the subtrees is maximized.

// Return the maximum product of the sums of the two subtrees. Since the answer may be too large, return it modulo 10^9 + 7.
// Note that you need to maximize the answer before taking the mod and not after taking it.

// Example 1:
//         1                               1
//       /   \                               \
//      2      3  =>          2               3 
//    /   \   /             /   \            /
//   4    5   6            4     5          6
//                           11             10       =  110
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_1_1699.png" />
// Input: root = [1,2,3,4,5,6]
// Output: 110
// Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11*10)

// Example 2:
//       1 
//         \ 
//          2                 1
//         /  \       =>       \               4
//        3    4                2            /   \
//           /   \             /            5     6
//          5     6           3  
//                            6      *         15       =  90
// <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_2_1699.png" />
// Input: root = [1,null,2,3,4,null,null,5,6]
// Output: 90
// Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15*6)

// Constraints:
//     The number of nodes in the tree is in the range [2, 5 * 10^4].
//     1 <= Node.val <= 10^4

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
func maxProduct(root *TreeNode) int {
    dp := []int{}
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0 }
        res := node.Val + dfs(node.Left) + dfs(node.Right)
        dp = append(dp, res)
        return res
    }
    res, sum := 0, dfs(root)
    for _, v := range dp {
        cand := v * (sum - v)
        if cand > res { 
            res = cand 
        }
    }
    return res % 1_000_000_007
}

func maxProduct1(root *TreeNode) int {
    var sumTree func(root *TreeNode) int
    sumTree = func(root *TreeNode) int {
        if root == nil { return 0 }
        return root.Val + sumTree(root.Left) + sumTree(root.Right)
    }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    res, total := 0, sumTree(root)
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0}
        left, right := dfs(node.Left), dfs(node.Right)
        res = max(res, (total - left) * left)
        res = max(res, (total - right) * right)
        return node.Val + left + right
    }
    dfs(root)
    return res % 1000000007
}

func main() {
    // Example 1:
    //         1                               1
    //       /   \                               \
    //      2      3  =>          2               3 
    //    /   \   /             /   \            /
    //   4    5  6             4     5          6
    //                           11             10       =  110
    // <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_1_1699.png" />
    // Input: root = [1,2,3,4,5,6]
    // Output: 110
    // Explanation: Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11*10)
    tree1 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode { 4, nil, nil }, &TreeNode { 5, nil, nil } },
        &TreeNode { 3, &TreeNode { 6, nil, nil }, nil },
    }
    fmt.Println(maxProduct(tree1)) // 110
    // Example 2:
    //       1 
    //         \ 
    //          2                 1
    //         /  \       =>       \               4
    //        3    4                2            /   \
    //           /   \             /            5     6
    //          5     6           3  
    //                            6      *         15       =  90
    // <img src="https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2020/02/02/sample_2_1699.png" />
    // Input: root = [1,null,2,3,4,null,null,5,6]
    // Output: 90
    // Explanation: Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15*6)
    tree2 := &TreeNode {
        1,
        nil,
        &TreeNode { 2, &TreeNode { 3, nil, nil }, &TreeNode { 4, &TreeNode { 5, nil, nil }, &TreeNode { 6, nil, nil } } },
    }
    fmt.Println(maxProduct(tree2)) // 90

    tree11 := &TreeNode {
        1,
        &TreeNode { 2, &TreeNode { 4, nil, nil }, &TreeNode { 5, nil, nil } },
        &TreeNode { 3, &TreeNode { 6, nil, nil }, nil },
    }
    fmt.Println(maxProduct1(tree11)) // 110
    tree12 := &TreeNode {
        1,
        nil,
        &TreeNode { 2, &TreeNode { 3, nil, nil }, &TreeNode { 4, &TreeNode { 5, nil, nil }, &TreeNode { 6, nil, nil } } },
    }
    fmt.Println(maxProduct1(tree12)) // 90
}