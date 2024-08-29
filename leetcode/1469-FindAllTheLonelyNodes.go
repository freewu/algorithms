package main

// 1469. Find All The Lonely Nodes
// In a binary tree, a lonely node is a node that is the only child of its parent node. 
// The root of the tree is not lonely because it does not have a parent node.

// Given the root of a binary tree, return an array containing the values of all lonely nodes in the tree. 
// Return the list in any order.

// Example 1:
//             1
//           /   \
//          2     3
//           \ 
//             (4)
// <img src="https://assets.leetcode.com/uploads/2020/06/03/e1.png">
// Input: root = [1,2,3,null,4]
// Output: [4]
// Explanation: Light blue node is the only lonely node.
// Node 1 is the root and is not lonely.
// Nodes 2 and 3 have the same parent and are not lonely.

// Example 2:
//             7
//           /    \
//          1      4
//         /      /  \
//       (6)     5    3
//                      \
//                      (2) 
// <img src="https://assets.leetcode.com/uploads/2020/06/03/e2.png">
// Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2]
// Output: [6,2]
// Explanation: Light blue nodes are lonely nodes.
// Please remember that order doesn't matter, [2,6] is also an acceptable answer.

// Example 3:
//             11
//           /     \
//          99     88
//         /         \
//       (77)        (66)
//       /             \
//     (55)            (44)
//     /                 \
//   (33)                (22)
// <img src="https://assets.leetcode.com/uploads/2020/06/03/tree.png">
// Input: root = [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
// Output: [77,55,33,66,44,22]
// Explanation: Nodes 99 and 88 share the same parent. Node 11 is the root.
// All other nodes are lonely.

// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     1 <= Node.val <= 10^6

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
func getLonelyNodes(root *TreeNode) []int {
    res := []int{}
    if root == nil || root.Left == nil && root.Right == nil {
        return res
    }
    if root.Left == nil { // 如果没有左子节点,把右边子节点的值写入返回数组
        res = append(res, root.Right.Val)
    } else if root.Right == nil { // 如果没有右子节点,把左边子节点的值写入返回数组
        res = append(res, root.Left.Val)
    }
    res = append(res, getLonelyNodes(root.Left)...)
    res = append(res, getLonelyNodes(root.Right)...)
    return res
}

func getLonelyNodes1(root *TreeNode) []int {
    res := []int{}
    var traverse func(root *TreeNode)
    traverse = func(root *TreeNode) {
        if root == nil {
            return
        }
        if root.Left != nil && root.Right == nil {
            res = append(res, root.Left.Val)
        }
        if root.Left == nil && root.Right != nil {
            res = append(res, root.Right.Val)
        }
        traverse(root.Left)
        traverse(root.Right)
    }
    traverse(root)
    return res
}

func main() {
    // Example 1:
    //             1
    //           /   \
    //          2     3
    //           \ 
    //            (4)
    // <img src="https://assets.leetcode.com/uploads/2020/06/03/e1.png">
    // Input: root = [1,2,3,null,4]
    // Output: [4]
    // Explanation: Light blue node is the only lonely node.
    // Node 1 is the root and is not lonely.
    // Nodes 2 and 3 have the same parent and are not lonely.
    tree1 := &TreeNode {
        1,
        &TreeNode{2, nil,  &TreeNode{4, nil, nil, }, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(getLonelyNodes(tree1)) // [4]
    // Example 2:
    //             7
    //           /    \
    //          1      4
    //         /      /  \
    //       (6)     5    3
    //                      \
    //                      (2) 
    // <img src="https://assets.leetcode.com/uploads/2020/06/03/e2.png">
    // Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2]
    // Output: [6,2]
    // Explanation: Light blue nodes are lonely nodes.
    // Please remember that order doesn't matter, [2,6] is also an acceptable answer.
    tree2 := &TreeNode {
        7,
        &TreeNode{1, &TreeNode{6, nil, nil, }, nil, },
        &TreeNode{4, &TreeNode{5, nil, nil, }, &TreeNode{3, nil, &TreeNode{2, nil, nil, }, }, },
    }
    fmt.Println(getLonelyNodes(tree2)) // [6,2]
    // Example 3:
    //             11
    //           /     \
    //          99     88
    //         /         \
    //       (77)        (66)
    //       /             \
    //     (55)            (44)
    //     /                 \
    //   (33)                (22)
    // <img src="https://assets.leetcode.com/uploads/2020/06/03/tree.png">
    // Input: root = [11,99,88,77,null,null,66,55,null,null,44,33,null,null,22]
    // Output: [77,55,33,66,44,22]
    // Explanation: Nodes 99 and 88 share the same parent. Node 11 is the root.
    // All other nodes are lonely.
    tree3 := &TreeNode {
        11,
        &TreeNode{99, &TreeNode{77, &TreeNode{55, &TreeNode{33, nil, nil, }, nil, }, nil, }, nil, },
        &TreeNode{88, nil, &TreeNode{66, nil, &TreeNode{44, nil, &TreeNode{22, nil, nil, }, }, }, },
    }
    fmt.Println(getLonelyNodes(tree3)) // [77,55,33,66,44,22]

    fmt.Println(getLonelyNodes1(tree1)) // [4]
    fmt.Println(getLonelyNodes1(tree2)) // [6,2]
    fmt.Println(getLonelyNodes1(tree3)) // [77,55,33,66,44,22]
}