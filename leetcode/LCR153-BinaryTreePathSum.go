package main

// LCR 153. 二叉树中和为目标值的路径
// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
// 叶子节点 是指没有子节点的节点。

// 示例 1：
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsumii1.jpg" />
// 输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
// 输出：[[5,4,11,2],[5,8,4,5]]

// 示例 2：
// <img src="https://assets.leetcode.com/uploads/2021/01/18/pathsum2.jpg" />
// 输入：root = [1,2,3], targetSum = 5
// 输出：[]

// 示例 3：
// 输入：root = [1,2], targetSum = 0
// 输出：[]

// 提示：
//     树中节点总数在范围 [0, 5000] 内
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
func pathTarget(root *TreeNode, target int) [][]int {
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
    return dfs(root, []int{}, target, 0)
}

// bfs
func pathTarget1(root *TreeNode, target int) [][]int {
    if root == nil {
        return nil
    }
    res := [][]int{}
    nodes, sums, paths := []*TreeNode{ root }, []int{ target - root.Val }, [][]int{{ root.Val }}
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
    fmt.Println(pathTarget(tree1, 22)) // [[5,4,11,2],[5,8,4,5]]
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
    fmt.Println(pathTarget(tree2, 5)) // []
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
    fmt.Println(pathTarget(tree3,0)) // []

    fmt.Println(pathTarget1(tree1, 22)) // [[5,4,11,2],[5,8,4,5]]
    fmt.Println(pathTarget1(tree2, 5)) // []
    fmt.Println(pathTarget1(tree3, 0)) // []
}