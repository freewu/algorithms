package main

// 3902. Zigzag Level Sum of Binary Tree
// You are given the root of a binary tree.

// Traverse the tree level by level using a zigzag pattern:
//     1. At odd-numbered levels (1-indexed), traverse nodes from left to right.
//     2. At even-numbered levels, traverse nodes from right to left.

// While traversing a level in the specified direction, process nodes in order and stop immediately before the first node that violates the condition:
//     1. At odd levels: the node does not have a left child.
//     2. At even levels: the node does not have a right child.

// Only the nodes processed before this stopping condition contribute to the level sum.

// Return an integer array ans where ans[i] is the sum of the node values that are processed at level i + 1.

// Example 1:
// Input: root = [5,2,8,1,null,9,6]
// Output: [5,8,0]
// Explanation:
//                 5 
//               /    \
//              2      8
//             /      /  \
//            1      9    6
// ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22054am.png" />
// At level 1, nodes are processed left to right. Node 5 is included, thus ans[0] = 5.
// At level 2, nodes are processed right to left. Node 8 is included, but node 2 lacks a right child, so processing stops, thus ans[1] = 8.
// At level 3, nodes are processed left to right. The first node 1 lacks a left child, so no nodes are included, and ans[2] = 0.
// Thus, ans = [5, 8, 0].

// Example 2:
// Input: root = [1,2,3,4,5,null,7]
// Output: [1,5,0]
// Explanation:
//                 1 
//               /    \
//              2      3
//             /  \      \
//            4     5     7
// <img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22232am.png" />
// At level 1, nodes are processed left to right. Node 1 is included, thus ans[0] = 1.
// At level 2, nodes are processed right to left. Nodes 3 and 2 are included since both have right children, thus ans[1] = 3 + 2 = 5.
// At level 3, nodes are processed left to right. The first node 4 lacks a left child, so no nodes are included, and ans[2] = 0.
// Thus, ans = [1, 5, 0].

// Constraints:
//     The number of nodes in the tree is in the range [1, 10^5].
//     -10^5 <= Node.val <= 10^5

import "fmt"

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
func zigzagLevelSum(root *TreeNode) []int64 {
    res := []int64{}
    if root == nil { return res } // 处理根节点为空的情况
    queue := []*TreeNode{root}
    level := 1
    for len(queue) > 0 {
        var childNodes []*TreeNode
        violated := false
        // 追加当前层的和，初始为0
        res = append(res, 0)
        // 逆序遍历 curNodes[::-1]
        for i := len(queue) - 1; i >= 0; i-- {
            node := queue[i]
            if level & 1 != 0 {  // level & 1 等价于 奇数层
                // 交换左右节点
                node.Left, node.Right = node.Right, node.Left
            }
            if node.Right == nil {
                violated = true
            }
            if !violated { // 未违反则累加值
                res[len(res)-1] += int64(node.Val)
            }
            if node.Right != nil {
                childNodes = append(childNodes, node.Right)
            }
            if node.Left != nil {
                childNodes = append(childNodes, node.Left)
            }
        }
        level++
        queue = childNodes
    }
    return res
}

func main() {
    // Example 1:
    // Input: root = [5,2,8,1,null,9,6]
    // Output: [5,8,0]
    // Explanation:
    //                 5 
    //               /    \
    //              2      8
    //             /      /  \
    //            1      9    6
    // ​​​​​​<img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22054am.png" />
    // At level 1, nodes are processed left to right. Node 5 is included, thus ans[0] = 5.
    // At level 2, nodes are processed right to left. Node 8 is included, but node 2 lacks a right child, so processing stops, thus ans[1] = 8.
    // At level 3, nodes are processed left to right. The first node 1 lacks a left child, so no nodes are included, and ans[2] = 0.
    // Thus, ans = [5, 8, 0].
    tree1 := &TreeNode{
        Val: 5, 
        Left:  &TreeNode{ 
            Val: 2,
            Left: &TreeNode{ Val: 1},
        }, 
        Right: &TreeNode{ 
            Val: 8,
            Left: &TreeNode{ Val: 9},
            Right: &TreeNode{ Val: 6},
        },
    }
    fmt.Println(zigzagLevelSum(tree1)) // [5,8,0]
    // Example 2:
    // Input: root = [1,2,3,4,5,null,7]
    // Output: [1,5,0]
    // Explanation:
    //                 1 
    //               /    \
    //              2      3
    //             /  \      \
    //            4     5     7
    // <img src="https://assets.leetcode.com/uploads/2026/04/12/screenshot-2026-04-13-at-22232am.png" />
    // At level 1, nodes are processed left to right. Node 1 is included, thus ans[0] = 1.
    // At level 2, nodes are processed right to left. Nodes 3 and 2 are included since both have right children, thus ans[1] = 3 + 2 = 5.
    // At level 3, nodes are processed left to right. The first node 4 lacks a left child, so no nodes are included, and ans[2] = 0.
    // Thus, ans = [1, 5, 0].
    tree2 := &TreeNode{
        Val: 1, 
        Left:  &TreeNode{ 
            Val: 2,
            Left: &TreeNode{ Val: 4},
            Right: &TreeNode{ Val: 5},
        }, 
        Right: &TreeNode{ 
            Val: 3,
            Right: &TreeNode{ Val: 7},
        },
    }
    fmt.Println(zigzagLevelSum(tree2)) // [1,5,0]
}

