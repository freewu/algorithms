package main

// 1373. Maximum Sum BST in Binary Tree
// Given a binary tree root, return the maximum sum of all keys of any sub-tree which is also a Binary Search Tree (BST).

// Assume a BST is defined as follows:
//     The left subtree of a node contains only nodes with keys less than the node's key.
//     The right subtree of a node contains only nodes with keys greater than the node's key.
//     Both the left and right subtrees must also be binary search trees.

// Example 1:
//              1
//           /    \
//          4      3
//        /   \   /   \
//       2     4 2     5
//                    /  \
//                   4    6
// <img src="https://assets.leetcode.com/uploads/2020/01/30/sample_1_1709.png" />
// Input: root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
// Output: 20
// Explanation: Maximum sum in a valid Binary search tree is obtained in root node with key equal to 3.

// Example 2:
//         4
//        /
//       3
//     /   \
//    1     2
// <img src="https://assets.leetcode.com/uploads/2020/01/30/sample_2_1709.png" />
// Input: root = [4,3,null,1,2]
// Output: 2
// Explanation: Maximum sum in a valid Binary search tree is obtained in a single root node with key equal to 2.

// Example 3:
// Input: root = [-4,-2,-5]
// Output: 0
// Explanation: All values are negatives. Return an empty BST.

// Constraints:
//     The number of nodes in the tree is in the range [1, 4 * 10^4].
//     -4 * 10^4 <= Node.val <= 4 * 10^4

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
func maxSumBST(root *TreeNode) int {
    res, inf := 0, 1 << 31
    min := func (x, y int) int { if x < y { return x; }; return y; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(root *TreeNode) (int, int, int)
    dfs = func(root *TreeNode) (int, int, int) {
        if root == nil { return inf, -inf, 0 }
        lmin, lmax, lsum := dfs(root.Left)
        rmin, rmax, rsum := dfs(root.Right)
        if root.Val <= lmax || root.Val >= rmin { // not bst
            return -inf, inf, 0
        }
        sum := lsum + rsum + root.Val
        res = max(res, sum)
        return min(lmin, root.Val), max(rmax, root.Val), sum
    }
    dfs(root)
    return res
}

func main() {
    // Example 1:
    //              1
    //           /    \
    //          4      3
    //        /   \   /   \
    //       2     4 2     5
    //                    /  \
    //                   4    6
    // <img src="https://assets.leetcode.com/uploads/2020/01/30/sample_1_1709.png" />
    // Input: root = [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
    // Output: 20
    // Explanation: Maximum sum in a valid Binary search tree is obtained in root node with key equal to 3.
    tree1 := &TreeNode{
        1, 
        &TreeNode{4, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}, },
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}, }, },
    }
    fmt.Println(maxSumBST(tree1)) // 20
    // Example 2:
    //         4
    //        /
    //       3
    //     /   \
    //    1     2
    // <img src="https://assets.leetcode.com/uploads/2020/01/30/sample_2_1709.png" />
    // Input: root = [4,3,null,1,2]
    // Output: 2
    // Explanation: Maximum sum in a valid Binary search tree is obtained in a single root node with key equal to 2.
    tree2 := &TreeNode{
        4, 
        &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{2, nil, nil},},
        nil, 
    }
    fmt.Println(maxSumBST(tree2)) // 2
    // Example 3:
    // Input: root = [-4,-2,-5]
    // Output: 0
    // Explanation: All values are negatives. Return an empty BST.
    tree3 := &TreeNode{
        -4, 
        &TreeNode{-2, nil, nil},
        &TreeNode{-5, nil, nil}, 
    }
    fmt.Println(maxSumBST(tree3)) // 0
}