package main

// 1161. Maximum Level Sum of a Binary Tree
// Given the root of a binary tree, the level of its root is 1, the level of its children is 2, and so on.
// Return the smallest level x such that the sum of all the values of nodes at level x is maximal.

// Example 1:
//         1
//       /   \
//      7     0
//    /  \
//   7    -8
// <img src="https://assets.leetcode.com/uploads/2019/05/03/capture.JPG" />
// Input: root = [1,7,0,7,-8,null,null]
// Output: 2
// Explanation: 
// Level 1 sum = 1.
// Level 2 sum = 7 + 0 = 7.
// Level 3 sum = 7 + -8 = -1.
// So we return the level with the maximum sum which is level 2.

// Example 2:
// Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
// Output: 2
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -10^5 <= Node.val <= 10^5

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
// bfs
func maxLevelSum(root *TreeNode) int {
    res, index, level, queue := -1 << 32 - 1, -1, 0, []*TreeNode{root}
    for len(queue) > 0 {
        sum, l := 0, len(queue) // 只取本层的 l 
        for i := 0; i < l; i++ {
            node := queue[0] // 出队列
            queue = queue[1:]
            sum += node.Val // 累加每层的值
            if node.Left != nil { // 如果存在左节点，加入到队列中
                queue = append(queue, node.Left)
            }   
            if node.Right != nil { // 如果存在右节点，加入到队列中
                queue = append(queue, node.Right)
            }
        }
        level++
        if sum > res { // 如果是大的则
            res, index = sum, level 
        }
    }
    return index
}

// dfs
func maxLevelSum1(root *TreeNode) int {
    levelSums := []int{}
    var dfs func(node *TreeNode, level int)
    dfs = func(node *TreeNode, level int) {
        if node == nil { return }
        if len(levelSums) == level {
            levelSums = append(levelSums, node.Val)
        } else {
            levelSums[level] += node.Val
        }
        dfs(node.Left, level + 1)
        dfs(node.Right, level + 1)
    }
    dfs(root, 0)
    maxSum, index := levelSums[0], 0
    for i := 1; i < len(levelSums); i++ {
        if levelSums[i] > maxSum {
            maxSum = levelSums[i]
            index = i 
        }
    }
    return index + 1
}

func main() {
    // Example 1:
    //         1
    //       /   \
    //      7     0
    //    /  \
    //   7    -8
    // <img src="https://assets.leetcode.com/uploads/2019/05/03/capture.JPG" />
    // Input: root = [1,7,0,7,-8,null,null]
    // Output: 2
    // Explanation: 
    // Level 1 sum = 1.
    // Level 2 sum = 7 + 0 = 7.
    // Level 3 sum = 7 + -8 = -1.
    // So we return the level with the maximum sum which is level 2.
    tree1 := &TreeNode{
        1, 
        &TreeNode{7, &TreeNode{7, nil, nil}, &TreeNode{-8, nil, nil}, },
        &TreeNode{0, nil, nil},
    }
    fmt.Println(maxLevelSum(tree1)) // 2
    // Example 2:
    // Input: root = [989,null,10250,98693,-89388,null,null,null,-32127]
    // Output: 2
    tree2 := &TreeNode{
        989, 
        nil,
        &TreeNode{10250, &TreeNode{98693, nil, nil}, &TreeNode{-89388, nil, &TreeNode{-32127, nil, nil}, }, },
    }
    fmt.Println(maxLevelSum(tree2)) // 2

    fmt.Println(maxLevelSum1(tree1)) // 2
    fmt.Println(maxLevelSum1(tree2)) // 2
}