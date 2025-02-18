package main

// 面试题 04.04. Check Balance LCCI
// Implement a function to check if a binary tree is balanced. 
// For the purposes of this question, a balanced tree is defined to be a tree such that the heights of the two subtrees of any node never differ by more than one.

// Example 1:
// Given tree [3,9,20,null,null,15,7]
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return true.

// Example 2:
// Given [1,2,2,3,3,null,null,4,4]
//       1
//      / \
//     2   2
//    / \
//   3   3
//  / \
// 4   4
// return false.

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
func isBalanced(root *TreeNode) bool { // 自顶向下的递归
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var height func(node *TreeNode) int // 获取树的高度
    height = func(node *TreeNode) int {
        if node == nil { return 0 }
        return max(height(node.Left), height(node.Right)) + 1
    }
    var dfs func(node *TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil { return true }
        return abs(height(node.Left) - height(node.Right)) <= 1 && dfs(node.Left) && dfs(node.Right)
    }
    return dfs(root)
}

// 自底向上的递归
func isBalanced1(root *TreeNode) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var height func(node *TreeNode) int
    height = func(node *TreeNode) int {
        if node == nil { return 0 }
        leftHeight, rightHeight := height(node.Left), height(node.Right)
        if leftHeight == -1 || rightHeight == -1 || abs(leftHeight - rightHeight) > 1 {
            return -1
        }
        return max(leftHeight, rightHeight) + 1
    }
    return height(root) >= 0
}

func main() {
    // Example 1:
    // Given tree [3,9,20,null,null,15,7]
    //     3
    //    / \
    //   9  20
    //     /  \
    //    15   7
    // return true.
    tree1 := &TreeNode {
        3,
        &TreeNode{9, nil, nil,},
        &TreeNode{20,  &TreeNode{15, nil, nil,},  &TreeNode{7, nil, nil,},},
    }
    fmt.Println(isBalanced(tree1)) // true
    // Example 2:
    // Given [1,2,2,3,3,null,null,4,4]
    //       1
    //      / \
    //     2   2
    //    / \
    //   3   3
    //  / \
    // 4   4
    // return false.
    tree2 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil,}, &TreeNode{4, nil, nil,},}, &TreeNode{3, nil, nil,},},
        &TreeNode{2, nil, nil,},
    }
    fmt.Println(isBalanced(tree2)) // false

    tree11 := &TreeNode {
        3,
        &TreeNode{9, nil, nil,},
        &TreeNode{20,  &TreeNode{15, nil, nil,},  &TreeNode{7, nil, nil,},},
    }
    fmt.Println(isBalanced1(tree11)) // true
    tree12 := &TreeNode {
        1,
        &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil,}, &TreeNode{4, nil, nil,},}, &TreeNode{3, nil, nil,},},
        &TreeNode{2, nil, nil,},
    }
    fmt.Println(isBalanced1(tree12)) // false
}