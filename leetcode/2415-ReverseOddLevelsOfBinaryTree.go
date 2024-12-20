package main

// 2415. Reverse Odd Levels of Binary Tree
// Given the root of a perfect binary tree, reverse the node values at each odd level of the tree.
//     For example, suppose the node values at level 3 are [2,1,3,4,7,11,29,18], then it should become [18,29,11,7,4,3,1,2].

// Return the root of the reversed tree.

// A binary tree is perfect if all parent nodes have two children and all leaves are on the same level.

// The level of a node is the number of edges along the path between it and the root node.

// Example 1:
// <img src="https://assets.leetcode.com/uploads/2022/07/28/first_case1.png" />
// Input: root = [2,3,5,8,13,21,34]
// Output: [2,5,3,8,13,21,34]
// Explanation: 
// The tree has only one odd level.
// The nodes at level 1 are 3, 5 respectively, which are reversed and become 5, 3.

// Example 2:
// <img src="https://assets.leetcode.com/uploads/2022/07/28/second_case3.png" />
// Input: root = [7,13,11]
// Output: [7,11,13]
// Explanation: 
// The nodes at level 1 are 13, 11, which are reversed and become 11, 13.

// Example 3:
// Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
// Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
// Explanation: 
// The odd levels have non-zero values.
// The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
// The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1 after the reversal.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 2^14].
//     0 <= Node.val <= 10^5
//     root is a perfect binary tree.

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
func reverseOddLevels(root *TreeNode) *TreeNode {
    queue := []*TreeNode{ root }
    for i := 0; len(queue) > 0; i++ {
        t := []*TreeNode{}
        for k := len(queue); k > 0; k-- {
            node := queue[0] // pop
            queue = queue[1:]
            if i % 2 == 1 { // 找出奇数层的节点
                t = append(t, node)
            }
            if node.Left != nil {
                queue = append(queue, node.Left)
                queue = append(queue, node.Right)
            }
        }
        for l, r := 0, len(t) - 1; l < r; l, r = l + 1, r - 1 { // 交换 奇数层数值
            t[l].Val, t[r].Val = t[r].Val, t[l].Val 
        }
    }
    return root
}

func reverseOddLevels1(root *TreeNode) *TreeNode {
    if root == nil { return root }
    queue, flag := []*TreeNode{ root }, false
    for len(queue) > 0 {
        n := len(queue)
        for i := 0; i < n; i++ {
            node := queue[i]
            if node.Left != nil  { queue = append(queue, node.Left) }
            if node.Right != nil { queue = append(queue, node.Right) }
        }
        if flag { // 奇数层才需要交换
            for i := 0; i < n / 2; i++ {
                a, b := queue[i], queue[n-i-1]
                a.Val, b.Val = b.Val, a.Val
            }
        }
        flag = !flag
        queue = queue[n:]
    }
    return root
}

func main() {
    // Example 1:
    // <img src="https://assets.leetcode.com/uploads/2022/07/28/first_case1.png" />
    // Input: root = [2,3,5,8,13,21,34]
    // Output: [2,5,3,8,13,21,34]
    // Explanation: 
    // The tree has only one odd level.
    // The nodes at level 1 are 3, 5 respectively, which are reversed and become 5, 3.
    tree1 := &TreeNode {
        2,
        &TreeNode { 3, &TreeNode { 8, nil, nil },  &TreeNode { 13, nil, nil }, },
        &TreeNode { 5, &TreeNode { 21, nil, nil }, &TreeNode { 34, nil, nil }, },
    }
    fmt.Println(reverseOddLevels(tree1))
    // Example 2:
    // <img src="https://assets.leetcode.com/uploads/2022/07/28/second_case3.png" />
    // Input: root = [7,13,11]
    // Output: [7,11,13]
    // Explanation: 
    // The nodes at level 1 are 13, 11, which are reversed and become 11, 13.
    tree2 := &TreeNode {
        7,
        &TreeNode { 13, nil, nil, },
        &TreeNode { 11, nil, nil, },
    }
    fmt.Println(reverseOddLevels(tree2))
    // Example 3:
    // Input: root = [0,1,2,0,0,0,0,1,1,1,1,2,2,2,2]
    // Output: [0,2,1,0,0,0,0,2,2,2,2,1,1,1,1]
    // Explanation: 
    // The odd levels have non-zero values.
    // The nodes at level 1 were 1, 2, and are 2, 1 after the reversal.
    // The nodes at level 3 were 1, 1, 1, 1, 2, 2, 2, 2, and are 2, 2, 2, 2, 1, 1, 1, 1 after the reversal.
    tree3 := &TreeNode {
        0,
        &TreeNode { 1, &TreeNode { 0, &TreeNode { 1, nil, nil }, &TreeNode { 1, nil, nil }, }, &TreeNode { 0, &TreeNode { 1, nil, nil }, &TreeNode { 1, nil, nil }, }, },
        &TreeNode { 2, &TreeNode { 0, &TreeNode { 2, nil, nil }, &TreeNode { 2, nil, nil }, }, &TreeNode { 0, &TreeNode { 2, nil, nil }, &TreeNode { 2, nil, nil }, }, },
    }
    fmt.Println(reverseOddLevels(tree3))

    tree11 := &TreeNode {
        2,
        &TreeNode { 3, &TreeNode { 8, nil, nil },  &TreeNode { 13, nil, nil }, },
        &TreeNode { 5, &TreeNode { 21, nil, nil }, &TreeNode { 34, nil, nil }, },
    }
    fmt.Println(reverseOddLevels1(tree11))
    tree12 := &TreeNode {
        7,
        &TreeNode { 13, nil, nil, },
        &TreeNode { 11, nil, nil, },
    }
    fmt.Println(reverseOddLevels1(tree12))
    tree13 := &TreeNode {
        0,
        &TreeNode { 1, &TreeNode { 0, &TreeNode { 1, nil, nil }, &TreeNode { 1, nil, nil }, }, &TreeNode { 0, &TreeNode { 1, nil, nil }, &TreeNode { 1, nil, nil }, }, },
        &TreeNode { 2, &TreeNode { 0, &TreeNode { 2, nil, nil }, &TreeNode { 2, nil, nil }, }, &TreeNode { 0, &TreeNode { 2, nil, nil }, &TreeNode { 2, nil, nil }, }, },
    }
    fmt.Println(reverseOddLevels1(tree13))
}