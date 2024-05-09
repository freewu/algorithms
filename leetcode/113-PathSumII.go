package main

// 113. Path Sum II
// Given the root of a binary tree and an integer targetSum, 
// return all root-to-leaf paths where the sum of the node values in the path equals targetSum. 
// Each path should be returned as a list of the node values, not node references.

// A root-to-leaf path is a path starting from the root and ending at any leaf node. 
// A leaf is a node with no children.

// Example 1:
//            （5）
//            /   \
//          (4)   (8)
//          /     /  \
//        (11)   13   (4)
//        /  \        /  \
//       7   (2)     (5)  1
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg" />
// Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// Output: [[5,4,11,2],[5,8,4,5]]
// Explanation: There are two paths whose sum equals targetSum:
// 5 + 4 + 11 + 2 = 22
// 5 + 8 + 4 + 5 = 22

// Example 2:
//      1
//    /   \
//   2     3
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
// Input: root = [1,2,3], targetSum = 5
// Output: []

// Example 3:
// Input: root = [1,2], targetSum = 0
// Output: []
 
// Constraints:
//     The number of nodes in the tree is in the range [0, 5000].
//     -1000 <= Node.val <= 1000
//     -1000 <= targetSum <= 1000

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
// dfs
func pathSum(root *TreeNode, targetSum int) [][]int {
    var dfs func(node *TreeNode, res []int, targetSum, currentSum int) [][]int 
    dfs = func(node *TreeNode, res []int, targetSum, currentSum int) [][]int {
        if node == nil {
            return [][]int{}
        }
        res = append(res, node.Val)
        currentSum += node.Val
        if node.Left == nil && node.Right == nil && currentSum == targetSum {
            return [][]int{res}
        }
        // Without copying it to another recursion It would give a pointer of the list
        newGroupOfNodes := make([]int, len(res))
        copy(newGroupOfNodes, res)
        return append(dfs(node.Left, newGroupOfNodes, targetSum, currentSum), dfs(node.Right, newGroupOfNodes, targetSum, currentSum)...)
    }
    return dfs(root, []int{}, targetSum, 0)
}

// bfs
func pathSum1(root *TreeNode, targetSum int) [][]int {
    if root == nil {
        return nil
    }
    res := [][]int{}
    nodes, sums, paths := []*TreeNode{ root }, []int{ targetSum - root.Val }, [][]int{{ root.Val }}
    for len(nodes) > 0 {
        curr, sum, path := nodes[len(nodes)-1], sums[len(sums)-1], paths[len(paths)-1]
        nodes, sums, paths = nodes[:len(nodes)-1], sums[:len(sums)-1], paths[:len(paths)-1]
        if curr.Left == nil && curr.Right == nil && sum == 0 { // 满足要求的路径
            res = append(res, append([]int{}, path...))
        }
        if curr.Right != nil {
            nodes = append(nodes, curr.Right)
            sums = append(sums, sum-curr.Right.Val)
            paths = append(paths, append(append([]int{}, path...), curr.Right.Val))
        }
        if curr.Left != nil {
            nodes = append(nodes, curr.Left)
            sums = append(sums, sum-curr.Left.Val)
            paths = append(paths, append(append([]int{}, path...), curr.Left.Val))
        }
    }
    return res
}

func main() {
    // Example 1:
    //            （5）
    //            /   \
    //          (4)   (8)
    //          /     /  \
    //        (11)   13   (4)
    //        /  \        /  \
    //       7   (2)     (5)  1
    // <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg" />
    // Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
    // Output: [[5,4,11,2],[5,8,4,5]]
    // Explanation: There are two paths whose sum equals targetSum:
    // 5 + 4 + 11 + 2 = 22
    // 5 + 8 + 4 + 5 = 22
    tree1 := &TreeNode{
        5, 
        &TreeNode{
            4, 
            &TreeNode{11, &TreeNode{7, nil, nil }, &TreeNode{2, nil, nil }, }, 
            nil,
        }, 
        &TreeNode{
            8, 
            &TreeNode{13, nil, nil }, 
            &TreeNode{4, &TreeNode{5, nil, nil }, &TreeNode{1, nil, nil } },
        }, 
    }
    fmt.Println(pathSum(tree1, 22)) // [[5,4,11,2],[5,8,4,5]]
    // Example 2:
    //      1
    //    /   \
    //   2     3
    // <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
    // Input: root = [1,2,3], targetSum = 5
    // Output: []
    tree2 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil }, 
        &TreeNode{3, nil, nil }, 
    }
    fmt.Println(pathSum(tree2, 5)) // []
    // Example 3:
    //      1
    //    /  
    //   2     
    // Input: root = [1,2], targetSum = 0
    // Output: []
    tree3 := &TreeNode{
        1, 
        &TreeNode{2, nil, nil }, 
        nil, 
    }
    fmt.Println(pathSum(tree3,0)) // []

    fmt.Println(pathSum1(tree1, 22)) // [[5,4,11,2],[5,8,4,5]]
    fmt.Println(pathSum1(tree2, 5)) // []
    fmt.Println(pathSum1(tree3, 0)) // []
}