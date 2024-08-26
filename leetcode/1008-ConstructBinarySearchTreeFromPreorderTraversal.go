package main

// 1008. Construct Binary Search Tree from Preorder Traversal
// Given an array of integers preorder, which represents the preorder traversal of a BST (i.e., binary search tree), construct the tree and return its root.
// It is guaranteed that there is always possible to find a binary search tree with the given requirements for the given test cases.
// A binary search tree is a binary tree where for every node, any descendant of Node.left has a value strictly less than Node.val, and any descendant of Node.right has a value strictly greater than Node.val.
// A preorder traversal of a binary tree displays the value of the node first, then traverses Node.left, then traverses Node.right.

// Example 1:
//             8
//           /   \
//          5     10
//        /   \     \
//       1     7     12
// <img src="https://assets.leetcode.com/uploads/2019/03/06/1266.png" />
// Input: preorder = [8,5,1,7,10,12]
// Output: [8,5,10,1,7,null,12]

// Example 2:
//         1
//           \
//             3
// Input: preorder = [1,3]
// Output: [1,null,3]

// Constraints:
//     1 <= preorder.length <= 100
//     1 <= preorder[i] <= 1000
//     All the values of preorder are unique.

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
func bstFromPreorder(preorder []int) *TreeNode {
    less := func(n int, arr []int) []int {
        res := []int{}
        for _, v := range arr { if v < n { res = append(res, v) } }
        return res
    }
    greater := func(n int, arr []int) []int {
        res := []int{}
        for _, v := range arr {  if v > n { res = append(res, v) } }
        return res
    }
    var helper func(preorder []int) *TreeNode
    helper = func(preorder []int) *TreeNode {
        if len(preorder) == 0 { 
            return nil 
        }
        node := &TreeNode{
            Val: preorder[0],
            Left: helper(less(preorder[0], preorder)),
            Right: helper(greater(preorder[0], preorder)),
        }
        return node
    }
    return helper(preorder)
}

func bstFromPreorder1(preorder []int) *TreeNode {
    var build func(preorder []int, start, end int) *TreeNode
    build = func(preorder []int, start, end int) *TreeNode {
        if start > end {
            return nil
        }
        rootVal := preorder[start]
        root := &TreeNode{Val: rootVal}
        p := start + 1
        for p <= end && preorder[p] < rootVal {
            p++
        }
        root.Left, root.Right = build(preorder, start+1, p-1), build(preorder, p, end) // 左子树区间: [start+1, p-1], 右子树区间: [p, end]
        return root
    }
    return build(preorder, 0, len(preorder)-1)
}

func main() {
    // Example 1:
    //             8
    //           /   \
    //          5     10
    //        /   \     \
    //       1     7     12
    // <img src="https://assets.leetcode.com/uploads/2019/03/06/1266.png" />
    // Input: preorder = [8,5,1,7,10,12]
    // Output: [8,5,10,1,7,null,12]
    fmt.Println(bstFromPreorder([]int{8,5,1,7,10,12})) // &{8 0xc000008078 0xc0000080a8}
    // Example 2:
    //         1
    //           \
    //             3
    // Input: preorder = [1,3]
    // Output: [1,null,3]
    fmt.Println(bstFromPreorder([]int{1,3})) // &{1 <nil> 0xc0000080f0}

    fmt.Println(bstFromPreorder1([]int{8,5,1,7,10,12})) // &{8 0xc000008078 0xc0000080a8}
    fmt.Println(bstFromPreorder1([]int{1,3})) // &{1 <nil> 0xc0000080f0}
}