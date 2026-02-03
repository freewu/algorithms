package main

// 3831. Median of a Binary Search Tree Level
// You are given the root of a Binary Search Tree (BST) and an integer level.

// The root node is at level 0. 
// Each level represents the distance from the root.

// Return the median value of all node values present at the given level. 
// If the level does not exist or contains no nodes, return -1.

// The median is defined as the middle element after sorting the values at that level in non-decreasing order. 
// If the number of values at that level is even, return the upper median (the larger of the two middle elements after sorting).

// Example 1:
// ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20801pm.png" />
// Input: root = [4,null,5,null,7], level = 2
// Output: 7
// Explanation:
// The nodes at level = 2 are [7]. The median value is 7.

// Example 2:
// ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20926pm.png" />
// Input: root = [6,3,8], level = 1
// Output: 8
// Explanation:
// The nodes at level = 1 are [3, 8]. There are two possible median values, so the larger one 8 is the answer.

// Example 3:
// ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-21001pm.png" />
// Input: root = [2,1], level = 2
// Output: -1
// Explanation:
// There is no node present at level = 2​​​​​​​, so the answer is -1.

// Constraints:
//     The number of nodes in the tree is in the range [1, 2 * 10^5].
//     1 <= Node.val <= 10^6
//     0 <= level <= 2 * 10^​​​​​​​5

import "fmt"
import "sort"

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
func levelMedian(root *TreeNode, level int) int {
    if root == nil { return -1 } // 处理根节点为空的边界情况
    queue := []*TreeNode{root} // 使用队列进行层序遍历，初始时加入根节点
    curr := 0
    // 遍历每一层
    for len(queue) > 0 {
        n := len(queue) // 当前层的节点数量
        arr := make([]int, 0, n) // 存储当前层的所有节点值
        for i := 0; i < n; i++ { // 处理当前层的所有节点 
            // 取出队首节点
            node := queue[0]
            queue = queue[1:]
            arr = append(arr, node.Val) // 记录当前节点值
            if node.Left != nil { // 将子节点加入队列（下一层）
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }
        }
        if curr == level { // 如果当前层级是目标层级，计算中位数
            if len(arr) == 0 {  return -1 }
            sort.Ints(arr) // 替换为标准库排序
            return arr[len(arr)/2]
        }
        curr++ // 进入下一层
    }
    return -1 // 目标层级不存在，返回-1
}

func main() {
    // Example 1:
    // ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20801pm.png" />
    // Input: root = [4,null,5,null,7], level = 2
    // Output: 7
    // Explanation:
    // The nodes at level = 2 are [7]. The median value is 7.
    tree1 := &TreeNode { 
        4, 
        nil, 
        &TreeNode { 5, nil, &TreeNode { 7, nil, nil, }, },
    }
    fmt.Println(levelMedian(tree1, 2)) // 7
    // Example 2:
    // ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-20926pm.png" />
    // Input: root = [6,3,8], level = 1
    // Output: 8
    // Explanation:
    // The nodes at level = 1 are [3, 8]. There are two possible median values, so the larger one 8 is the answer.
    tree2 := &TreeNode { 
        6, 
        &TreeNode { 3, nil, nil, }, 
        &TreeNode { 8, nil, nil, },
    }
    fmt.Println(levelMedian(tree2, 1)) // 8
    // Example 3:
    // ​​​​​​​​​​​​​<img src="https://assets.leetcode.com/uploads/2026/01/27/screenshot-2026-01-27-at-21001pm.png" />
    // Input: root = [2,1], level = 2
    // Output: -1
    // Explanation:
    // There is no node present at level = 2​​​​​​​, so the answer is -1.
    tree3 := &TreeNode { 
        2, 
        &TreeNode { 1, nil, nil, }, 
        nil,
    }
    fmt.Println(levelMedian(tree3, 2)) // -1
}