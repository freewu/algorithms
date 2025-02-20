package main

// 面试题 04.12. Paths with Sum LCCI
// You are given a binary tree in which each node contains an integer value (which might be positive or negative). 
// Design an algorithm to count the number of paths that sum to a given value. 
// The path does not need to start or end at the root or a leaf, but it must go downwards (traveling only from parent nodes to child nodes).

// Example:
// Given the following tree and  sum = 22,
//               5
//              / \
//             4   8
//            /   / \
//           11  13  4
//          /  \    / \
//         7    2  5   1
// Output: 3
// Explanation: Paths that have sum 22 are: [5,4,11,2], [5,8,4,5], [4,11,7]

// Note:
//     node number <= 10000

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
func pathSum(root *TreeNode, sum int) int {
    res, mp := 0, make(map[int]int)
    var dfs func(node *TreeNode, target int)
    dfs = func(node *TreeNode, target int) {
        if node == nil {  return }
        target += node.Val
        if sum == target {
            res++
        }
        if v, ok := mp[target - sum]; ok {
            res += v
        }
        mp[target]++
        dfs(node.Left, target)
        dfs(node.Right, target)
        mp[target]--
    }
    dfs(root, 0)
    return res
}

func pathSum1(root *TreeNode, target int) int {
    res, path := 0, []int{0}
    var dfs func(node *TreeNode)
    dfs = func(node *TreeNode) {
        if node == nil { return }
        sum := path[len(path) - 1] + node.Val
        for _, v := range path {
            if sum - v == target {
                res++
            }
        }
        path = append(path, sum)
        dfs(node.Left)
        dfs(node.Right)
        path = path[0:len(path) - 1]
    }
    dfs(root)
    return res
}

func main() {
    // Example:
    // Given the following tree and  sum = 22,
    //               5
    //              / \
    //             4   8
    //            /   / \
    //           11  13  4
    //          /  \    / \
    //         7    2  5   1
    // Output: 3
    // Explanation: Paths that have sum 22 are: [5,4,11,2], [5,8,4,5], [4,11,7]
    tree1 := &TreeNode {
        5,
        &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil, }, &TreeNode{2, nil, nil, }, }, nil, },
        &TreeNode{8, &TreeNode{13, nil, nil, }, &TreeNode{4, &TreeNode{5, nil, nil, }, &TreeNode{1, nil, nil, }, }, },
    }
    fmt.Println(pathSum(tree1, 22)) // 3

    tree11 := &TreeNode {
        5,
        &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil, }, &TreeNode{2, nil, nil, }, }, nil, },
        &TreeNode{8, &TreeNode{13, nil, nil, }, &TreeNode{4, &TreeNode{5, nil, nil, }, &TreeNode{1, nil, nil, }, }, },
    }
    fmt.Println(pathSum1(tree11, 22)) // 3
}