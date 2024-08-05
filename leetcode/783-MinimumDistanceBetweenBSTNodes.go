package main

// 783. Minimum Distance Between BST Nodes
// Given the root of a Binary Search Tree (BST), 
// return the minimum difference between the values of any two different nodes in the tree.

// Example 1:
//         4
//       /   \
//      2     6
//     /  \
//    1    3
// <img src="https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg" />
// Input: root = [4,2,6,1,3]
// Output: 1

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg" />
//         1
//       /   \
//      0    48
//          /   \
//         12    49
// Input: root = [1,0,48,null,null,12,49]
// Output: 1

// Constraints:
//     The number of nodes in the tree is in the range [2, 100].
//     0 <= Node.val <= 10^5

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
func minDiffInBST(root *TreeNode) int {
    var prev *TreeNode
    res := 1 << 32 - 1
    min := func (x, y int) int { if x < y { return x; }; return y; }
    var visit func(root *TreeNode)
    visit = func(root *TreeNode) {
        if root == nil {
            return
        }
        visit(root.Left)
        if prev != nil { // inorder 中序遍历
            res = min(res, (root.Val - prev.Val))
        }
        prev = root
        visit(root.Right)
    }
    visit(root)
    return res
}

func main() {
    // Example 1:
    //         4
    //       /   \
    //      2     6
    //     /  \
    //    1    3
    // <img src="https://assets.leetcode.com/uploads/2021/02/05/bst1.jpg" />
    // Input: root = [4,2,6,1,3]
    // Output: 1
    tree1 := &TreeNode{
        4, 
        &TreeNode{2, &TreeNode{1, nil, nil, }, &TreeNode{3, nil, nil, }, },
        &TreeNode{6, nil, nil, },
    }
    fmt.Println(minDiffInBST(tree1)) // 1
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2021/02/05/bst2.jpg" />
    //         1
    //       /   \
    //      0    48
    //          /   \
    //         12    49
    // Input: root = [1,0,48,null,null,12,49]
    // Output: 1
    tree2 := &TreeNode{
        1, 
        &TreeNode{0, nil, nil, },
        &TreeNode{48, &TreeNode{12, nil, nil, }, &TreeNode{49, nil, nil, }, },
    }
    fmt.Println(minDiffInBST(tree2)) // 1
}