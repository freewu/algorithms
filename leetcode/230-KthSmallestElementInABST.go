package main

// 230. Kth Smallest Element in a BST
// Given the root of a binary search tree, and an integer k, 
// return the kth smallest value (1-indexed) of all the values of the nodes in the tree.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2021/01/28/kthtree1.jpg" / >
//        3
//       /  \
//       1   4
//        \
//         2
// Input: root = [3,1,4,null,2], k = 1
// Output: 1

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2021/01/28/kthtree2.jpg" / >
//                 5
//                / \
//               3   6
//              / \
//             2   4
//            /
//           1
// Input: root = [5,3,6,2,4,null,null,1], k = 3
// Output: 3

// Constraints:
//     The number of nodes in the tree is n.
//     1 <= k <= n <= 10^4
//     0 <= Node.val <= 10^4
 
// Follow up: 
//     If the BST is modified often (i.e., we can do insert and delete operations) 
//     and you need to find the kth smallest frequently, how would you optimize?

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
func kthSmallest(root *TreeNode, k int) int {
    var inorder func (root *TreeNode ,k, level int) (int,int)
    inorder = func (root *TreeNode ,k, level int) (int,int) {
        if root == nil {
            return -1, level
        }
        res :=0
        res, level = inorder(root.Left, k, level)
        if res != -1 {
            return res, level
        }
        if level + 1 == k {
            return root.Val, level
        }
        return inorder(root.Right, k, level + 1)
    }
    res, _ := inorder(root, k, 0)
    return res
}

// stack
func kthSmallest1(root *TreeNode, k int) int {
    stack := []*TreeNode{}
    for {
        for root != nil {
            stack = append(stack, root)
            root = root.Left
        }
        stack, root = stack[:len(stack) - 1], stack[len(stack)-1]
        k--
        if k == 0 {
            return root.Val
        }
        root = root.Right
    }
}

func main() {
    tree1 := &TreeNode {
        3,
        &TreeNode { 
            1, 
            nil,
            &TreeNode{2, nil, nil},
        },
        &TreeNode{4, nil, nil},
    }
    fmt.Println(kthSmallest(tree1, 1)) // 1

    tree2 := &TreeNode {
        5,
        &TreeNode { 
            3, 
            &TreeNode { 
                2, 
                &TreeNode{1, nil, nil},
                nil,
            },
            &TreeNode{4, nil, nil},
        },
        &TreeNode{6, nil, nil},
    }
    fmt.Println(kthSmallest(tree2, 3)) // 3

    fmt.Println(kthSmallest1(tree1, 1)) // 1
    fmt.Println(kthSmallest1(tree2, 3)) // 3
}