package main

// 637. Average of Levels in Binary Tree
// Given the root of a binary tree, return the average value of the nodes on each level in the form of an array. Answers within 10-5 of the actual answer will be accepted.
 
// Example 1:
//         3          3.000
//        /  \
//       9   20       (9 + 20) / 2 = 14.5000
//           /  \ 
//          15   7    (15 + 7) / 2 = 11.000
// <img src="https://assets.leetcode.com/uploads/2021/03/09/avg1-tree.jpg" / >
// Input: root = [3,9,20,null,null,15,7]
// Output: [3.00000,14.50000,11.00000]
// Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
// Hence return [3, 14.5, 11].

// Example 2:
//          3          3.000
//         /  \
//         9   20       (9 + 20) / 2 = 14.5000
//        /  \ 
//       15   7         (15 + 7) / 2 = 11.000
// <img src="https://assets.leetcode.com/uploads/2021/03/09/avg2-tree.jpg" / >
// Input: root = [3,9,20,15,7]
// Output: [3.00000,14.50000,11.00000]
 
// Constraints:
//     The number of nodes in the tree is in the range [1, 10^4].
//     -2^31 <= Node.val <= 2^31 - 1

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
// dfs
func averageOfLevels(root *TreeNode) []float64 {
    res, cnt := []float64{}, []int{}
    var dfs func (node *TreeNode, level int, res *[]float64, cnt *[]int)
    dfs = func (node *TreeNode, level int, res *[]float64, cnt *[]int) {
        if node == nil { return }
        if level == len(*res) {
            (*res) = append(*res, float64(node.Val))
            (*cnt) = append(*cnt, 1)
        } else {
            (*res)[level] += float64(node.Val) // 累加同层的和
            (*cnt)[level]++ // 统计同层的个数
        }
        dfs(node.Left, level + 1, res, cnt)
        dfs(node.Right, level + 1, res, cnt)
    }
    dfs(root, 0, &res, &cnt)
    for i := range res {
        res[i] /= float64(cnt[i])
    }
    return res
}

// bfs
func averageOfLevels1(root *TreeNode) []float64 {
    res := []float64{}
    if root == nil {
        return res
    }
    queue := []*TreeNode{ root }
    for len(queue) > 0 {
        queueLen := len(queue)
        sum, cnt := 0, 0
        for i := 0; i < queueLen; i++ {
            currentNode := queue[0]
            queue = queue[1:]
            sum += currentNode.Val // 累加同层的和
            cnt++ // 统计同层的个数
            if currentNode.Left != nil {
                queue = append(queue, currentNode.Left)
            }
            if currentNode.Right != nil {
                queue = append(queue, currentNode.Right)
            }
        }
        res = append(res, float64(sum) / float64(cnt))
    }
    return res
}

func main() {
    // Example 1:
    //         3          3.000
    //        /  \
    //       9   20       (9 + 20) / 2 = 14.5000
    //           /  \ 
    //          15   7    (15 + 7) / 2 = 11.000
    // Input: root = [3,9,20,null,null,15,7]
    // Output: [3.00000,14.50000,11.00000]
    // Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5, and on level 2 is 11.
    // Hence return [3, 14.5, 11].
    tree1 := &TreeNode {
        3,
        &TreeNode{9, nil, nil},
        &TreeNode {
            20,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
    }
    fmt.Println(averageOfLevels(tree1)) // [3.00000,14.50000,11.00000]
    // Example 2:
    //          3          3.000
    //         /  \
    //         9   20       (9 + 20) / 2 = 14.5000
    //        /  \ 
    //       15   7         (15 + 7) / 2 = 11.000
    // <img src="https://assets.leetcode.com/uploads/2021/03/09/avg2-tree.jpg" / >
    // Input: root = [3,9,20,15,7]
    // Output: [3.00000,14.50000,11.00000]
    tree2 := &TreeNode {
        3,
        &TreeNode {
            9,
            &TreeNode{15, nil, nil},
            &TreeNode{7, nil, nil},
        },
        &TreeNode{20, nil, nil},
    }
    fmt.Println(averageOfLevels(tree2)) // [3.00000,14.50000,11.00000]

    fmt.Println(averageOfLevels1(tree1)) // [3.00000,14.50000,11.00000]
    fmt.Println(averageOfLevels1(tree2)) // [3.00000,14.50000,11.00000]
}