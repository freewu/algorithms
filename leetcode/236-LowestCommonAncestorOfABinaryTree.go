package main

// 236. Lowest Common Ancestor of a Binary Tree
// Given a binary tree, find the lowest common ancestor (LCA) of two given nodes in the tree.
// According to the definition of LCA on Wikipedia: 
//     The lowest common ancestor is defined between two nodes p 
//     and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Output: 3
// Explanation: The LCA of nodes 5 and 1 is 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2018/12/14/binarytree.png" />
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The LCA of nodes 5 and 4 is 5, since a node can be a descendant of itself according to the LCA definition.

// Example 3:
// Input: root = [1,2], p = 1, q = 2
// Output: 1
 
// Constraints:
//     The number of nodes in the tree is in the range [2, 10^5].
//     -10^9 <= Node.val <= 10^9
//     All Node.val are unique.
//     p != q
//     p and q will exist in the tree.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// 对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，
// 满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // 一个节点也可以是它自己的祖先
    if root == nil || root == q || root == p {
        return root
    }
    // 递归查找 left / right
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    fmt.Println("root: ", root)
    fmt.Println("left: ", left)
    fmt.Println("right: ", right)
    // 如果 left 和 right 都不为 nil 返回 root
    if left != nil && right != nil {
        return root
    }
    if left == nil {
        return right
    }
    return left
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode {
            5,
            &TreeNode{6, nil, nil},
            &TreeNode{
                2, 
                &TreeNode{7, nil, nil},
                &TreeNode{4, nil, nil},
            },
        },
        &TreeNode {
            1,
            &TreeNode{0, nil, nil},
            &TreeNode{8, nil, nil},
        },
    }
    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{5, nil, nil},&TreeNode{1, nil, nil})) // 3
    fmt.Println(lowestCommonAncestor(tree1,&TreeNode{5, nil, nil},&TreeNode{4, nil, nil})) // 5

    tree2 := &TreeNode {
        1,
        &TreeNode{2, nil, nil},
        nil,
    }
    //fmt.Println(lowestCommonAncestor(tree2,&TreeNode{1, nil, nil},&TreeNode{2, nil, nil})) // 1
    fmt.Println(lowestCommonAncestor(tree2,&TreeNode{1, nil, nil},&TreeNode{2, nil, nil})) // 1
}