package main

// 2265. Count Nodes Equal to Average of Subtree
// Given the root of a binary tree, 
// return the number of nodes where the value of the node is equal to the average of the values in its subtree.

// Note:
//     The average of n elements is the sum of the n elements divided by n and rounded down to the nearest integer.
//     A subtree of root is a tree consisting of root and all of its descendants.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315203925-1.png" />
// Input: root = [4,8,5,0,1,null,6]
// Output: 5
// Explanation: 
// For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4.
// For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
// For the node with value 0: The average of its subtree is 0 / 1 = 0.
// For the node with value 1: The average of its subtree is 1 / 1 = 1.
// For the node with value 6: The average of its subtree is 6 / 1 = 6.

// Example 2:
// Input: root = [1]
// Output: 1
// Explanation: For the node with value 1: The average of its subtree is 1 / 1 = 1.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     0 <= Node.val <= 1000

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
func averageOfSubtree(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        count, sum := 1, node.Val
        if nil != node.Left {
            lcount, lsum := dfs(node.Left)
            count += lcount
            sum += lsum
        }
        if nil != node.Right {
            rcount, rsum := dfs(node.Right)
            count += rcount
            sum += rsum
        }
        if node.Val == sum / count {
            res++
        }
        return count, sum
    }
    dfs(root)
    return res
}

func averageOfSubtree1(root *TreeNode) int {
    res := 0
    var dfs func(node *TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        if node == nil { return 0, 0 }
        lcount, lsum := dfs(node.Left)
        rcount, rsum := dfs(node.Right)
        count, sum := lcount + rcount + 1, lsum + rsum + node.Val
        if sum / count == node.Val {
            res++
        }
        return count, sum
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/03/15/image-20220315203925-1.png" />
    // Input: root = [4,8,5,0,1,null,6]
    // Output: 5
    // Explanation: 
    // For the node with value 4: The average of its subtree is (4 + 8 + 5 + 0 + 1 + 6) / 6 = 24 / 6 = 4.
    // For the node with value 5: The average of its subtree is (5 + 6) / 2 = 11 / 2 = 5.
    // For the node with value 0: The average of its subtree is 0 / 1 = 0.
    // For the node with value 1: The average of its subtree is 1 / 1 = 1.
    // For the node with value 6: The average of its subtree is 6 / 1 = 6.
    tree1 := &TreeNode {
        4,
        &TreeNode { 8, &TreeNode { 0, nil, nil, }, &TreeNode { 1, nil, nil, }, },
        &TreeNode { 5, nil, &TreeNode { 6, nil, nil, }, },
    }
    fmt.Println(averageOfSubtree(tree1)) // 5
    // Example 2:
    // Input: root = [1]
    // Output: 1
    // Explanation: For the node with value 1: The average of its subtree is 1 / 1 = 1.
    tree2 := &TreeNode { 1, nil, nil }
    fmt.Println(averageOfSubtree(tree2)) // 1

    fmt.Println(averageOfSubtree1(tree1)) // 5
    fmt.Println(averageOfSubtree1(tree2)) // 1
}