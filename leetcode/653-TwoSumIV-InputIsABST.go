package main

// 653. Two Sum IV - Input is a BST
// Given the root of a binary search tree and an integer k, 
// return true if there exist two elements in the BST such that their sum is equal to k, or false otherwise.

// Example 1:
//            {5}
//           /   \
//         (3)   (6)
//        /   \     \
//      [2]    {4}   [7]
// <img src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_1.jpg" />
// Input: root = [5,3,6,2,4,null,7], k = 9
// Output: true

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_2.jpg" />
// Input: root = [5,3,6,2,4,null,7], k = 28
// Output: false

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^4 <= Node.val <= 10^4
//     root is guaranteed to be a valid binary search tree.
//     -10^5 <= k <= 10^5

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
func findTarget(root *TreeNode, k int) bool {
    mp := make(map[int]int)
    var dfs func(*TreeNode) bool 
    dfs = func(node *TreeNode) bool {
        if node == nil {
            return false
        }
        if _, ok := mp[node.Val]; ok { // 如果发现存在返回 true
            return true
        }
        mp[k - node.Val]++ // 遍历树 把  k - node.Val 存在 map 中
        return dfs(node.Left) || dfs(node.Right)
    }
    return dfs(root)
}

func findTarget1(root *TreeNode, k int) bool {
    arr := make([]int, 0)
    var ergodic func(root *TreeNode)
    ergodic = func(root *TreeNode) {
        if root == nil {
            return
        }
        ergodic(root.Left)
        arr = append(arr, root.Val)
        ergodic(root.Right)
    }
    ergodic(root) // // 中序遍历得到有序数组
    left, right := 0, len(arr)-1
    for left < right {
        if arr[left] + arr[right] < k {
            left++
        } else if arr[left] + arr[right] > k {
            right--
        } else {
            return true
        }
    }
    return false
}

func main() {
    // Example 1:
    //            {5}
    //           /   \
    //         (3)   (6)
    //        /   \     \
    //      [2]    {4}   [7]
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_1.jpg" />
    // Input: root = [5,3,6,2,4,null,7], k = 9
    // Output: true
    tree1 := &TreeNode {
        5,
        &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}, },
        &TreeNode{6, nil,                    &TreeNode{7, nil, nil}, },
    }
    fmt.Println(findTarget(tree1, 9)) // true
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2020/09/21/sum_tree_2.jpg" />
    // Input: root = [5,3,6,2,4,null,7], k = 28
    // Output: false
    fmt.Println(findTarget(tree1, 28)) // false

    fmt.Println(findTarget1(tree1, 9)) // true
    fmt.Println(findTarget1(tree1, 28)) // false
}