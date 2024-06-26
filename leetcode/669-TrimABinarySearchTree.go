package main

// 669. Trim a Binary Search Tree
// Given the root of a binary search tree and the lowest and highest boundaries as low and high, 
// trim the tree so that all its elements lies in [low, high]. 
// Trimming the tree should not change the relative structure of the elements that will remain in the tree (i.e., any node's descendant should remain a descendant). It can be proven that there is a unique answer.

// Return the root of the trimmed binary search tree. Note that the root may change depending on the given bounds.

// Example 1:
//         1                     1
//       /    \      =>           \ 
//      0      2                    2
// <img src="https://assets.leetcode.com/uploads/2020/09/09/trim1.jpg" /> 
// Input: root = [1,0,2], low = 1, high = 2
// Output: [1,null,2]

// Example 2:
//             3                 3
//           /   \              /
//          0     4            2
//           \         =>     /
//            2              1
//           /
//          1
// <img src="https://assets.leetcode.com/uploads/2020/09/09/trim2.jpg" /> 
// Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
// Output: [3,2,null,1]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     0 <= Node.val <= 10^4
//     The value of each node in the tree is unique.
//     root is guaranteed to be a valid binary search tree.
//     0 <= low <= high <= 10^4

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
func trimBST(root *TreeNode, low int, high int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val >= low && root.Val <= high {
        root.Left  = trimBST(root.Left, low, high) // Trim the left subtree
        root.Right = trimBST(root.Right, low, high)  // Trim the right subtree
    } else if root.Val < low {
        // if the root val is less than low then getting values from left 
        // will be even lower(binary tree rule) so we need to set the root to root.right;
        root = trimBST(root.Right, low, high) 
    } else if root.Val > high {
        // if the root val is greater than high then getting values from right 
        // will be even higher(binary tree rule) so we need to the root to root.left.
        root = trimBST(root.Left, low, high) 
    }
    return root
}

func trimBST1(root *TreeNode, low int, high int) *TreeNode {
    var dfs func(root *TreeNode, low int, high int) *TreeNode
    dfs = func(root *TreeNode, low int, high int) *TreeNode {
        if root == nil {
            return nil
          }
          // Trim the left and right subtrees
          root.Left, root.Right = trimBST(root.Left, low, high), trimBST(root.Right, low, high)
          // If the current node's value is less than low, return the right subtree
          if root.Val < low {
            return root.Right
          }
          // If the current node's value is greater than high, return the left subtree
          if root.Val > high {
            return root.Left
          }
          // Otherwise, the current node is within the range
          return root
    }
    return dfs(root,low, high)
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2020/09/09/trim1.jpg" /> 
    // Input: root = [1,0,2], low = 1, high = 2
    // Output: [1,null,2]
    tree1 := &TreeNode {
        1,
        &TreeNode{0, nil, nil, },
        &TreeNode{2, nil, nil, },
    }
    fmt.Println(trimBST(tree1,1,2))
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/09/trim2.jpg" /> 
    // Input: root = [3,0,4,null,2,null,null,1], low = 1, high = 3
    // Output: [3,2,null,1]
    tree2 := &TreeNode {
        3,
        &TreeNode{0, nil, &TreeNode{2, &TreeNode{1, nil, nil, }, nil, }, },
        &TreeNode{4, nil, nil, },
    }
    fmt.Println(trimBST(tree2,1,3))

    fmt.Println(trimBST1(tree1,1,2))
    fmt.Println(trimBST1(tree2,1,3))
}