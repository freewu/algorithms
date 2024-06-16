package main

// 1120. Maximum Average Subtree
// Given the root of a binary tree, return the maximum average value of a subtree of that tree. 
// Answers within 10^-5 of the actual answer will be accepted.

// A subtree of a tree is any node of that tree plus all its descendants.
// The average value of a tree is the sum of its values, divided by the number of nodes.

// Example 1:
//         5
//       /   \
//      6     1
// <img src="https://assets.leetcode.com/uploads/2019/04/09/1308_example_1.png" />
// Input: root = [5,6,1]
// Output: 6.00000
// Explanation: 
// For the node with value = 5 we have an average of (5 + 6 + 1) / 3 = 4.
// For the node with value = 6 we have an average of 6 / 1 = 6.
// For the node with value = 1 we have an average of 1 / 1 = 1.
// So the answer is 6 which is the maximum.

// Example 2:
// Input: root = [0,null,1]
// Output: 1.00000

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     0 <= Node.val <= 10^5

import "fmt"
import "math"

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
func maximumAverageSubtree(root *TreeNode) float64 {
    res := float64(0.0)
    var dfs func(node *TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        left, right := dfs(node.Left), dfs(node.Right)
        if node.Left != nil  {  node.Val += node.Left.Val;  }
        if node.Right != nil {  node.Val += node.Right.Val; }
        count := left + right + 1
        res = math.Max(res, float64(node.Val) / float64(count))
        return count
    }
    dfs(root)
    return res
}

func maximumAverageSubtree1(root *TreeNode) float64 {
    maxAverage := float64(0)
    var ptr *float64 = &maxAverage 
    var dfs func(node *TreeNode,ptr *float64) [2]float64
    dfs = func (node *TreeNode,ptr *float64) [2]float64 {
        left, right, num := [2]float64{ 0,0 }, [2]float64{ 0,0 }, [2]float64{ 0,0 }
        if node.Left != nil {  left = dfs(node.Left, ptr) }
        if node.Right != nil { right= dfs(node.Right,ptr) }
        num[0] = (left[0] * left[1] + right[0] * right[1] + float64(node.Val)) / (left[1] + right[1] + 1)
        num[1] = (left[1] + right[1] + 1)
        if num[0] > *ptr {
            *ptr = num[0]
        }
        return num
    }
    dfs(root, ptr)
    return maxAverage
}

func main() {
    // Example 1:
    //         5
    //       /   \
    //      6     1
    // <img src="https://assets.leetcode.com/uploads/2019/04/09/1308_example_1.png" />
    // Input: root = [5,6,1]
    // Output: 6.00000
    // Explanation: 
    // For the node with value = 5 we have an average of (5 + 6 + 1) / 3 = 4.
    // For the node with value = 6 we have an average of 6 / 1 = 6.
    // For the node with value = 1 we have an average of 1 / 1 = 1.
    // So the answer is 6 which is the maximum.
    tree1 := &TreeNode {
        5,
        &TreeNode{6, nil, nil, },
        &TreeNode{1, nil, nil, },
    }
    fmt.Println(maximumAverageSubtree(tree1)) // 6.000
    // Example 2:
    // Input: root = [0,null,1]
    // Output: 1.00000
    tree2 := &TreeNode {
        0,
        nil,
        &TreeNode{1, nil, nil, },
    }
    fmt.Println(maximumAverageSubtree(tree2)) // 1.000

    fmt.Println(maximumAverageSubtree1(tree1)) // 6.000
    fmt.Println(maximumAverageSubtree1(tree2)) // 1.000
}