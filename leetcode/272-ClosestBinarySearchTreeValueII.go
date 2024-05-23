package main

// 272. Closest Binary Search Tree Value II
// Given the root of a binary search tree, a target value, and an integer k, return the k values in the BST that are closest to the target. 
// You may return the answer in any order.
// You are guaranteed to have only one unique set of k values in the BST that are closest to the target.

// Example 1:
//         4
//       /   \ 
//      2     5
//    /   \
//   1     3
// <img src="https://assets.leetcode.com/uploads/2021/03/12/closest1-1-tree.jpg" />
// Input: root = [4,2,5,1,3], target = 3.714286, k = 2
// Output: [4,3]

// Example 2:
// Input: root = [1], target = 0.000000, k = 1
// Output: [1]
 
// Constraints:
//     The number of nodes in the tree is n.
//     1 <= k <= n <= 10^4.
//     0 <= Node.val <= 10^9
//     -10^9 <= target <= 10^9
 
// Follow up: Assume that the BST is balanced. Could you solve it in less than O(n) runtime (where n = total nodes)?

import "fmt"
import "math"

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
func closestKValues(root *TreeNode, target float64, k int) []int {
    p, res := 0, make([]int, k)
    for i := 0; i < k; i++ {
        res[i] = int(target) + 1e10
    }
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        dfs(node.Left)
        if math.Abs(target - float64(node.Val)) < math.Abs(target - float64(res[p % k])) {
            res[p % k] = node.Val
            p++
        }
        dfs(node.Right)
    }
    dfs(root)
    return res
}

func closestKValues1(root *TreeNode, target float64, k int) []int {
    res, count := []int{}, 0
    var InOrder func (root *TreeNode, target float64, k int)
    InOrder = func (root *TreeNode, target float64, k int) {
        if root == nil {
            return 
        }
        InOrder(root.Left, target, k)
        if count < k {
            res = append(res, root.Val)
            count++
        } else {
            if math.Abs(float64(root.Val) - target) < math.Abs(float64(res[0]) - target) {
                res = append(res, root.Val)
                res = res[1:]
            } else {
                return 
            }
        }
        InOrder(root.Right, target, k)
    }
    InOrder(root, target, k)
    return res 
}

func main() {
    // Example 1:
    //         4
    //       /   \ 
    //      2     5
    //    /   \
    //   1     3
    // <img src="https://assets.leetcode.com/uploads/2021/03/12/closest1-1-tree.jpg" />
    // Input: root = [4,2,5,1,3], target = 3.714286, k = 2
    // Output: [4,3]
    tree1 := &TreeNode {
        4,
        &TreeNode{2, &TreeNode{1, nil, nil }, &TreeNode{3, nil, nil }, },
        &TreeNode{5, nil, nil },
    }
    fmt.Println(closestKValues(tree1, 3.714286, 2)) // [4,3]
    // Example 2:
    // Input: root = [1], target = 0.000000, k = 1
    // Output: [1]
    tree2 := &TreeNode{1, nil, nil }
    fmt.Println(closestKValues(tree2, 0.000000, 1)) // [1]

    fmt.Println(closestKValues1(tree1, 3.714286, 2)) // [4,3]
    fmt.Println(closestKValues1(tree2, 0.000000, 1)) // [1]
}