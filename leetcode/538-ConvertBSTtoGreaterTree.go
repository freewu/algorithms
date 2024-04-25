package main

// 538. Convert BST to Greater Tree
// Given the root of a Binary Search Tree (BST), 
// convert it to a Greater Tree such that every key of the original BST is changed to the original key plus the sum of all keys greater than the original key in BST.
// As a reminder, a binary search tree is a tree that satisfies these constraints:
//     The left subtree of a node contains only nodes with keys less than the node's key.
//     The right subtree of a node contains only nodes with keys greater than the node's key.
//     Both the left and right subtrees must also be binary search trees.
 
// Example 1:
// Input: root = [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
// Output: [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]

// Example 2:
// Input: root = [0,null,1]
// Output: [1,null,1]
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 10^4].
//     -10^4 <= Node.val <= 10^4
//     All the values in the tree are unique.
//     root is guaranteed to be a valid binary search tree.

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
// func convertBST(root *TreeNode) *TreeNode {
//     sum := 0
//     var inOrder func(node *TreeNode,sum *int) 
//     inOrder = func (node *TreeNode,sum *int) {
//         if node == nil {
//             return
//         }
//         inOrder(node.Right,sum)
//         node.Val += *sum
//         *sum = node.Val
//         inOrder(node.Left,sum)
//     }
//     inOrder(root,&sum)
//     return root
// }

func convertBST(root *TreeNode) *TreeNode {
    var traverse func(root *TreeNode, sum int) int 
    traverse = func (root *TreeNode, sum int) int {
        if root == nil {
            return sum
        }
        // right
        sum = traverse(root.Right, sum)
        root.Val += sum
        // left
        return traverse(root.Left, root.Val)
    }
    traverse(root, 0)
    return root
}




func main() {
    // [4,1,6,0,2,5,7,null,null,null,3,null,null,null,8]
    tree1 := &TreeNode {
        4,
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, &TreeNode{3, nil, nil}, }, },
        &TreeNode{6, &TreeNode{5, nil, nil}, &TreeNode{7, nil, &TreeNode{8, nil, nil}, }, },
    }
    fmt.Println(tree1.Val) // 4
    fmt.Println(convertBST(tree1)) // [30,36,21,36,35,26,15,null,null,null,33,null,null,null,8]
    fmt.Println(tree1.Val) // 30

    // [0,null,1] 
    tree2 := &TreeNode {
        0,
        nil,
        &TreeNode{1, nil, nil},
    }
    fmt.Println(tree2.Val) // 0
    fmt.Println(convertBST(tree2)) // [1,null,1]
    fmt.Println(tree2.Val) // 1
}