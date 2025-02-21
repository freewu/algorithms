package main

// 面试题 17.12. BiNode LCCI
// The data structure TreeNode is used for binary tree, but it can also used to represent a single linked list (where left is null, and right is the next node in the list). 
// Implement a method to convert a binary search tree (implemented with TreeNode) into a single linked list. 
// The values should be kept in order and the operation should be performed in place (that is, on the original data structure).

// Return the head node of the linked list after converting.

// Note: This problem is slightly different from the original one in the book.

// Example:
// Input:  [4,2,5,1,3,null,6,0]
// Output:  [0,null,1,null,2,null,3,null,4,null,5,null,6]

// Note:
//     The number of nodes will not exceed 100000.

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
func convertBiNode(root *TreeNode) *TreeNode {
    dummy := &TreeNode{ Val: -1 }
    var pre *TreeNode
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil { return }
        inorder(node.Left)
        if pre == nil {
            pre = node
            dummy.Right = node
        } else {
            pre.Right = node
            pre = node
        }
        node.Left = nil
        inorder(node.Right)
    }
    inorder(root)
    return dummy.Right
}

func main() {
    // Example:
    // Input:  [4,2,5,1,3,null,6,0]
    // Output:  [0,null,1,null,2,null,3,null,4,null,5,null,6]
    tree1 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, &TreeNode{0, nil, nil, }, nil, }, &TreeNode{3, nil, nil, }, },
        &TreeNode{5, nil,&TreeNode{6, nil, nil, }, },
    }
    fmt.Println(convertBiNode(tree1)) // 
}