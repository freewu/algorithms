package main

// 1214. Two Sum BSTs
// Given the roots of two binary search trees, root1 and root2, 
// return true if and only if there is a node in the first tree and a node in the second tree whose values sum up to a given integer target.

// Example 1:
//         2         1
//        /  \      /  \
//       1    4    0    3
// <img src="https://assets.leetcode.com/uploads/2021/02/10/ex1.png" />
// Input: root1 = [2,1,4], root2 = [1,0,3], target = 5
// Output: true
// Explanation: 2 and 3 sum up to 5.

// Example 2: 
//         0            5
//        /  \        /   \
//      -10  10      1     7
//                 /   \
//                0     2
// <img src="https://assets.leetcode.com/uploads/2021/02/10/ex2.png" />
// Input: root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
// Output: false

// Constraints:
//     The number of nodes in each tree is in the range [1, 5000].
//     -10^9 <= Node.val, target <= 10^9

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
func twoSumBSTs(root1 *TreeNode, root2 *TreeNode, target int) bool {
    var search func(node *TreeNode, target int) bool 
    search = func(node *TreeNode, target int) bool {
        if node == nil {
            return false
        }
        if node.Val == target {
            return true
        } else if node.Val<target {
            return search(node.Right, target)
        } else {
            return search(node.Left, target)
        }
    }
    var traversal func(node *TreeNode) bool
    traversal = func(node *TreeNode) bool {
        if node.Left!=nil && traversal(node.Left) {
            return true
        }
        if search(root2, target-node.Val) { // 二分查找 root2
            return true
        }
        if node.Right!=nil && traversal(node.Right) {
            return true
        }
        return false
    }
    return traversal(root1) // 遍历 root1
}

func main() {
    // Example 1:
    //         2         1
    //        /  \      /  \
    //       1    4    0    3
    // <img src="https://assets.leetcode.com/uploads/2021/02/10/ex1.png" />
    // Input: root1 = [2,1,4], root2 = [1,0,3], target = 5
    // Output: true
    // Explanation: 2 and 3 sum up to 5.
    tree11 := &TreeNode{
        2, 
        &TreeNode{1, nil, nil},
        &TreeNode{4, nil, nil},
    }
    tree12 := &TreeNode{
        1, 
        &TreeNode{0, nil, nil},
        &TreeNode{3, nil, nil},
    } 
    fmt.Println(twoSumBSTs(tree11, tree12, 5)) // true
    // Example 2: 
    //         0            5
    //        /  \        /   \
    //      -10  10      1     7
    //                 /   \
    //                0     2
    // <img src="https://assets.leetcode.com/uploads/2021/02/10/ex2.png" />
    // Input: root1 = [0,-10,10], root2 = [5,1,7,0,2], target = 18
    // Output: false
    tree21 := &TreeNode{
        0, 
        &TreeNode{-10, nil, nil},
        &TreeNode{-10, nil, nil},
    }
    tree22 := &TreeNode{
        5, 
        &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, nil}, },
        &TreeNode{7, nil, nil},
    } 
    fmt.Println(twoSumBSTs(tree21, tree22, 18)) // false
}