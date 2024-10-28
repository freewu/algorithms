package main

// 2764. Is Array a Preorder of Some ‌Binary Tree
// Given a 0-indexed integer 2D array nodes, 
// your task is to determine if the given array represents the preorder traversal of some binary tree.

// For each index i, nodes[i] = [id, parentId], 
// where id is the id of the node at the index i and parentId is the id of its parent in the tree 
// (if the node has no parent, then parentId == -1).

// Return true if the given array represents the preorder traversal of some tree, and false otherwise.

// Note: the preorder traversal of a tree is a recursive way to traverse a tree in which we first visit the current node, 
// then we do the preorder traversal for the left child, and finally, we do it for the right child.

// Example 1:
// Input: nodes = [[0,-1],[1,0],[2,0],[3,2],[4,2]]
// Output: true
// Explanation: The given nodes make the tree in the picture below.
// We can show that this is the preorder traversal of the tree, first we visit node 0, then we do the preorder traversal of the right child which is [1], then we do the preorder traversal of the left child which is [2,3,4].
// <img src="https://assets.leetcode.com/uploads/2023/07/04/1.png" />
//         0
//       /   \
//      1     2
//          /    \
//         3      4

// Example 2:
// Input: nodes = [[0,-1],[1,0],[2,0],[3,1],[4,1]]
// Output: false
// Explanation: The given nodes make the tree in the picture below.
// For the preorder traversal, first we visit node 0, then we do the preorder traversal of the right child which is [1,3,4], but we can see that in the given order, 2 comes between 1 and 3, so, it's not the preorder traversal of the tree.
// <img src="https://assets.leetcode.com/uploads/2023/07/04/2.png" />
//         0
//       /   \
//      1     2
//     /  \
//    3    4

// Constraints:
//     1 <= nodes.length <= 10^5
//     nodes[i].length == 2
//     0 <= nodes[i][0] <= 10^5
//     -1 <= nodes[i][1] <= 10^5
//     The input is generated such that nodes make a binary tree.

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func isPreorder(nodes [][]int) bool {
    mp := make(map[int]*TreeNode)
    for _, v := range nodes {
        id := v[0]
        mp[id] = &TreeNode{ Val: id}
    }
    // step1 还原二叉树
    var root *TreeNode
    for _, v := range nodes {
        id, parentId := v[0], v[1]
        if parentId == -1 {
            root = mp[id]
        } else {
            parent := mp[parentId]
            if parent.Left == nil {
                parent.Left = mp[id]
                continue
            }
            if parent.Right == nil {
                parent.Right = mp[id]
                continue
            }
        }
    }
    // step2 得出其先序遍历
    preOrderNums := []int{}
    var preOrder func(root *TreeNode)
    preOrder = func(root *TreeNode) {
        if root == nil { return }
        preOrderNums = append(preOrderNums, root.Val)
        preOrder(root.Left)
        preOrder(root.Right)
    }
    preOrder(root)
    // 与 nodes 中的遍历进行比较
    for i, v := range nodes {
        if preOrderNums[i] != v[0] {
            return false
        }
    }
    return true
}

func main() {
    // Example 1:
    // Input: nodes = [[0,-1],[1,0],[2,0],[3,2],[4,2]]
    // Output: true
    // Explanation: The given nodes make the tree in the picture below.
    // We can show that this is the preorder traversal of the tree, first we visit node 0, then we do the preorder traversal of the right child which is [1], then we do the preorder traversal of the left child which is [2,3,4].
    // <img src="https://assets.leetcode.com/uploads/2023/07/04/1.png" />
    //         0
    //       /   \
    //      1     2
    //          /    \
    //         3      4
    fmt.Println(isPreorder([][]int{{0,-1},{1,0},{2,0},{3,2},{4,2}})) // true
    // Example 2:
    // Input: nodes = [[0,-1],[1,0],[2,0],[3,1],[4,1]]
    // Output: false
    // Explanation: The given nodes make the tree in the picture below.
    // For the preorder traversal, first we visit node 0, then we do the preorder traversal of the right child which is [1,3,4], but we can see that in the given order, 2 comes between 1 and 3, so, it's not the preorder traversal of the tree.
    // <img src="https://assets.leetcode.com/uploads/2023/07/04/2.png" />
    //         0
    //       /   \
    //      1     2
    //     /  \
    //    3    4
    fmt.Println(isPreorder([][]int{{0,-1},{1,0},{2,0},{3,1},{4,1}})) // false
}