package main

// 617. Merge Two Binary Trees
// You are given two binary trees root1 and root2.
// Imagine that when you put one of them to cover the other, some nodes of the two trees are overlapped while the others are not. 
// You need to merge the two trees into a new binary tree. The merge rule is that if two nodes overlap, then sum node values up as the new value of the merged node. Otherwise, the NOT null node will be used as the node of the new tree.

// Return the merged tree.
// Note: The merging process must start from the root nodes of both trees.

// Example 1: 
//     1             2                    3
//    / \           / \                  / \
//   3   2         1   3         =>     4   5
//  /               \   \              / \   \
// 5                 4   7            5   4   7
// <img src="https://assets.leetcode.com/uploads/2021/02/05/merge.jpg" />
// Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
// Output: [3,4,5,5,4,null,7]

// Example 2:
//      1      1                2
//            /      =>        /
//           2                2
// Input: root1 = [1], root2 = [1,2]
// Output: [2,2]
 
// Constraints:
//     The number of nodes in both trees is in the range [0, 2000].
//     -10^4 <= Node.val <= 10^4

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
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil { return root2; }
    if root2 == nil { return root1; }
    var merge func(root1 *TreeNode, root2 *TreeNode) 
    merge = func (root1 *TreeNode, root2 *TreeNode) {
        root1.Val += root2.Val
        if root1.Left != nil && root2.Left != nil {
            merge(root1.Left,root2.Left)
        } else if root1.Left == nil {
            root1.Left = root2.Left
        }
        if root1.Right != nil && root2.Right != nil {
            merge(root1.Right,root2.Right)
        } else if root1.Right == nil{
            root1.Right = root2.Right
        } 
    }
    merge(root1,root2)
    return root1   
}

func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    if root1 == nil && root2 != nil {
        return root2
    } else if root1 != nil && root2 == nil {
        return root1
    } else if root1 != nil && root2 != nil {
        root1.Val += root2.Val
        root1.Left = mergeTrees(root1.Left, root2.Left)
        root1.Right = mergeTrees(root1.Right, root2.Right)
    }
    return root1
}

func main() {
    tree11 := &TreeNode{
        1, 
        &TreeNode{3, &TreeNode{5, nil, nil}, nil},
        &TreeNode{2, nil, nil},
    }
    tree12 := &TreeNode{
        2, 
        &TreeNode{1, nil, &TreeNode{4, nil, nil} },
        &TreeNode{3, nil,  &TreeNode{7, nil, nil} },
    }
    fmt.Println(mergeTrees(tree11,tree12).Val) // 3

    tree21 := &TreeNode{1, nil, nil}
    tree22 := &TreeNode{1, &TreeNode{2, nil, nil}, nil }
    fmt.Println(mergeTrees(tree21,tree22).Val) // 2

    tree111 := &TreeNode{
        1, 
        &TreeNode{3, &TreeNode{5, nil, nil}, nil},
        &TreeNode{2, nil, nil},
    }
    tree112 := &TreeNode{
        2, 
        &TreeNode{1, nil, &TreeNode{4, nil, nil} },
        &TreeNode{3, nil,  &TreeNode{7, nil, nil} },
    }
    fmt.Println(mergeTrees1(tree111,tree112).Val) // 3
    tree121 := &TreeNode{1, nil, nil}
    tree122 := &TreeNode{1, &TreeNode{2, nil, nil}, nil }
    fmt.Println(mergeTrees1(tree121,tree122).Val) // 2
}