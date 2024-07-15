package main

// 2196. Create Binary Tree From Descriptions
// You are given a 2D integer array descriptions where descriptions[i] = [parenti, childi, isLefti] indicates that parenti is the parent of childi in a binary tree of unique values. Furthermore,
//     If isLefti == 1, then childi is the left child of parenti.
//     If isLefti == 0, then childi is the right child of parenti.

// Construct the binary tree described by descriptions and return its root.
// The test cases will be generated such that the binary tree is valid.

// Example 1:
//          50
//        /    \
//      20      80
//     /  \     /
//   15    17  19
// <img src="https://assets.leetcode.com/uploads/2022/02/09/example1drawio.png" />
// Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
// Output: [50,20,80,15,17,19]
// Explanation: The root node is the node with value 50 since it has no parent.
// The resulting binary tree is shown in the diagram.

// Example 2:
//      1
//     /
//    2
//     \ 
//      3
//     /
//    4
// <img src="https://assets.leetcode.com/uploads/2022/02/09/example2drawio.png" />
// Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
// Output: [1,2,null,null,3,4]
// Explanation: The root node is the node with value 1 since it has no parent.
// The resulting binary tree is shown in the diagram.

// Constraints:
//     1 <= descriptions.length <= 10^4
//     descriptions[i].length == 3
//     1 <= parenti, childi <= 10^5
//     0 <= isLefti <= 1
//     The binary tree described by descriptions is valid.

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
func createBinaryTree(descriptions [][]int) *TreeNode {
    n, res := len(descriptions), &TreeNode{}
    if len(descriptions) == 0 { return res }
    cache := map[int]*TreeNode{}
    for i := 0; i < n; i++ {
        cache[descriptions[i][1]] = &TreeNode{ Val: descriptions[i][1], }
    }
    for i := 0;i < n; i++ {
        p, c, child := descriptions[i][0], descriptions[i][1], descriptions[i][2]
        pNode, isExist := cache[p]
        if !isExist {
            pNode = &TreeNode{  Val: p, }
            res = pNode
            cache[p] = res
        }
        cNode := cache[c]
        if child == 1 {
            pNode.Left = cNode
        } else {
            pNode.Right = cNode
        }
    }
    return res
}

func main() {
    // Example 1:
    //          50
    //        /    \
    //      20      80
    //     /  \     /
    //   15    17  19
    // <img src="https://assets.leetcode.com/uploads/2022/02/09/example1drawio.png" />
    // Input: descriptions = [[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]
    // Output: [50,20,80,15,17,19]
    // Explanation: The root node is the node with value 50 since it has no parent.
    // The resulting binary tree is shown in the diagram.
    fmt.Println(createBinaryTree([][]int{{20,15,1},{20,17,0},{50,20,1},{50,80,0},{80,19,1}})) // &{50 0xc000008090 0xc0000080a8}
    // Example 2:
    //      1
    //     /
    //    2
    //     \ 
    //      3
    //     /
    //    4
    // <img src="https://assets.leetcode.com/uploads/2022/02/09/example2drawio.png" />
    // Input: descriptions = [[1,2,1],[2,3,0],[3,4,1]]
    // Output: [1,2,null,null,3,4]
    // Explanation: The root node is the node with value 1 since it has no parent.
    // The resulting binary tree is shown in the diagram.
    fmt.Println(createBinaryTree([][]int{{1,2,1},{2,3,0},{3,4,1}})) // &{1 0xc000008120 <nil>}
}