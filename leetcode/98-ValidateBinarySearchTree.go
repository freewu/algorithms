package main

// 98. Validate Binary Search Tree
// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
// A valid BST is defined as follows:
//     The left subtree of a node contains only nodes with keys less than the node's key.
//     The right subtree of a node contains only nodes with keys greater than the node's key.
//     Both the left and right subtrees must also be binary search trees.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg" />
// Input: root = [2,1,3]
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg" />
// Input: root = [5,1,4,null,null,3,6]
// Output: false
// Explanation: The root node's value is 5 but its right child's value is 4.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -2^31 <= Node.val <= 2^31 - 1

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
func isValidBST(root *TreeNode) bool {
    var isBST func(root *TreeNode, prev *int) bool
    isBST = func(root *TreeNode, prev *int) bool {
        if root == nil { // If nil, it is a BST
            return true
        }
        // The idea is to do inorder traversal while maintaining the prev value
        // for any node if prev is greater than the root's value, return false
        if isBST(root.Left, prev) == false { // Check for left subtree
            return false
        }
        if root.Val <= *prev { // Compare the current node and prev value
            return false
        }
        *prev = root.Val // Update prev with current before moving ahead
        return isBST(root.Right, prev) // Check for right subtree
    }
    var prev = -1 << 63
    return isBST(root, &prev)
}

// func isValidBST1(root *TreeNode) bool {
//     if root == nil {
//         return true
//     }
//     if root.Left == nil && root.Right == nil {
//         return true
//     }
   
//     hitL := isValidBST(root.Left)
//     hitR := isValidBST(root.Right)
//     if !hitL || !hitR {
//         return false
//     }
//     maxLeft := findMax(root.Left,-math.MaxInt)
//     minRight := findMin(root.Right,math.MaxInt)
//     return hitL && hitR && (root.Val > maxLeft) && (root.Val < minRight)
// }

// func findMax(node *TreeNode, maxNum int)  int {
//     if node == nil {
//         return maxNum
//     }
//     if node.Val > maxNum {
//         maxNum = node.Val
//     }
//     maxNum = findMax(node.Left,maxNum)
//     maxNum = findMax(node.Right,maxNum)
//     return maxNum
// }

// func findMin(node *TreeNode, minNum int)  int {
//     if node == nil {
//         return minNum
//     }
//     if node.Val < minNum {
//         minNum = node.Val
//     }
//     minNum = findMin(node.Left,minNum)
//     minNum = findMin(node.Right,minNum)
//     return minNum
// }

func main() {
    tree1 := &TreeNode {
        2,
        &TreeNode{1, nil, nil},
        &TreeNode{3, nil, nil},
    }
    fmt.Println(isValidBST(tree1)) // true

    // Explanation: The root node's value is 5 but its right child's value is 4.

    tree2 := &TreeNode {
        5,
        &TreeNode{1, nil, nil},
        &TreeNode{
            4, 
            &TreeNode{3, nil, nil},
            &TreeNode{6, nil, nil},
        },
    }
    fmt.Println(isValidBST(tree2)) // false
}