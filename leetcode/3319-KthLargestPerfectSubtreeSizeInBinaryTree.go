package main

// 3319. K-th Largest Perfect Subtree Size in Binary Tree
// You are given the root of a binary tree and an integer k.

// Return an integer denoting the size of the kth largest perfect binary subtree, or -1 if it doesn't exist.

// A perfect binary tree is a tree where all leaves are on the same level, and every parent has two children.

// Example 1:
// Input: root = [5,3,6,5,2,5,7,1,8,null,null,6,8], k = 2
// Output: 3
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/10/14/tmpresl95rp-1.png" />
// The roots of the perfect binary subtrees are highlighted in black. Their sizes, in non-increasing order are [3, 3, 1, 1, 1, 1, 1, 1].
// The 2nd largest size is 3.

// Example 2:
// Input: root = [1,2,3,4,5,6,7], k = 1
// Output: 7
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/10/14/tmp_s508x9e-1.png" />
// The sizes of the perfect binary subtrees in non-increasing order are [7, 3, 3, 1, 1, 1, 1]. The size of the largest perfect binary subtree is 7.

// Example 3:
// Input: root = [1,2,3,null,4], k = 3
// Output: -1
// Explanation:
// <img src="https://assets.leetcode.com/uploads/2024/10/14/tmp74xnmpj4-1.png" />
// The sizes of the perfect binary subtrees in non-increasing order are [1, 1]. There are fewer than 3 perfect binary subtrees.

// Constraints:
//     The number of nodes in the tree is in the range [1, 2000].
//     1 <= Node.val <= 2000
//     1 <= k <= 1024

import "fmt"
import "sort"

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
func kthLargestPerfectSubtree(root *TreeNode, k int) int {
    arr := []int{}
    var dfs func(root *TreeNode) int 
    dfs = func(root *TreeNode) int {
        if root == nil { return 0 }
        left, right := dfs(root.Left), dfs(root.Right)
        if left > -1 && left == right {
            arr = append(arr, left * 2 + 1)
            return left * 2 + 1
        } else {
            return -1
        }
    }
    dfs(root)
    if len(arr) < k { return -1 } // 不够长度
    sort.Slice(arr, func(i, j int) bool { // 从大到小
        return arr[i] > arr[j]
    })
    // fmt.Println(arr)
    return arr[k - 1]
}

func main() {
    // Example 1:
    // Input: root = [5,3,6,5,2,5,7,1,8,null,null,6,8], k = 2
    // Output: 3
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/10/14/tmpresl95rp-1.png" />
    // The roots of the perfect binary subtrees are highlighted in black. Their sizes, in non-increasing order are [3, 3, 1, 1, 1, 1, 1, 1].
    // The 2nd largest size is 3.
    tree1 := &TreeNode{
        5, 
        &TreeNode{3, &TreeNode{5, &TreeNode{1, nil, nil, }, &TreeNode{8, nil, nil, }, }, &TreeNode{2, nil, nil, }, },
        &TreeNode{6, &TreeNode{5, &TreeNode{6, nil, nil, }, &TreeNode{8, nil, nil, }, }, &TreeNode{7, nil, nil, }, },
    }
    fmt.Println(kthLargestPerfectSubtree(tree1, 2)) // 3
    // Example 2:
    // Input: root = [1,2,3,4,5,6,7], k = 1
    // Output: 7
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/10/14/tmp_s508x9e-1.png" />
    // The sizes of the perfect binary subtrees in non-increasing order are [7, 3, 3, 1, 1, 1, 1]. The size of the largest perfect binary subtree is 7.
    tree2 := &TreeNode{
        1, 
        &TreeNode{2, &TreeNode{4, nil, nil, }, &TreeNode{5, nil, nil, }, },
        &TreeNode{3, &TreeNode{6, nil, nil, }, &TreeNode{7, nil, nil, }, },
    } 
    fmt.Println(kthLargestPerfectSubtree(tree2, 1)) // 7
    // Example 3:
    // Input: root = [1,2,3,null,4], k = 3
    // Output: -1
    // Explanation:
    // <img src="https://assets.leetcode.com/uploads/2024/10/14/tmp74xnmpj4-1.png" />
    // The sizes of the perfect binary subtrees in non-increasing order are [1, 1]. There are fewer than 3 perfect binary subtrees.
    tree3 := &TreeNode{
        1, 
        &TreeNode{2, nil, &TreeNode{4, nil, nil, }, },
        &TreeNode{3, nil, nil, },
    } 
    fmt.Println(kthLargestPerfectSubtree(tree3, 3)) // -1
}