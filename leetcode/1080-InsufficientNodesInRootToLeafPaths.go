package main

// 1080. Insufficient Nodes in Root to Leaf Paths
// Given the root of a binary tree and an integer limit, 
// delete all insufficient nodes in the tree simultaneously, 
// and return the root of the resulting binary tree.

// A node is insufficient if every root to leaf path intersecting this node has a sum strictly less than limit.

// A leaf is a node with no children.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-11.png" />
// Input: root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
// Output: [1,2,3,4,null,null,7,8,9,null,14]

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-3.png" />
// Input: root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
// Output: [5,4,8,11,null,17,4,7,null,null,null,5]

// Example 3:
// <img src="https://assets.leetcode.com/uploads/2019/06/11/screen-shot-2019-06-11-at-83301-pm.png" />
// Input: root = [1,2,-3,-5,null,4,null], limit = -1
// Output: [1,null,-3,4]

// Constraints:
//     The number of nodes in the tree is in the range [1, 5000].
//     -10^5 <= Node.val <= 10^5
//     -10^9 <= limit <= 10^9

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
// // 递归
// func sufficientSubset(root *TreeNode, limit int) *TreeNode {
//     if root.Left == nil && root.Right == nil {
//         if limit - root.Val > 0 {
//             return nil
//         }
//         return root
//     }
//     if root.Left != nil {
//         root.Left = sufficientSubset(root.Left, limit - root.Val)
//     }
//     if root.Right != nil {
//         root.Right = sufficientSubset(root.Right, limit - root.Val)
//     }
//     if root.Left == nil && root.Right == nil {
//         return nil
//     }
//     return root
// }

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
    if root == nil {
        return nil
    }
    limit -= root.Val
    if root.Left == root.Right { // root 是叶子
        if limit > 0 { // 从根到叶子的路径和小于 limit，删除叶子
            return nil
        }
        return root // 否则不删除
    }
    root.Left, root.Right = sufficientSubset(root.Left, limit), sufficientSubset(root.Right, limit)
    if root.Left ==  root.Right { // 如果儿子都被删除，就删 root
        return nil
    }
    return root // 否则不删 root
}

// dfs
func sufficientSubset1(root *TreeNode, limit int) *TreeNode {
    var dfs func(root *TreeNode, limit int) *TreeNode
    dfs = func(root *TreeNode, limit int) *TreeNode {
        if root.Left == nil && root.Right == nil {
            if limit - root.Val > 0 {
                return nil
            }
            return root
        }
        if root.Left != nil {
            root.Left = dfs(root.Left, limit - root.Val)
        }
        if root.Right != nil {
            root.Right = dfs(root.Right, limit - root.Val)
        }
        if root.Left == nil && root.Right == nil {
            return nil
        }
        return root
    }
    return dfs(root, limit)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-11.png" />
    // Input: root = [1,2,3,4,-99,-99,7,8,9,-99,-99,12,13,-99,14], limit = 1
    // Output: [1,2,3,4,null,null,7,8,9,null,14]
    tree1 := &TreeNode{
        1, 
        &TreeNode{ 2, &TreeNode{ 4,   &TreeNode{ 9,  nil, nil, }, &TreeNode{  9, nil, nil, }, }, &TreeNode{ -99, &TreeNode{ -99, nil, nil, }, &TreeNode{ -99, nil, nil, }, }, },
        &TreeNode{ 3, &TreeNode{ -99, &TreeNode{ 12, nil, nil, }, &TreeNode{ 13, nil, nil, }, }, &TreeNode{ 7,   &TreeNode{ -99, nil, nil, }, &TreeNode{ 14,  nil, nil, }, }, },
    }
    fmt.Println(sufficientSubset(tree1, 1)) // &{1 0xc000094048 0xc0000940f0} [1,2,3,4,null,null,7,8,9,null,14]
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2019/06/05/insufficient-3.png" />
    // Input: root = [5,4,8,11,null,17,4,7,1,null,null,5,3], limit = 22
    // Output: [5,4,8,11,null,17,4,7,null,null,null,5]
    tree2 := &TreeNode{
        5, 
        &TreeNode{ 4, &TreeNode{ 11, &TreeNode{ 7, nil, nil, }, &TreeNode{ 1, nil, nil, } }, nil, },
        &TreeNode{ 8, &TreeNode{ 17, nil, nil, }, &TreeNode{ 4,&TreeNode{ 5, nil, nil, }, &TreeNode{ 3, nil, nil, }, }, },
    }
    fmt.Println(sufficientSubset(tree2, 22)) // &{5 0xc0000941c8 0xc000094228} [5,4,8,11,null,17,4,7,null,null,null,5]
    // Example 3:
    // <img src="https://assets.leetcode.com/uploads/2019/06/11/screen-shot-2019-06-11-at-83301-pm.png" />
    // Input: root = [1,2,-3,-5,null,4,null], limit = -1
    // Output: [1,null,-3,4]
    tree3 := &TreeNode{
        1, 
        &TreeNode{ 2,  &TreeNode{ -5, nil, nil, }, nil, },
        &TreeNode{ -3, &TreeNode{  4, nil, nil, }, nil, },
    }
    fmt.Println(sufficientSubset(tree3, -1)) // &{1 <nil> 0xc000094300} [1,null,-3,4]

    fmt.Println(sufficientSubset1(tree1, 1)) // &{1 0xc000094048 0xc0000940f0} [1,2,3,4,null,null,7,8,9,null,14]
    fmt.Println(sufficientSubset1(tree2, 22)) // &{5 0xc0000941c8 0xc000094228} [5,4,8,11,null,17,4,7,null,null,null,5]
    fmt.Println(sufficientSubset1(tree3, -1)) // &{1 <nil> 0xc000094300} [1,null,-3,4]
}