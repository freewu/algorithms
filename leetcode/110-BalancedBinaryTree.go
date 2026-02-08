package main

// 110. Balanced Binary Tree
// Given a binary tree, determine if it is height-balanced.

// Example 1:
//         3
//        /  \
//       9   20
//          /  \
//         15   7
// <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg" />
// Input: root = [3,9,20,null,null,15,7]
// Output: true

// Example 2:
//             1 
//           /   \
//          2     2
//        /   \
//       3     3
//     /   \
//    4     4
// <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg" />
// Input: root = [1,2,2,3,3,null,null,4,4]
// Output: false

// Example 3:
// <img src="" />
// Input: root = []
// Output: true
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 5000].
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
// 递归
func isBalanced(root *TreeNode) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    // iterative function to measure the maximum depth of a tree
    var maxDepth func(root *TreeNode) int 
    maxDepth = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        return max(maxDepth(root.Left), maxDepth(root.Right))+1
    }
    var helper func(root *TreeNode) bool
    helper = func(root *TreeNode) bool {
        if root == nil {
            return true
        }
        // getting maximum depth possible from both left and right children
        // check if the difference is not more than 1, if it is then the tree is unbalanced and we can stop our check
        if abs(maxDepth(root.Left) -  maxDepth(root.Right)) > 1 {
            return false
        }
        // if both children are balanced then this tree is also balanced
        return helper(root.Left) && helper(root.Right)
    }
    return helper(root)
}

func isBalanced1(root *TreeNode) bool {
    abs := func(x int) int { if x < 0 { return -x; }; return x; }
    max := func (x, y int) int { if x > y { return x; }; return y; }
    var dfs func(*TreeNode) int
    dfs = func(node *TreeNode) int {
        if node == nil { return 0  }
        left := dfs(node.Left)
        if left == -1 { return -1 }
        right := dfs(node.Right)
        if right == -1 || abs(left - right) > 1 { return -1 }
        return max(left, right) + 1
    }
    return dfs(root) != -1
}

func main() {
    // Example 1:
    //         3
    //        /  \
    //       9   20
    //          /  \
    //         15   7
    // <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_1.jpg" />
    // Input: root = [3,9,20,null,null,15,7]
    // Output: true
    tree1 := &TreeNode {
        3,
        &TreeNode{9, nil, nil},
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(isBalanced(tree1)) // true
    // Example 2:
    //             1 
    //           /   \
    //          2     2
    //        /   \
    //       3     3
    //     /   \
    //    4     4
    // <img src="https://assets.leetcode.com/uploads/2020/10/06/balance_2.jpg" />
    // Input: root = [1,2,2,3,3,null,null,4,4]
    // Output: false
    tree2 := &TreeNode {
        1,
        &TreeNode {
            2,
            &TreeNode{
                3, 
                &TreeNode{4, nil, nil }, 
                &TreeNode{4, nil, nil }, 
            },
            &TreeNode{3, nil, nil },
        },
        &TreeNode{2, nil, nil},
    }
    fmt.Println(isBalanced(tree2)) // false
    // Example 3:
    // <img src="" />
    // Input: root = []
    // Output: true
    fmt.Println(isBalanced(nil)) // false

    fmt.Println(isBalanced1(tree1)) // true
    fmt.Println(isBalanced1(tree2)) // false
    fmt.Println(isBalanced1(nil)) // false
}