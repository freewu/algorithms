package main 

// 508. Most Frequent Subtree Sum
// Given the root of a binary tree, return the most frequent subtree sum. 
// If there is a tie, return all the values with the highest frequency in any order.

// The subtree sum of a node is defined as the sum of all the node values formed by the subtree rooted at that node (including the node itself).

// Example 1:
//         5
//       /   \
//      2    -3
// <img src="https://assets.leetcode.com/uploads/2021/04/24/freq1-tree.jpg" />
// Input: root = [5,2,-3]
// Output: [2,-3,4]

// Example 2:
//         5
//        /  \ 
//       2   -5
// <img src="https://assets.leetcode.com/uploads/2021/04/24/freq2-tree.jpg" />
// Input: root = [5,2,-5]
// Output: [2]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5

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
func findFrequentTreeSum(root *TreeNode) []int {
    res, mx, mp := []int{}, 0, make(map[int]int)
    var dfs func(root *TreeNode) int
    dfs = func(root *TreeNode) int {
        if root == nil {
            return 0
        }
        curr := root.Val + dfs(root.Left) + dfs(root.Right)
        mp[curr]++
        if mp[curr] > mx {
            mx = mp[curr]
        }
        return curr
    }
    dfs(root)
    // for _, v:= range mp { // 找到最大值
    //     if v > mx {
    //         mx = v
    //     }
    // }
    for k, v := range mp {
        if v == mx {
            res = append(res, k)
        }
    }
    return res
}

func main() {
    // Example 1:
    //         5
    //       /   \
    //      2    -3
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/freq1-tree.jpg" />
    // Input: root = [5,2,-3]
    // Output: [2,-3,4]
    tree1 := &TreeNode {
        5,
        &TreeNode{2, nil, nil},
        &TreeNode{-3, nil, nil},
    }
    fmt.Println(findFrequentTreeSum(tree1)) // [2,-3,4]
    // Example 2:
    //         5
    //        /  \ 
    //       2   -5
    // <img src="https://assets.leetcode.com/uploads/2021/04/24/freq2-tree.jpg" />
    // Input: root = [5,2,-5]
    // Output: [2]
    tree2 := &TreeNode {
        5,
        &TreeNode{2, nil, nil},
        &TreeNode{-5, nil, nil},
    }
    fmt.Println(findFrequentTreeSum(tree2)) // [2]
}