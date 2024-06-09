package main

// 968. Binary Tree Cameras
// You are given the root of a binary tree. 
// We install cameras on the tree nodes where each camera at a node can monitor its parent, itself, and its immediate children.
// Return the minimum number of cameras needed to monitor all nodes of the tree.

// Example 1:
//         (0)
//         /
//       (c)
//      /   \   
//    (0)   (0)
// <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_01.png" />
// Input: root = [0,0,null,0,0]
// Output: 1
// Explanation: One camera is enough to monitor all nodes if placed as shown.

// Example 2:
//             (0)
//             /
//           (c)
//           /
//         (0)
//         /
//       (c)
//         \ 
//          (0)
// <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_02.png" />
// Input: root = [0,0,null,0,null,0,null,null,0]
// Output: 2
// Explanation: At least two cameras are needed to monitor all nodes of the tree. The above image shows one of the valid configurations of camera placement.
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 1000].
//     Node.val == 0

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
func minCameraCover(root *TreeNode) int {
    var dfs func(root *TreeNode) (int, int)
    dfs = func(root *TreeNode) (int, int) {
        if root == nil {
            return 0, 1
        }
        c1, l1 := dfs(root.Left)
        c2, l2 := dfs(root.Right)
        if l1 == 2 || l2 == 2 { // child and itself do not has camera, need camera on parent
            return c1 + c2 + 1, 0
        }
        if l1 == 0 || l2 == 0 { // with camera, not need camera, affect to parent
            return c1 + c2, 1
        }
        return c1 + c2, 2 // child has camera, not need camera, not affect to parent
    }
    c, l := dfs(root)
    if l == 2 {
        c++
    }
    return c
}

func minCameraCover1(root *TreeNode) int {
    cameras := 0
    var dfs func(root *TreeNode, cameras *int) int
    dfs = func(root *TreeNode, cameras *int) int {
        if root == nil {
            return 0
        }
        r := dfs(root.Left, cameras) | dfs(root.Right, cameras)
        *cameras += r & 1
        return int("1202"[r] - '0')
    }
    r := dfs(root, &cameras) & 1
    return cameras + r
}


func main() {
    // Example 1:
    //         (0)
    //         /
    //       (c)
    //      /   \   
    //    (0)   (0)
    // <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_01.png" />
    // Input: root = [0,0,null,0,0]
    // Output: 1
    // Explanation: One camera is enough to monitor all nodes if placed as shown.
    tree1 := &TreeNode {
        0,
        &TreeNode{ 0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}, },
        nil,
    }
    fmt.Println(minCameraCover(tree1)) // 1
    // Example 2:
    //             (0)
    //             /
    //           (c)
    //           /
    //         (0)
    //         /
    //       (c)
    //         \ 
    //          (0)
    // <img src="https://assets.leetcode.com/uploads/2018/12/29/bst_cameras_02.png" />
    // Input: root = [0,0,null,0,null,0,null,null,0]
    // Output: 2
    // Explanation: At least two cameras are needed to monitor all nodes of the tree. The above image shows one of the valid configurations of camera placement.
    tree2 := &TreeNode {
        0,
        &TreeNode{ 0, &TreeNode{0, &TreeNode{0, nil, &TreeNode{0, nil, nil} }, nil}, nil, },
        nil,
    }
    fmt.Println(minCameraCover(tree2)) // 2

    fmt.Println(minCameraCover1(tree1)) // 1
    fmt.Println(minCameraCover1(tree2)) // 2
}