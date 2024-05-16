package main

// 250. Count Univalue Subtrees
// Given the root of a binary tree, return the number of uni-value subtrees.
// A uni-value subtree means all nodes of the subtree have the same value.

// Example 1:
//       5
//      /  \
//     1   (5)
//    / \    \
//  (5)  (5)  (5)
// <img src="https://assets.leetcode.com/uploads/2020/08/21/unival_e1.jpg" />
// Input: root = [5,1,5,5,5,null,5]
// Output: 4

// Example 2:
// Input: root = []
// Output: 0

// Example 3:
//       (5)
//      /   \
//    (5)    (5)
//    / \      \
//  (5)  (5)    (5)
// Input: root = [5,5,5,5,5,null,5]
// Output: 6

// Constraints:
//     The number of the node in the tree will be in the range [0, 1000].
//     -1000 <= Node.val <= 1000

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
func countUnivalSubtrees(root *TreeNode) int {
    res := 0
    var isUnivalSubtrees func(root *TreeNode) bool
    isUnivalSubtrees = func (root *TreeNode) bool {
        if root == nil {
            return true
        }
        isUnival := true
        if root.Left != nil {
            if !isUnivalSubtrees(root.Left) || root.Val != root.Left.Val {
                isUnival = false
            }
        }
        if root.Right != nil {
            if !isUnivalSubtrees(root.Right) || root.Val != root.Right.Val {
                isUnival = false
            }
        }
        if isUnival {
            res++
        }
        return isUnival
    }
    isUnivalSubtrees(root)
    return res
}

func countUnivalSubtrees1(root *TreeNode) int {
    if root == nil {
        return 0
    }
    res := 0
    var dfs func(node *TreeNode, ans *int) (int, bool)
    dfs = func(node *TreeNode, ans *int) (int, bool) {
        tag := true
        if node.Left != nil {
            t, tag1 := dfs(node.Left, ans)
            if !tag1 || t != node.Val { tag = false; }
        }
        if node.Right != nil {
            t, tag2 := dfs(node.Right, ans)
            if !tag2 || t != node.Val { tag = false; }
        }
        if tag { *ans++; }
        return node.Val, tag
    }
    dfs(root, &res)
    return res
}

func main() {
    // Example 1:
    //       5
    //      /  \
    //     1   (5)
    //    / \    \
    //  (5)  (5)  (5)
    // <img src="https://assets.leetcode.com/uploads/2020/08/21/unival_e1.jpg" />
    // Input: root = [5,1,5,5,5,null,5]
    // Output: 4
    tree1 := &TreeNode {
        5,
        &TreeNode{1, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}, },
        &TreeNode{5, nil,  &TreeNode{5, nil, nil}, },
    }
    fmt.Println(countUnivalSubtrees(tree1)) // 5
    // Example 2:
    // Input: root = []
    // Output: 0
    fmt.Println(countUnivalSubtrees(nil)) // 1
    // Example 3:
    //       (5)
    //      /   \
    //    (5)    (5)
    //    / \      \
    //  (5)  (5)    (5)
    // Input: root = [5,5,5,5,5,null,5]
    // Output: 6
    tree3 := &TreeNode {
        5,
        &TreeNode{5, &TreeNode{5, nil, nil}, &TreeNode{5, nil, nil}, },
        &TreeNode{5, nil,  &TreeNode{5, nil, nil}, },
    }
    fmt.Println(countUnivalSubtrees(tree3)) // 6

    fmt.Println(countUnivalSubtrees1(tree1)) // 4
    fmt.Println(countUnivalSubtrees1(nil)) // 0
    fmt.Println(countUnivalSubtrees1(tree3)) // 6
}