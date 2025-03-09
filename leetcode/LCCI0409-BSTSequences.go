package main

// 面试题 04.09. BST Sequences LCCI
// A binary search tree was created by traversing through an array from left to right and inserting each element. 
// Given a binary search tree with distinct elements, print all possible arrays that could have led to this tree.

// Example:
// Given the following tree:
//         2
//        / \
//       1   3
// Output:
// [
//    [2,1,3],
//    [2,3,1]
// ]

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
func BSTSequences(root *TreeNode) [][]int {
    res := [][]int{}
    if root == nil {
        res = append(res, []int{})
        return res 
    }
    queue := []*TreeNode{root}
    var dfs func(queue []*TreeNode, list []int)
    dfs = func(queue []*TreeNode, list []int) {
        if len(queue) == 0 {
            res = append(res, append([]int{}, list...))
            return 
        }
        i := 0
        for i < len(queue) {
            queue[i], queue[0] = queue[0], queue[i]
            node := queue[0]
            newQueue := queue[1:]
            if node.Left != nil {
                newQueue = append(newQueue, node.Left)
            }
            if node.Right != nil {
                newQueue = append(newQueue, node.Right)
            }
            dfs(newQueue, append(list, queue[0].Val))
            queue[i], queue[0] = queue[0], queue[i]
            i++
        }
    }
    dfs(queue, []int{})
    return res
}

func main() {
    // Example:
    // Given the following tree:
    //         2
    //        / \
    //       1   3
    // Output:
    // [
    //    [2,1,3],
    //    [2,3,1]
    // ]
    tree1 := &TreeNode {
        2,
        &TreeNode{1, nil, nil, },
        &TreeNode{3, nil, nil, },
    }
    fmt.Println(BSTSequences(tree1)) // [[2,1,3], [2,3,1]]
    
    fmt.Println(BSTSequences(nil)) // [[]]
}