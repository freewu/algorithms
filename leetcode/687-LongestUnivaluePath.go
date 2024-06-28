package main

// 687. Longest Univalue Path
// Given the root of a binary tree, return the length of the longest path, where each node in the path has the same value. 
// This path may or may not pass through the root.

// The length of the path between two nodes is represented by the number of edges between them.

// Example 1:
//          (5)
//         /   \
//        4    (5)
//      /   \     \
//     4     4    (5)
// <img src="https://assets.leetcode.com/uploads/2020/10/13/ex1.jpg" />
// Input: root = [5,4,5,1,1,null,5]
// Output: 2
// Explanation: The shown image shows that the longest path of the same value (i.e. 5).

// Example 2:
//             1
//           /   \
//         (4)    5
//        /   \     \
//      (4)   (4)     5
// <img src="https://assets.leetcode.com/uploads/2020/10/13/ex2.jpg" />
// Input: root = [1,4,5,4,4,null,5]
// Output: 2
// Explanation: The shown image shows that the longest path of the same value (i.e. 4).

// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -1000 <= Node.val <= 1000
//     The depth of the tree will not exceed 1000.

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
// dfs
func longestUnivaluePath(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        v, sum := 0, 0
        if node.Left != nil {
            vl := dfs(node.Left)
            if node.Val == node.Left.Val {
                v = vl + 1
                sum += vl + 1
            }
        }
        if node.Right != nil {
            vr := dfs(node.Right)
            if node.Val == node.Right.Val {
                if vr + 1 > v {
                    v = vr + 1
                }
                sum += vr + 1
            }
        }
        if sum > res {
            res = sum
        }
        return v
    }
    dfs(root)
    return res
}

func longestUnivaluePath1(root *TreeNode) int {
    res := 0
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int{
        if node == nil {
            return -1
        }
        left := dfs(node.Left) + 1
        right := dfs(node.Right) + 1
        if node.Left != nil && node.Left.Val != node.Val {
            left = 0
        }
        if node.Right != nil && node.Right.Val != node.Val {
            right = 0
        }
        res = max(res,left + right)
        return max(left,right)
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    //          (5)
    //         /   \
    //        4    (5)
    //      /   \     \
    //     4     4    (5)
    // <img src="https://assets.leetcode.com/uploads/2020/10/13/ex1.jpg" />
    // Input: root = [5,4,5,1,1,null,5]
    // Output: 2
    // Explanation: The shown image shows that the longest path of the same value (i.e. 5).
    tree1 := &TreeNode {
        5,
        &TreeNode{4, &TreeNode{4, nil, nil, },  &TreeNode{4, nil, nil, }, },
        &TreeNode{5, nil,                       &TreeNode{5, nil, nil, }, },
    }
    fmt.Println(longestUnivaluePath(tree1)) // 2
    // Example 2:
    //             1
    //           /   \
    //         (4)    5
    //        /   \     \
    //      (4)   (4)     5
    // <img src="https://assets.leetcode.com/uploads/2020/10/13/ex2.jpg" />
    // Input: root = [1,4,5,4,4,null,5]
    // Output: 2
    // Explanation: The shown image shows that the longest path of the same value (i.e. 4).
    tree2 := &TreeNode {
        1,
        &TreeNode{4, &TreeNode{4, nil, nil, },  &TreeNode{4, nil, nil, }, },
        &TreeNode{5, nil,                       &TreeNode{5, nil, nil, }, },
    }
    fmt.Println(longestUnivaluePath(tree2)) // 2

    fmt.Println(longestUnivaluePath1(tree1)) // 2
    fmt.Println(longestUnivaluePath1(tree2)) // 2
}