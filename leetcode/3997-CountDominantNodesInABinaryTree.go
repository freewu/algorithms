package main

// 3997. Count Dominant Nodes in a Binary Tree
// You are given the root of a complete binary tree.

// A node x is called dominant if its value is equal to the maximum value among all nodes in the subtree rooted at x.

// Return the number of dominant nodes in the tree.

// Example 1:
//             5
//            / \
//           3   8
//          / \   / \
//         2   4 7   1
// <img src="https://assets.leetcode.com/uploads/2026/06/13/tnew.png" />
// Input: root = [5,3,8,2,4,7,1]
// Output: 5
// Explanation:
// The leaf nodes with values 2, 4, 7, and 1 are dominant.
// The node with value 8 is dominant because its value is the maximum value in its subtree [8, 7, 1].
// Thus, the answer is 5.

// Example 2:
//             1
//            / \
//           2   3
//          / \
//         1   2
// <img src="https://assets.leetcode.com/uploads/2026/06/15/t9.png" />
// Input: root = [1,2,3,1,2]
// Output: 4
// Explanation:
// The leaf nodes with values 1, 2, and 3 are dominant.
// The node with value 2 whose subtree is [2, 1, 2] is dominant because its value is the maximum value in its subtree.
// Thus, the answer is 4.

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     1 <= Node.val <= 10^9
//     The tree is guaranteed to be a complete binary tree.

import "fmt"

type TreeNode struct {
    Val   int
    Left  *TreeNode
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
func countDominantNodes(root *TreeNode) int {
    res := 0
    var dfs func(t *TreeNode) int
    dfs = func(t *TreeNode) int {
        if t == nil { return -1 }
        if t.Left == nil && t.Right == nil {
            res ++
            return t.Val
        }
        s := max(dfs(t.Left), dfs(t.Right))
        if t.Val >= s {
            res ++
        }
        return max(s, t.Val)
    }
    dfs(root)
    return res
}

func countDominantNodes1(root *TreeNode) int {
    var subCount func(node *TreeNode) (int, int)
    subCount = func(node *TreeNode) (int, int) { // Return num dominant and max.
        if node == nil { return 0, 0 }
        leftDom, leftMax := subCount(node.Left)
        rightDom, rightMax := subCount(node.Right)
        numDom := leftDom + rightDom
        subMax := max(leftMax, rightMax)
        totMax := max(subMax, node.Val)
        if node.Val >= subMax {
            numDom++
        }
        return numDom, totMax
    }
    res, _ := subCount(root)
    return res
}

func main() {
    // Example 1:
    //             5
    //            / \
    //           3   8
    //          / \   / \
    //         2   4 7   1
    // <img src="https://assets.leetcode.com/uploads/2026/06/13/tnew.png" />
    // Input: root = [5,3,8,2,4,7,1]
    // Output: 5
    // Explanation:
    // The leaf nodes with values 2, 4, 7, and 1 are dominant.
    // The node with value 8 is dominant because its value is the maximum value in its subtree [8, 7, 1].
    // Thus, the answer is 5.
    tree1 := &TreeNode{
        5, 
        &TreeNode{ 3, 
            &TreeNode{ 2, nil, nil, }, 
            &TreeNode{ 4, nil, nil, }, 
        },
        &TreeNode{ 8, 
            &TreeNode{ 7, nil, nil, }, 
            &TreeNode{ 1, nil, nil, }, 
        },
    }
    fmt.Println(countDominantNodes(tree1)) // 5
    // Example 2:
    //             1
    //            / \
    //           2   3
    //          / \
    //         1   2
    // <img src="https://assets.leetcode.com/uploads/2026/06/15/t9.png" />
    // Input: root = [1,2,3,1,2]
    // Output: 4
    // Explanation:
    // The leaf nodes with values 1, 2, and 3 are dominant.
    // The node with value 2 whose subtree is [2, 1, 2] is dominant because its value is the maximum value in its subtree.
    // Thus, the answer is 4.
    tree2 := &TreeNode{
        1, 
        &TreeNode{ 2, 
            &TreeNode{ 1, nil, nil, }, 
            &TreeNode{ 2, nil, nil, }, 
        },
        &TreeNode{ 3, nil, nil, },
    }
    fmt.Println(countDominantNodes(tree2)) // 4

    tree11 := &TreeNode{
        5, 
        &TreeNode{ 3, 
            &TreeNode{ 2, nil, nil, }, 
            &TreeNode{ 4, nil, nil, }, 
        },
        &TreeNode{ 8, 
            &TreeNode{ 7, nil, nil, }, 
            &TreeNode{ 1, nil, nil, }, 
        },
    }
    fmt.Println(countDominantNodes1(tree11)) // 5
    tree12 := &TreeNode{
        1, 
        &TreeNode{ 2, 
            &TreeNode{ 1, nil, nil, }, 
            &TreeNode{ 2, nil, nil, }, 
        },
        &TreeNode{ 3, nil, nil, },
    }
    fmt.Println(countDominantNodes1(tree12)) // 4
}