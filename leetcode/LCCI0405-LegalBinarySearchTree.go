package main

// 面试题 04.05. Legal Binary Search Tree LCCI
// Implement a function to check if a binary tree is a binary search tree.

// Example 1:
// Input:
//     2
//    / \
//   1   3
// Output: true

// Example 2:
// Input:
//     5
//    / \
//   1   4
//      / \
//     3   6
// Output: false
// Explanation: Input: [5,1,4,null,null,3,6].
// the value of root node is 5, but its right child has value 4.

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
func isValidBST(root *TreeNode) bool {
    var dfs func(node *TreeNode, left, right int) bool
    dfs = func(node *TreeNode, left, right int) bool { // 前序遍历
        if node == nil { return true }
        x := node.Val 
        return left < x && x < right && dfs(node.Left, left, x) && dfs(node.Right, x, right)
    }
    return dfs(root, -1 << 61, 1 << 61)
}

func isValidBST1(root *TreeNode) bool {
    pre := -1 << 61
    var dfs func(*TreeNode) bool
    dfs = func(node *TreeNode) bool { // 中序遍历
        if node == nil { return true }
        if !dfs(node.Left) || node.Val <= pre { return false }
        pre = node.Val
        return dfs(node.Right)
    }
    return dfs(root)
}

func isValidBST2(root *TreeNode) bool {
    var dfs func(node *TreeNode) (int, int)
    dfs = func(node *TreeNode) (int, int) {
        if node == nil { return 1 << 61, -1 << 61 }
        lmn, lmx := dfs(node.Left)
        rmn, rmx := dfs(node.Right)
        v := node.Val
        if v <= lmx || v >= rmn { return -1 << 61, 1 << 61 }
        return min(lmn, v), max(rmx, v)
    }
    _, mx := dfs(root)
    return mx != 1 << 61
}

func main() {
    // Example 1:
    // Input:
    //     2
    //    / \
    //   1   3
    // Output: true
    tree1 := &TreeNode {
        2,
        &TreeNode{1, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(isValidBST(tree1)) // true
    // Example 2:
    // Input:
    //     5
    //    / \
    //   1   4
    //      / \
    //     3   6
    // Output: false
    // Explanation: Input: [5,1,4,null,null,3,6].
    // the value of root node is 5, but its right child has value 4.
    tree2 := &TreeNode {
        5,
        &TreeNode{1, nil, nil, },
        &TreeNode{4, &TreeNode{3, nil, nil, }, &TreeNode{6, nil, nil, }, },
    }
    fmt.Println(isValidBST(tree2)) // false

    fmt.Println(isValidBST1(tree1)) // true
    fmt.Println(isValidBST1(tree2)) // false

    fmt.Println(isValidBST2(tree1)) // true
    fmt.Println(isValidBST2(tree2)) // false
}