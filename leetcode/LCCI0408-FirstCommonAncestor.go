package main

// 面试题 04.08. First Common Ancestor LCCI
// Design an algorithm and write code to find the first common ancestor of two nodes in a binary tree. 
// Avoid storing additional nodes in a data structure. 
// NOTE: This is not necessarily a binary search tree.

// For example, Given the following tree: root = [3,5,1,6,2,0,8,null,null,7,4]
//     3
//    / \
//   5   1
//  / \ / \
// 6  2 0  8
//   / \
//  7   4

// Example 1:
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
// Input: 3
// Explanation: The first common ancestor of node 5 and node 1 is node 3.

// Example 2:
// Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
// Output: 5
// Explanation: The first common ancestor of node 5 and node 4 is node 5.

// Notes:
//     All node values are pairwise distinct.
//     p, q are different node and both can be found in the given tree.

import "fmt"

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
func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
    var dfs func(root, p, q *TreeNode) *TreeNode
    dfs = func(root, p, q *TreeNode) *TreeNode {
        if root == nil || root == p || root == q { return root } // 公共祖先表示：左子树出现p、右子树出现q；或者左子树出现q，右子树出现p
        left, right := dfs(root.Left, p, q), dfs(root.Right, p, q) // /遍历左,右子树
        if left != nil && right != nil { // 判断是否找到公共祖先 如果左右子树均找到了，则当前节点为公共祖先节点
            return root
        }
        if left != nil { // p、q 均在左子树上，则继续往下找
            return left
        }
        return right
    }
    return dfs(root, p, q)
}

func main() {
    // Example 1:
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
    // Input: 3
    // Explanation: The first common ancestor of node 5 and node 1 is node 3.
    tree1 := &TreeNode{
        3,
        &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{2, &TreeNode{7, nil, nil, }, &TreeNode{ 4, nil, nil, }, }, },
        &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{8, nil, nil, }, },
    }
    fmt.Println(lowestCommonAncestor(tree1, &TreeNode{5, nil, nil, }, &TreeNode{1, nil, nil, })) // 3
    // Example 2:
    // Input: root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 4
    // Output: 5
    // Explanation: The first common ancestor of node 5 and node 4 is node 5.
    tree2 := &TreeNode{
        3,
        &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{2, &TreeNode{7, nil, nil, }, &TreeNode{ 4, nil, nil, }, }, },
        &TreeNode{1, &TreeNode{0, nil, nil, }, &TreeNode{8, nil, nil, }, },
    }
    fmt.Println(lowestCommonAncestor(tree2, &TreeNode{5, nil, nil, }, &TreeNode{4, nil, nil, })) // 5
}